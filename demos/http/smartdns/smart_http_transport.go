package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"sync/atomic"
	"time"
)

// SmartHTTPTransport implements intelligent HTTP transport with IP failover
type SmartHTTPTransport struct {
	resolver   *SmartResolver
	transport  *http.Transport
	roundRobin map[string]*int64 // Round-robin counter for each host
}

// NewSmartHTTPTransport creates a new smart HTTP transport
func NewSmartHTTPTransport() *SmartHTTPTransport {
	resolver := NewSmartResolver()

	transport := &http.Transport{
		DialContext:           nil, // Will be set in RoundTrip
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		DisableKeepAlives:     false,
	}

	return &SmartHTTPTransport{
		resolver:   resolver,
		transport:  transport,
		roundRobin: make(map[string]*int64),
	}
}

// Close cleans up resources
func (st *SmartHTTPTransport) Close() {
	if st.resolver != nil {
		st.resolver.Close()
	}
}

// RoundTrip implements http.RoundTripper interface
func (st *SmartHTTPTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.URL == nil {
		return nil, errors.New("request URL is nil")
	}

	host := req.URL.Hostname()
	if host == "" {
		return nil, errors.New("empty hostname in request")
	}

	// Check if it's already an IP address
	if net.ParseIP(host) != nil {
		// Direct IP request, use original transport
		return st.transport.RoundTrip(req)
	}

	// Resolve DNS and get healthy IPs
	ctx := req.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	// Perform DNS lookup to populate IP pool
	_, err := st.resolver.LookupIPAddr(ctx, host)
	if err != nil {
		return nil, fmt.Errorf("DNS lookup failed for %s: %w", host, err)
	}

	// Get healthy IPs
	healthyIPs := st.resolver.GetHealthyIPs(host)
	if len(healthyIPs) == 0 {
		return nil, fmt.Errorf("no healthy IPs available for host %s", host)
	}

	// Select IP using round-robin
	selectedIP := st.selectIP(host, healthyIPs)

	// Create transport with custom dialer for the selected IP
	transportCopy := st.transport.Clone()
	transportCopy.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		// Replace hostname with selected IP
		_, port, err := net.SplitHostPort(addr)
		if err != nil {
			// Default ports
			if req.URL.Scheme == "https" {
				port = "443"
			} else {
				port = "80"
			}
		}

		targetAddr := net.JoinHostPort(selectedIP, port)
		dialer := &net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}
		return dialer.DialContext(ctx, network, targetAddr)
	}

	// Execute request
	startTime := time.Now()
	resp, err := transportCopy.RoundTrip(req)
	duration := time.Since(startTime)

	// Update statistics
	success := err == nil
	statusCode := 0
	if resp != nil {
		statusCode = resp.StatusCode
	}

	st.resolver.UpdateIPStats(host, selectedIP, success, statusCode)

	// Log the request for debugging
	fmt.Printf(
		"Request to %s using IP %s: status=%d, duration=%v, success=%v\n",
		host, selectedIP, statusCode, duration, success,
	)

	// If request failed, try another IP if available
	if err != nil || (resp != nil && resp.StatusCode >= 500) {
		return st.retryWithAnotherIP(req, host, selectedIP, healthyIPs)
	}

	return resp, err
}

// selectIP selects an IP using round-robin algorithm
func (st *SmartHTTPTransport) selectIP(host string, ips []string) string {
	if len(ips) == 1 {
		return ips[0]
	}

	// Initialize counter if not exists
	if _, exists := st.roundRobin[host]; !exists {
		counter := int64(0)
		st.roundRobin[host] = &counter
	}

	// Get next IP using round-robin
	counter := st.roundRobin[host]
	index := atomic.AddInt64(counter, 1) % int64(len(ips))
	return ips[index]
}

