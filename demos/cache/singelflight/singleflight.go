package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

// CacheItem 缓存项结构
type CacheItem struct {
	Value      any
	ExpireTime time.Time
}

// Cache 带过期时间的缓存系统
type Cache struct {
	mu    sync.RWMutex
	items map[string]*CacheItem
	sg    singleflight.Group
}

// NewCache 创建新的缓存实例
func NewCache() *Cache {
	c := &Cache{
		items: make(map[string]*CacheItem),
	}
	// 启动定期清理过期缓存的goroutine
	go c.cleanExpiredItems()
	return c
}

// Get 获取缓存数据，如果不存在则通过加载函数获取
func (c *Cache) Get(ctx context.Context, key string, ttl time.Duration, loader func(context.Context, string) (any, error)) (any, error) {
	// 先尝试从缓存中获取
	if val, ok := c.getFromCache(key); ok {
		log.Printf("[Cache Hit] key=%s, value=%v", key, val)
		return val, nil
	}

	log.Printf("[Cache Miss] key=%s, using SingleFlight to load data", key)

	// 使用SingleFlight防止缓存击穿
	// 多个并发请求会被合并为一个
	v, err, shared := c.sg.Do(key, func() (any, error) {
		log.Printf("[SingleFlight] loading data for key=%s", key)
		
		// 再次检查缓存（可能在等待期间其他goroutine已经加载）
		if val, ok := c.getFromCache(key); ok {
			log.Printf("[Cache Hit in SingleFlight] key=%s, value=%v", key, val)
			return val, nil
		}

		// 从数据源加载数据
		data, err := loader(ctx, key)
		if err != nil {
			log.Printf("[Load Error] key=%s, error=%v", key, err)
			return nil, err
		}

		// 保存到缓存
		c.Set(key, data, ttl)
		log.Printf("[Loaded and Cached] key=%s, value=%v, ttl=%v", key, data, ttl)
		return data, nil
	})

	if shared {
		log.Printf("[SingleFlight Merged] key=%s, this request was merged with others", key)
	}

	return v, err
}

// getFromCache 从缓存中获取数据
func (c *Cache) getFromCache(key string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, exists := c.items[key]
	if !exists {
		return nil, false
	}

	// 检查是否过期
	if time.Now().After(item.ExpireTime) {
		return nil, false
	}

	return item.Value, true
}

// Set 设置缓存数据
func (c *Cache) Set(key string, value any, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = &CacheItem{
		Value:      value,
		ExpireTime: time.Now().Add(ttl),
	}
	log.Printf("[Cache Set] key=%s, value=%v, ttl=%v", key, value, ttl)
}

// Delete 删除缓存数据
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
	log.Printf("[Cache Delete] key=%s", key)
}

// cleanExpiredItems 定期清理过期的缓存项
func (c *Cache) cleanExpiredItems() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		now := time.Now()
		for key, item := range c.items {
			if now.After(item.ExpireTime) {
				delete(c.items, key)
				log.Printf("[Cache Expired] key=%s", key)
			}
		}
		c.mu.Unlock()
	}
}

// simulateDBQuery 模拟数据库查询
func simulateDBQuery(ctx context.Context, key string) (any, error) {
	log.Printf("[DB Query Start] querying database for key=%s", key)
	
	// 模拟数据库查询延迟
	time.Sleep(2 * time.Second)
	
	// 模拟返回数据
	result := fmt.Sprintf("data_for_%s_from_db", key)
	log.Printf("[DB Query Complete] key=%s, result=%s", key, result)
	
	return result, nil
}

