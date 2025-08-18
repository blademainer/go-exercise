# 智能HTTP客户端 - IP故障检测与自动切换

这个项目实现了一个智能的HTTP客户端，具备以下核心功能：

## 核心特性

### 1. 多IP DNS解析
- 自动解析域名获取所有可用IP地址
- 维护每个域名的IP地址池
- 支持IPv4和IPv6（根据系统配置）

### 2. 智能负载均衡
- 使用轮询算法在健康的IP之间分配请求
- 支持权重调整（未来扩展）
- 实时更新可用IP列表

### 3. 健康检查与监控
- 实时监控每个IP的成功率
- 记录请求延迟、错误率等关键指标
- 支持自定义健康检查间隔和超时时间

### 4. 熔断器模式
- 自动检测故障IP并进行熔断
- 支持三种状态：健康(Healthy)、不健康(Unhealthy)、半开(Half-Open)
- 可配置的故障阈值和最小请求数

### 5. 半开恢复机制
- 定期对不健康的IP进行恢复检查
- 逐步恢复故障IP的流量
- 智能的恢复策略避免雪崩效应

## 架构设计

```
HTTP Client
    ↓
Smart Transport
    ↓
┌─────────────────┬─────────────────┬─────────────────┐
│   DNS Resolver  │ Health Monitor  │ Circuit Breaker │
│   多IP解析      │   健康检查      │    熔断器       │
└─────────────────┴─────────────────┴─────────────────┘
    ↓
IP Pool (IP地址池)
    ↓
Load Balancer (负载均衡)
    ↓
┌─────────────┬─────────────┐
│ Server 1    │ Server 2    │
│ (IP1)       │ (IP2)       │
└─────────────┴─────────────┘
```

## 使用方法

### 基本使用

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    // 创建智能HTTP客户端
    client := NewSmartHTTPClient()
    defer client.Close()
    
    // 发送请求（会自动进行IP故障检测和切换）
    resp, err := client.Get("http://example.com")
    if err != nil {
        fmt.Printf("Request failed: %v\n", err)
        return
    }
    defer resp.Body.Close()
    
    fmt.Printf("Response status: %d\n", resp.StatusCode)
    
    // 查看客户端状态
    client.PrintStatus()
}
```

### 高级配置

```go
// 创建自定义配置的智能客户端
resolver := NewSmartResolver()
// 可以调整以下参数：
// - 故障阈值：failureThreshold
// - 最小请求数：minRequests  
// - 恢复超时：recoveryTimeout
// - 健康检查间隔：healthCheckInterval

transport := &SmartHTTPTransport{
    resolver: resolver,
    // 其他自定义配置...
}

client := &http.Client{
    Transport: transport,
    Timeout:   30 * time.Second,
}
```

## 运行演示

### 1. 基本演示
```bash
cd demos/http/dns
go run .
```

### 2. 故障模拟演示
```bash
go run . simulate
```

### 3. 实时监控演示
```bash
go run . monitor
```

## 关键配置参数

| 参数 | 默认值 | 说明 |
|------|--------|------|
| `failureThreshold` | 3 | 连续失败多少次后触发熔断 |
| `minRequests` | 5 | 做出熔断决策的最小请求数 |
| `recoveryTimeout` | 30s | 从故障状态恢复检查的间隔 |
| `healthCheckInterval` | 10s | 健康检查的执行间隔 |
| `healthCheckTimeout` | 5s | 单次健康检查的超时时间 |

## 监控指标

智能客户端提供以下监控指标：

### IP级别指标
- `success_count`: 成功请求数
- `failure_count`: 失败请求数  
- `success_rate`: 成功率
- `consecutive_fails`: 连续失败次数
- `last_success`: 最后一次成功时间
- `last_failure`: 最后一次失败时间
- `status`: IP状态（健康/不健康/半开）

### 客户端级别指标
- `ip_pool_status`: 所有域名的IP池状态
- `round_robin_counters`: 轮询计数器状态

## 故障检测策略

### 1. 熔断触发条件
- 连续失败次数 >= `failureThreshold`
- 总请求数 >= `minRequests`
- HTTP状态码 >= 500
- 网络超时或连接错误

### 2. 恢复检测
- 不健康的IP会在`recoveryTimeout`后进入半开状态
- 半开状态下允许少量试探性请求
- 试探成功后逐步恢复正常流量

### 3. 负载均衡
- 优先使用健康状态的IP
- 在健康IP之间使用轮询算法
- 无健康IP时尝试半开状态的IP

## 实际应用场景

### 1. 微服务间通信
- 服务发现后的多实例负载均衡
- 实例故障时的自动切换
- 服务恢复的自动检测

### 2. 外部API调用
- 第三方服务的高可用访问
- CDN节点的智能选择
- API网关的故障转移

### 3. 数据库连接池
- 主从数据库的自动切换
- 读写分离场景下的负载均衡
- 数据库实例故障的快速检测

## 性能特性

- **低延迟**: 故障检测和切换在毫秒级完成
- **高并发**: 支持大量并发请求的智能路由
- **内存高效**: 采用轻量级的状态管理
- **CPU友好**: 异步健康检查不影响主请求流程

## 扩展性

该设计支持以下扩展：

1. **自定义负载均衡算法**: 权重轮询、最少连接等
2. **插件化监控**: 集成Prometheus、Grafana等监控系统
3. **配置中心集成**: 支持动态配置更新
4. **服务网格集成**: 与Istio、Linkerd等服务网格协作

## 注意事项

1. **DNS缓存**: 系统DNS缓存可能影响IP更新，建议配置合适的TTL
2. **连接复用**: 启用HTTP keep-alive可以提高性能
3. **超时设置**: 根据业务场景调整各级超时时间
4. **监控告警**: 建议集成监控系统进行故障告警

## 故障排查

### 常见问题

1. **所有IP都被标记为不健康**
   - 检查网络连通性
   - 确认健康检查URL的可达性
   - 调整故障阈值参数

2. **切换不及时**
   - 减少`failureThreshold`值
   - 缩短`healthCheckInterval`间隔
   - 检查DNS解析是否正常

3. **频繁切换**
   - 增加`minRequests`值
   - 延长`recoveryTimeout`时间
   - 检查网络稳定性

### 调试模式

启用详细日志可以帮助排查问题：

```go
// 在创建客户端前设置日志级别
log.SetLevel(log.DebugLevel)
```

## 贡献

欢迎提交Issue和Pull Request来改进这个项目！