// retryWithAnotherIP attempts to retry the request with a different IP
func (st *SmartHTTPTransport) retryWithAnotherIP(
	req *http.Request, host, failedIP string, availableIPs []string,
) (*http.Response, error) {
	// Find alternative IPs (excluding the failed one)
	var alternativeIPs []string
	for _, ip := range availableIPs {
		if ip != failedIP {
			alternativeIPs = append(alternativeIPs, ip)
		}
	}

	if len(alternativeIPs) == 0 {
		return nil, fmt.Errorf("no alternative IPs available for host %s", host)
	}

	// Try one alternative IP
	selectedIP := alternativeIPs[0]

	// Create transport with custom dialer for the alternative IP
	transportCopy := st.transport.Clone()
	transportCopy.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		_, port, err := net.SplitHostPort(addr)
		if err != nil {
			if req.URL.Scheme == "https" {
				port = "443"
			} else {
				port = "80"
			}
		}

		targetAddr := net.JoinHostPort(selectedIP, port)
		dialer := &net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}
		return dialer.DialContext(ctx, network, targetAddr)
	}

	// Execute retry request
	startTime := time.Now()
	resp, err := transportCopy.RoundTrip(req)
	duration := time.Since(startTime)

	// Update statistics for retry
	success := err == nil
	statusCode := 0
	if resp != nil {
		statusCode = resp.StatusCode
	}

	st.resolver.UpdateIPStats(host, selectedIP, success, statusCode)

	fmt.Printf(
		"Retry request to %s using IP %s: status=%d, duration=%v, success=%v\n",
		host, selectedIP, statusCode, duration, success,
	)

	return resp, err
}

// GetTransportStatus returns the current status of the transport
func (st *SmartHTTPTransport) GetTransportStatus() map[string]interface{} {
	return map[string]interface{}{
		"ip_pool_status": st.resolver.GetIPPoolStatus(),
		"round_robin_counters": func() map[string]int64 {
			result := make(map[string]int64)
			for host, counter := range st.roundRobin {
				result[host] = atomic.LoadInt64(counter)
			}
			return result
		}(),
	}
}

// SmartHTTPClient wraps http.Client with smart transport
type SmartHTTPClient struct {
	*http.Client
	transport *SmartHTTPTransport
}

// NewSmartHTTPClient creates a new smart HTTP client
func NewSmartHTTPClient() *SmartHTTPClient {
	transport := NewSmartHTTPTransport()

	client := &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second,
	}

	return &SmartHTTPClient{
		Client:    client,
		transport: transport,
	}
}

// Close cleans up resources
func (sc *SmartHTTPClient) Close() {
	if sc.transport != nil {
		sc.transport.Close()
	}
}

// GetStatus returns the current status of the client
func (sc *SmartHTTPClient) GetStatus() map[string]interface{} {
	if sc.transport == nil {
		return nil
	}
	return sc.transport.GetTransportStatus()
}

// PrintStatus prints the current status in a human-readable format
func (sc *SmartHTTPClient) PrintStatus() {
	status := sc.GetStatus()
	if status == nil {
		fmt.Println("No status available")
		return
	}

	fmt.Println("\n=== Smart HTTP Client Status ===")

	// Print IP pool status
	if ipPoolStatus, ok := status["ip_pool_status"].(map[string][]map[string]interface{}); ok {
		fmt.Println("\nIP Pool Status:")
		for host, ips := range ipPoolStatus {
			fmt.Printf("  Host: %s\n", host)
			for _, ipInfo := range ips {
				statusStr := "Unknown"
				if s, ok := ipInfo["status"].(IPStatus); ok {
					switch s {
					case IPStatusHealthy:
						statusStr = "Healthy"
					case IPStatusUnhealthy:
						statusStr = "Unhealthy"
					case IPStatusHalfOpen:
						statusStr = "Half-Open"
					}
				}

				fmt.Printf(
					"    IP: %s, Status: %s, Success: %v, Failure: %v, Rate: %.2f%%\n",
					ipInfo["ip"], statusStr,
					ipInfo["success_count"], ipInfo["failure_count"],
					ipInfo["success_rate"].(float64)*100,
				)
			}
		}
	}

	// Print round-robin counters
	if rrCounters, ok := status["round_robin_counters"].(map[string]int64); ok {
		fmt.Println("\nRound-Robin Counters:")
		for host, counter := range rrCounters {
			fmt.Printf("  %s: %d\n", host, counter)
		}
	}

	fmt.Println("================================\n")
}