// testConcurrentAccess 测试并发访问相同key的场景
func testConcurrentAccess(cache *Cache, key string, goroutineCount int) {
	log.Printf("\n========== Testing Concurrent Access ==========")
	log.Printf("Starting %d goroutines to request key=%s concurrently", goroutineCount, key)
	
	var wg sync.WaitGroup
	startTime := time.Now()

	for i := range goroutineCount {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			ctx := context.Background()
			val, err := cache.Get(ctx, key, 5*time.Second, simulateDBQuery)
			if err != nil {
				log.Printf("[Goroutine %d] Error: %v", id, err)
				return
			}
			log.Printf("[Goroutine %d] Got value: %v", id, val)
		}(i)
	}

	wg.Wait()
	duration := time.Since(startTime)
	log.Printf("========== Concurrent Access Complete ==========")
	log.Printf("Total time: %v (should be ~2s if SingleFlight works correctly)", duration)
	log.Printf("Expected: Without SingleFlight, it would take ~2s * %d goroutines\n", goroutineCount)
}

// testCacheExpiration 测试缓存过期
func testCacheExpiration(cache *Cache) {
	log.Printf("\n========== Testing Cache Expiration ==========")
	
	key := "expire_test"
	ctx := context.Background()
	
	// 第一次获取，会从数据库加载
	log.Printf("First access (will load from DB):")
	cache.Get(ctx, key, 3*time.Second, simulateDBQuery)
	
	// 立即第二次获取，应该从缓存命中
	log.Printf("\nSecond access (should hit cache):")
	cache.Get(ctx, key, 3*time.Second, simulateDBQuery)
	
	// 等待缓存过期
	log.Printf("\nWaiting 4 seconds for cache to expire...")
	time.Sleep(4 * time.Second)
	
	// 过期后再次获取，应该重新从数据库加载
	log.Printf("\nThird access after expiration (should reload from DB):")
	cache.Get(ctx, key, 3*time.Second, simulateDBQuery)
	
	log.Printf("========== Cache Expiration Test Complete ==========\n")
}

// testCacheOperations 测试基本缓存操作
func testCacheOperations(cache *Cache) {
	log.Printf("\n========== Testing Basic Cache Operations ==========")
	
	// 测试Set和Get
	cache.Set("user:1", "Alice", 10*time.Second)
	cache.Set("user:2", "Bob", 10*time.Second)
	
	if val, ok := cache.getFromCache("user:1"); ok {
		log.Printf("Retrieved: user:1 = %v", val)
	}
	
	// 测试Delete
	cache.Delete("user:1")
	if _, ok := cache.getFromCache("user:1"); !ok {
		log.Printf("Successfully deleted: user:1")
	}
	
	log.Printf("========== Basic Cache Operations Complete ==========\n")
}

func main() {
	log.SetFlags(log.Ltime | log.Lmicroseconds)
	log.Printf("========== SingleFlight Cache Demo Start ==========\n")

	// 创建缓存实例
	cache := NewCache()

	// 测试1: 基本缓存操作
	testCacheOperations(cache)

	// 测试2: 并发访问相同key（展示SingleFlight效果）
	testConcurrentAccess(cache, "product:100", 10)

	// 等待一段时间，确保上一个测试的缓存还在
	time.Sleep(1 * time.Second)

	// 测试3: 再次并发访问相同key（应该全部命中缓存）
	log.Printf("\n========== Testing Cache Hit Scenario ==========")
	log.Printf("Accessing the same key again (should all hit cache):")
	testConcurrentAccess(cache, "product:100", 5)

	// 测试4: 缓存过期测试
	testCacheExpiration(cache)

	// 测试5: 并发访问不同key
	log.Printf("\n========== Testing Concurrent Access with Different Keys ==========")
	var wg sync.WaitGroup
	keys := []string{"order:1", "order:2", "order:3"}
	
	for _, key := range keys {
		wg.Add(1)
		go func(k string) {
			defer wg.Done()
			ctx := context.Background()
			cache.Get(ctx, k, 5*time.Second, simulateDBQuery)
		}(key)
	}
	wg.Wait()
	log.Printf("========== Different Keys Test Complete ==========\n")

	log.Printf("\n========== SingleFlight Cache Demo Complete ==========")
}
