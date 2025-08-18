package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

// FailureSimulation demonstrates the smart client handling various failure scenarios
func FailureSimulation() {
	fmt.Println("\n=== IP Failure Simulation Demo ===")
	fmt.Println("This demo simulates various network failure scenarios to test the smart client")

	// Create smart HTTP client
	smartClient := NewSmartHTTPClient()
	defer smartClient.Close()

	// Test scenarios
	scenarios := []struct {
		name        string
		url         string
		description string
		requests    int
		delay       time.Duration
	}{
		{
			name:        "Normal Operation",
			url:         "http://www.baidu.com",
			description: "Testing normal operation with a reliable host",
			requests:    5,
			delay:       1 * time.Second,
		},
		{
			name:        "Timeout Scenario",
			url:         "http://httpbin.org/delay/10", // Will timeout
			description: "Testing timeout handling",
			requests:    3,
			delay:       2 * time.Second,
		},
		{
			name:        "Server Error Scenario",
			url:         "http://httpbin.org/status/500", // Returns 500
			description: "Testing server error handling",
			requests:    3,
			delay:       1 * time.Second,
		},
		{
			name:        "Non-existent Host",
			url:         "http://this-host-does-not-exist-12345.com",
			description: "Testing DNS resolution failure",
			requests:    2,
			delay:       1 * time.Second,
		},
	}

	for _, scenario := range scenarios {
		fmt.Printf("\n--- %s ---\n", scenario.name)
		fmt.Printf("Description: %s\n", scenario.description)
		fmt.Printf("URL: %s\n", scenario.url)
		fmt.Printf("Requests: %d\n", scenario.requests)

		runScenario(smartClient, scenario.url, scenario.requests, scenario.delay)

		// Print status after each scenario
		smartClient.PrintStatus()

		// Wait between scenarios
		time.Sleep(3 * time.Second)
	}

	fmt.Println("\n=== Simulation Complete ===")
}

// runScenario executes a specific test scenario
func runScenario(client *SmartHTTPClient, url string, requests int, delay time.Duration) {
	var wg sync.WaitGroup

	for i := 0; i < requests; i++ {
		wg.Add(1)
		go func(requestID int) {
			defer wg.Done()

			start := time.Now()
			ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
			defer cancel()

			req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
			if err != nil {
				log.Printf("Request %d: Failed to create request: %v", requestID, err)
				return
			}

			resp, err := client.Do(req)
			duration := time.Since(start)

			if err != nil {
				log.Printf("Request %d: FAILED after %v - %v", requestID, duration, err)
				return
			}
			defer resp.Body.Close()

			log.Printf("Request %d: SUCCESS - Status: %d, Duration: %v",
				requestID, resp.StatusCode, duration)
		}(i)

		if i < requests-1 {
			time.Sleep(delay)
		}
	}

	wg.Wait()
}

// MonitoringDemo shows real-time monitoring capabilities
func MonitoringDemo() {
	fmt.Println("\n=== Real-time Monitoring Demo ===")

	smartClient := NewSmartHTTPClient()
	defer smartClient.Close()

	// Test URLs with different reliability characteristics
	urls := []string{
		"http://www.baidu.com",          // Usually reliable
		"http://httpbin.org/status/200", // Always returns 200
		"http://httpbin.org/status/500", // Always returns 500
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Start monitoring goroutine
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				fmt.Println("\n--- Monitoring Update ---")
				smartClient.PrintStatus()
			}
		}
	}()

	// Generate load
	var wg sync.WaitGroup

	for workerID := 0; workerID < 3; workerID++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			ticker := time.NewTicker(2 * time.Second)
			defer ticker.Stop()

			urlIndex := 0

			for {
				select {
				case <-ctx.Done():
					log.Printf("Monitoring worker %d stopped", id)
					return
				case <-ticker.C:
					url := urls[urlIndex%len(urls)]
					urlIndex++

					go func() {
						reqCtx, reqCancel := context.WithTimeout(context.Background(), 5*time.Second)
						defer reqCancel()

						req, err := http.NewRequestWithContext(reqCtx, "GET", url, nil)
						if err != nil {
							return
						}

						start := time.Now()
						resp, err := smartClient.Do(req)
						duration := time.Since(start)

						if err != nil {
							log.Printf("Worker %d: %s FAILED after %v: %v", id, url, duration, err)
						} else {
							resp.Body.Close()
							log.Printf("Worker %d: %s SUCCESS - Status: %d, Duration: %v",
								id, url, resp.StatusCode, duration)
						}
					}()
				}
			}
		}(workerID)
	}

	wg.Wait()

	fmt.Println("\n=== Final Monitoring Status ===")
	smartClient.PrintStatus()
}

// RunFailureSimulation is the entry point for failure simulation
func RunFailureSimulation() {
	if len(os.Args) > 1 && os.Args[1] == "simulate" {
		FailureSimulation()
	} else if len(os.Args) > 1 && os.Args[1] == "monitor" {
		MonitoringDemo()
	} else {
		fmt.Println("Usage:")
		fmt.Println("  go run . simulate  - Run failure simulation demo")
		fmt.Println("  go run . monitor   - Run real-time monitoring demo")
		fmt.Println("  go run .          - Run basic smart client demo")
		fmt.Println()

		// Run the basic demo by default
		// This will call the main() function from dns.go
	}
}
