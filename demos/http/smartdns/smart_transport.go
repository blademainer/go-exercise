package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

// IPStatus represents the status of an IP address
type IPStatus int32

const (
	IPStatusHealthy IPStatus = iota
	IPStatusUnhealthy
	IPStatusHalfOpen
)

// IPInfo holds information about an IP address
type IPInfo struct {
	IP               string
	Status           IPStatus
	SuccessCount     int64
	FailureCount     int64
	LastSuccess      time.Time
	LastFailure      time.Time
	ConsecutiveFails int64
	mu               sync.RWMutex
}

// UpdateSuccess updates success statistics
func (ip *IPInfo) UpdateSuccess() {
	ip.mu.Lock()
	defer ip.mu.Unlock()

	atomic.AddInt64(&ip.SuccessCount, 1)
	atomic.StoreInt64(&ip.ConsecutiveFails, 0)
	ip.LastSuccess = time.Now()

	// If IP was unhealthy or half-open, mark as healthy
	if ip.Status != IPStatusHealthy {
		ip.Status = IPStatusHealthy
		fmt.Printf("IP %s recovered to healthy status\n", ip.IP)
	}
}

// UpdateFailure updates failure statistics
func (ip *IPInfo) UpdateFailure() {
	ip.mu.Lock()
	defer ip.mu.Unlock()

	atomic.AddInt64(&ip.FailureCount, 1)
	atomic.AddInt64(&ip.ConsecutiveFails, 1)
	ip.LastFailure = time.Now()
}

// GetSuccessRate calculates the success rate
func (ip *IPInfo) GetSuccessRate() float64 {
	total := atomic.LoadInt64(&ip.SuccessCount) + atomic.LoadInt64(&ip.FailureCount)
	if total == 0 {
		return 1.0 // No data, assume healthy
	}
	return float64(atomic.LoadInt64(&ip.SuccessCount)) / float64(total)
}

// ShouldCircuitBreak determines if this IP should be circuit broken
func (ip *IPInfo) ShouldCircuitBreak(threshold int64, minRequests int64) bool {
	ip.mu.RLock()
	defer ip.mu.RUnlock()

	total := atomic.LoadInt64(&ip.SuccessCount) + atomic.LoadInt64(&ip.FailureCount)
	consecutiveFails := atomic.LoadInt64(&ip.ConsecutiveFails)

	// Need minimum requests to make a decision
	if total < minRequests {
		return false
	}

	return consecutiveFails >= threshold
}

// SmartResolver implements custom DNS resolution with IP health tracking
type SmartResolver struct {
	resolver *net.Resolver
	ipPool   map[string][]*IPInfo
	mu       sync.RWMutex

	// Circuit breaker configuration
	failureThreshold int64
	minRequests      int64
	recoveryTimeout  time.Duration

	// Health check configuration
	healthCheckInterval time.Duration
	healthCheckTimeout  time.Duration

	// Context for stopping background tasks
	ctx    context.Context
	cancel context.CancelFunc
}

// NewSmartResolver creates a new smart resolver
func NewSmartResolver() *SmartResolver {
	ctx, cancel := context.WithCancel(context.Background())

	sr := &SmartResolver{
		resolver:            &net.Resolver{},
		ipPool:              make(map[string][]*IPInfo),
		failureThreshold:    3, // Circuit break after 3 consecutive failures
		minRequests:         5, // Need at least 5 requests to make decisions
		recoveryTimeout:     30 * time.Second,
		healthCheckInterval: 10 * time.Second,
		healthCheckTimeout:  5 * time.Second,
		ctx:                 ctx,
		cancel:              cancel,
	}

	// Start background health checker
	go sr.healthChecker()

	return sr
}

// Close stops background tasks
func (sr *SmartResolver) Close() {
	sr.cancel()
}

// LookupIPAddr performs DNS lookup and manages IP pool
func (sr *SmartResolver) LookupIPAddr(ctx context.Context, host string) ([]net.IPAddr, error) {
	// Perform actual DNS lookup
	ips, err := sr.resolver.LookupIPAddr(ctx, host)
	if err != nil {
		return nil, err
	}

	sr.mu.Lock()
	defer sr.mu.Unlock()

	// Initialize or update IP pool for this host
	if _, exists := sr.ipPool[host]; !exists {
		sr.ipPool[host] = make([]*IPInfo, 0, len(ips))
	}

	// Create IPInfo for new IPs
	existingIPs := make(map[string]*IPInfo)
	for _, ipInfo := range sr.ipPool[host] {
		existingIPs[ipInfo.IP] = ipInfo
	}

	var newIPInfos []*IPInfo
	for _, ip := range ips {
		ipStr := ip.IP.String()
		if existing, found := existingIPs[ipStr]; found {
			newIPInfos = append(newIPInfos, existing)
		} else {
			newIPInfo := &IPInfo{
				IP:          ipStr,
				Status:      IPStatusHealthy,
				LastSuccess: time.Now(),
			}
			newIPInfos = append(newIPInfos, newIPInfo)
			fmt.Printf("Added new IP %s for host %s\n", ipStr, host)
		}
	}

	sr.ipPool[host] = newIPInfos

	return ips, nil
}

