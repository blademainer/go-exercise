package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	// Check command line arguments for demo mode
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "simulate":
			FailureSimulation()
			return
		case "monitor":
			MonitoringDemo()
			return
		case "help":
			printUsage()
			return
		}
	}

	// Run basic demo
	runBasicDemo()
}

func printUsage() {
	fmt.Println("智能HTTP客户端演示程序")
	fmt.Println("用法:")
	fmt.Println("  go run . [模式]")
	fmt.Println()
	fmt.Println("可用模式:")
	fmt.Println("  (无参数)  - 运行基本智能客户端演示")
	fmt.Println("  simulate  - 运行故障模拟演示")
	fmt.Println("  monitor   - 运行实时监控演示")
	fmt.Println("  help      - 显示此帮助信息")
}

func runBasicDemo() {
	// Create smart HTTP client
	smartClient := NewSmartHTTPClient()
	defer smartClient.Close()

	// Set up signal handling for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("Received shutdown signal, stopping...")
		cancel()
	}()

	// Test URLs - you can modify these to test with your own domains
	testURLs := []string{
		"http://www.baidu.com",
		"http://www.google.com",
	}

	fmt.Println("=== Smart HTTP Client Demo ===")
	fmt.Println("This demo shows intelligent IP failover and health monitoring")
	fmt.Println("Press Ctrl+C to stop\n")

	// Start periodic status reporting
	go func() {
		ticker := time.NewTicker(15 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				smartClient.PrintStatus()
			}
		}
	}()

	// Run concurrent requests to test the smart client
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			runWorker(ctx, smartClient, testURLs, workerID)
		}(i)
	}

	// Wait for all workers to finish
	wg.Wait()

	// Print final status
	fmt.Println("\n=== Final Status ===")
	smartClient.PrintStatus()
}

// runWorker simulates continuous HTTP requests
func runWorker(ctx context.Context, client *SmartHTTPClient, urls []string, workerID int) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	urlIndex := 0

	for {
		select {
		case <-ctx.Done():
			log.Printf("Worker %d stopped", workerID)
			return
		case <-ticker.C:
			url := urls[urlIndex%len(urls)]
			urlIndex++

			go makeRequest(client, url, workerID)
		}
	}
}

// makeRequest performs a single HTTP request and handles the response
func makeRequest(client *SmartHTTPClient, url string, workerID int) {
	start := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Printf("Worker %d: Failed to create request for %s: %v", workerID, url, err)
		return
	}

	resp, err := client.Do(req)
	duration := time.Since(start)

	if err != nil {
		log.Printf("Worker %d: Request to %s failed after %v: %v", workerID, url, duration, err)
		return
	}
	defer resp.Body.Close()

	// Read response (limit to avoid memory issues)
	body, err := io.ReadAll(io.LimitReader(resp.Body, 1024))
	if err != nil {
		log.Printf("Worker %d: Failed to read response from %s: %v", workerID, url, err)
		return
	}

	log.Printf(
		"Worker %d: %s -> Status: %d, Size: %d bytes, Duration: %v",
		workerID, url, resp.StatusCode, len(body), duration,
	)
}