// GetHealthyIPs returns healthy IPs for a host
func (sr *SmartResolver) GetHealthyIPs(host string) []string {
	sr.mu.RLock()
	defer sr.mu.RUnlock()

	ipInfos, exists := sr.ipPool[host]
	if !exists {
		return nil
	}

	var healthyIPs []string
	var halfOpenIPs []string

	for _, ipInfo := range ipInfos {
		ipInfo.mu.RLock()
		status := ipInfo.Status
		ipInfo.mu.RUnlock()

		switch status {
		case IPStatusHealthy:
			healthyIPs = append(healthyIPs, ipInfo.IP)
		case IPStatusHalfOpen:
			halfOpenIPs = append(halfOpenIPs, ipInfo.IP)
		}
	}

	// If no healthy IPs, try half-open ones
	if len(healthyIPs) == 0 && len(halfOpenIPs) > 0 {
		return halfOpenIPs[:1] // Only try one half-open IP at a time
	}

	return healthyIPs
}

// UpdateIPStats updates statistics for an IP after a request
func (sr *SmartResolver) UpdateIPStats(host, ip string, success bool, statusCode int) {
	sr.mu.RLock()
	ipInfos, exists := sr.ipPool[host]
	sr.mu.RUnlock()

	if !exists {
		return
	}

	for _, ipInfo := range ipInfos {
		if ipInfo.IP == ip {
			if success && statusCode >= 200 && statusCode < 400 {
				ipInfo.UpdateSuccess()
			} else {
				ipInfo.UpdateFailure()

				// Check if should circuit break
				if ipInfo.ShouldCircuitBreak(sr.failureThreshold, sr.minRequests) {
					ipInfo.mu.Lock()
					if ipInfo.Status == IPStatusHealthy {
						ipInfo.Status = IPStatusUnhealthy
						fmt.Printf("Circuit breaker triggered for IP %s, marked as unhealthy\n", ip)
					}
					ipInfo.mu.Unlock()
				}
			}
			break
		}
	}
}

// healthChecker runs periodic health checks and recovery attempts
func (sr *SmartResolver) healthChecker() {
	ticker := time.NewTicker(sr.healthCheckInterval)
	defer ticker.Stop()

	for {
		select {
		case <-sr.ctx.Done():
			return
		case <-ticker.C:
			sr.performHealthChecks()
		}
	}
}

// performHealthChecks checks unhealthy IPs for recovery
func (sr *SmartResolver) performHealthChecks() {
	sr.mu.RLock()
	hostsToCheck := make(map[string][]*IPInfo)
	for host, ipInfos := range sr.ipPool {
		var unhealthyIPs []*IPInfo
		for _, ipInfo := range ipInfos {
			ipInfo.mu.RLock()
			if ipInfo.Status == IPStatusUnhealthy &&
				time.Since(ipInfo.LastFailure) >= sr.recoveryTimeout {
				unhealthyIPs = append(unhealthyIPs, ipInfo)
			}
			ipInfo.mu.RUnlock()
		}
		if len(unhealthyIPs) > 0 {
			hostsToCheck[host] = unhealthyIPs
		}
	}
	sr.mu.RUnlock()

	// Perform health checks for unhealthy IPs
	for host, ipInfos := range hostsToCheck {
		for _, ipInfo := range ipInfos {
			go sr.checkIPHealth(host, ipInfo)
		}
	}
}

// checkIPHealth performs a health check on a specific IP
func (sr *SmartResolver) checkIPHealth(host string, ipInfo *IPInfo) {
	healthCtx, cancel := context.WithTimeout(context.Background(), sr.healthCheckTimeout)
	defer cancel()

	// Create a simple HTTP client with this specific IP
	transport := &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			dialer := &net.Dialer{}
			// Replace host with IP for health check
			_, port, err := net.SplitHostPort(addr)
			if err != nil {
				port = "80" // Default port
			}
			return dialer.DialContext(ctx, network, net.JoinHostPort(ipInfo.IP, port))
		},
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   sr.healthCheckTimeout,
	}

	// Perform health check request
	req, err := http.NewRequestWithContext(healthCtx, "GET", "http://"+host, nil)
	if err != nil {
		return
	}

	resp, err := client.Do(req)
	if err == nil && resp != nil {
		resp.Body.Close()
		if resp.StatusCode >= 200 && resp.StatusCode < 400 {
			// Health check passed, move to half-open state
			ipInfo.mu.Lock()
			if ipInfo.Status == IPStatusUnhealthy {
				ipInfo.Status = IPStatusHalfOpen
				fmt.Printf("IP %s moved to half-open state\n", ipInfo.IP)
			}
			ipInfo.mu.Unlock()
		}
	}
}

// GetIPPoolStatus returns current status of all IPs
func (sr *SmartResolver) GetIPPoolStatus() map[string][]map[string]interface{} {
	sr.mu.RLock()
	defer sr.mu.RUnlock()

	result := make(map[string][]map[string]interface{})

	for host, ipInfos := range sr.ipPool {
		var hostStatus []map[string]interface{}
		for _, ipInfo := range ipInfos {
			ipInfo.mu.RLock()
			status := map[string]interface{}{
				"ip":                ipInfo.IP,
				"status":            ipInfo.Status,
				"success_count":     atomic.LoadInt64(&ipInfo.SuccessCount),
				"failure_count":     atomic.LoadInt64(&ipInfo.FailureCount),
				"success_rate":      ipInfo.GetSuccessRate(),
				"consecutive_fails": atomic.LoadInt64(&ipInfo.ConsecutiveFails),
				"last_success":      ipInfo.LastSuccess,
				"last_failure":      ipInfo.LastFailure,
			}
			ipInfo.mu.RUnlock()
			hostStatus = append(hostStatus, status)
		}
		result[host] = hostStatus
	}

	return result
}
