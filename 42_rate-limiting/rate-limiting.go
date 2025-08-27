/*
=============================================================================
                        🚦 GO RATE LIMITING TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Rate limiting controls the rate at which operations are performed.
Go provides several mechanisms including time.Ticker, token bucket,
and sliding window algorithms for implementing rate limiting.

🔑 KEY FEATURES:
• Control request/operation frequency
• Prevent system overload
• Implement backpressure
• Fair resource allocation

💡 REAL-WORLD ANALOGY:
Rate Limiting = Traffic Light System
- Ticker = Regular traffic light intervals
- Token bucket = Toll booth with limited tokens
- Sliding window = Traffic flow monitoring
- Backpressure = Traffic congestion management

🎯 WHY USE RATE LIMITING?
• Protect APIs from abuse
• Ensure fair resource usage
• Prevent system overload
• Comply with external API limits

=============================================================================
*/

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 🎯 BASIC RATE LIMITER WITH TICKER
func basicRateLimiter() {
	fmt.Println("🎯 Basic Rate Limiter (Ticker)")
	fmt.Println("==============================")

	// Allow 1 operation per second
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	requests := []string{"req1", "req2", "req3", "req4", "req5"}

	fmt.Println("Processing requests at 1 per second:")
	for _, req := range requests {
		<-ticker.C // Wait for ticker
		fmt.Printf("⏰ Processing %s at %s\n", req, time.Now().Format("15:04:05"))
	}
}

// 🪣 TOKEN BUCKET RATE LIMITER
type TokenBucket struct {
	tokens   chan struct{}
	ticker   *time.Ticker
	capacity int
	rate     time.Duration
	quit     chan bool
}

func NewTokenBucket(capacity int, rate time.Duration) *TokenBucket {
	tb := &TokenBucket{
		tokens:   make(chan struct{}, capacity),
		ticker:   time.NewTicker(rate),
		capacity: capacity,
		rate:     rate,
		quit:     make(chan bool),
	}

	// Fill bucket initially
	for i := 0; i < capacity; i++ {
		tb.tokens <- struct{}{}
	}

	// Start token refill goroutine
	go tb.refill()

	return tb
}

func (tb *TokenBucket) refill() {
	for {
		select {
		case <-tb.ticker.C:
			select {
			case tb.tokens <- struct{}{}:
				// Token added
			default:
				// Bucket full, skip
			}
		case <-tb.quit:
			return
		}
	}
}

func (tb *TokenBucket) Allow() bool {
	select {
	case <-tb.tokens:
		return true
	default:
		return false
	}
}

func (tb *TokenBucket) Wait() {
	<-tb.tokens
}

func (tb *TokenBucket) Stop() {
	tb.ticker.Stop()
	close(tb.quit)
}

// 🌊 SLIDING WINDOW RATE LIMITER
type SlidingWindow struct {
	mu       sync.Mutex
	requests []time.Time
	limit    int
	window   time.Duration
}

func NewSlidingWindow(limit int, window time.Duration) *SlidingWindow {
	return &SlidingWindow{
		requests: make([]time.Time, 0),
		limit:    limit,
		window:   window,
	}
}

func (sw *SlidingWindow) Allow() bool {
	sw.mu.Lock()
	defer sw.mu.Unlock()

	now := time.Now()
	cutoff := now.Add(-sw.window)

	// Remove old requests
	validRequests := sw.requests[:0]
	for _, req := range sw.requests {
		if req.After(cutoff) {
			validRequests = append(validRequests, req)
		}
	}
	sw.requests = validRequests

	// Check if we can allow this request
	if len(sw.requests) < sw.limit {
		sw.requests = append(sw.requests, now)
		return true
	}

	return false
}

func (sw *SlidingWindow) RequestCount() int {
	sw.mu.Lock()
	defer sw.mu.Unlock()
	return len(sw.requests)
}

// 🎛️ ADAPTIVE RATE LIMITER
type AdaptiveRateLimiter struct {
	mu           sync.Mutex
	currentRate  time.Duration
	minRate      time.Duration
	maxRate      time.Duration
	successCount int
	errorCount   int
	lastAdjust   time.Time
	ticker       *time.Ticker
	tokens       chan struct{}
}

func NewAdaptiveRateLimiter(initialRate, minRate, maxRate time.Duration) *AdaptiveRateLimiter {
	arl := &AdaptiveRateLimiter{
		currentRate: initialRate,
		minRate:     minRate,
		maxRate:     maxRate,
		lastAdjust:  time.Now(),
		tokens:      make(chan struct{}, 1),
	}
	
	arl.tokens <- struct{}{} // Initial token
	arl.startTicker()
	
	return arl
}

func (arl *AdaptiveRateLimiter) startTicker() {
	if arl.ticker != nil {
		arl.ticker.Stop()
	}
	
	arl.ticker = time.NewTicker(arl.currentRate)
	go func() {
		for range arl.ticker.C {
			select {
			case arl.tokens <- struct{}{}:
			default:
			}
		}
	}()
}

func (arl *AdaptiveRateLimiter) Allow() bool {
	select {
	case <-arl.tokens:
		return true
	default:
		return false
	}
}

func (arl *AdaptiveRateLimiter) RecordSuccess() {
	arl.mu.Lock()
	defer arl.mu.Unlock()
	arl.successCount++
	arl.adjustRate()
}

func (arl *AdaptiveRateLimiter) RecordError() {
	arl.mu.Lock()
	defer arl.mu.Unlock()
	arl.errorCount++
	arl.adjustRate()
}

func (arl *AdaptiveRateLimiter) adjustRate() {
	now := time.Now()
	if now.Sub(arl.lastAdjust) < 5*time.Second {
		return // Don't adjust too frequently
	}

	total := arl.successCount + arl.errorCount
	if total < 10 {
		return // Need more data
	}

	errorRate := float64(arl.errorCount) / float64(total)
	
	if errorRate > 0.1 { // More than 10% errors, slow down
		newRate := time.Duration(float64(arl.currentRate) * 1.5)
		if newRate <= arl.maxRate {
			arl.currentRate = newRate
			arl.startTicker()
			fmt.Printf("🐌 Slowing down to %v (error rate: %.1f%%)\n", arl.currentRate, errorRate*100)
		}
	} else if errorRate < 0.05 { // Less than 5% errors, speed up
		newRate := time.Duration(float64(arl.currentRate) * 0.8)
		if newRate >= arl.minRate {
			arl.currentRate = newRate
			arl.startTicker()
			fmt.Printf("🚀 Speeding up to %v (error rate: %.1f%%)\n", arl.currentRate, errorRate*100)
		}
	}

	// Reset counters
	arl.successCount = 0
	arl.errorCount = 0
	arl.lastAdjust = now
}

// 🌐 HTTP RATE LIMITER EXAMPLE
type HTTPRateLimiter struct {
	limiters map[string]*TokenBucket
	mu       sync.RWMutex
	capacity int
	rate     time.Duration
}

func NewHTTPRateLimiter(capacity int, rate time.Duration) *HTTPRateLimiter {
	return &HTTPRateLimiter{
		limiters: make(map[string]*TokenBucket),
		capacity: capacity,
		rate:     rate,
	}
}

func (hrl *HTTPRateLimiter) Allow(clientID string) bool {
	hrl.mu.RLock()
	limiter, exists := hrl.limiters[clientID]
	hrl.mu.RUnlock()

	if !exists {
		hrl.mu.Lock()
		// Double-check after acquiring write lock
		if limiter, exists = hrl.limiters[clientID]; !exists {
			limiter = NewTokenBucket(hrl.capacity, hrl.rate)
			hrl.limiters[clientID] = limiter
		}
		hrl.mu.Unlock()
	}

	return limiter.Allow()
}

func main() {
	fmt.Println("🚦 RATE LIMITING TUTORIAL")
	fmt.Println("=========================")

	// 🎯 DEMO 1: Basic Rate Limiter
	basicRateLimiter()

	// 🎯 DEMO 2: Token Bucket Rate Limiter
	fmt.Println("\n🪣 Token Bucket Rate Limiter")
	fmt.Println("============================")

	bucket := NewTokenBucket(3, 500*time.Millisecond) // 3 tokens, refill every 500ms
	defer bucket.Stop()

	fmt.Println("Token bucket (capacity: 3, refill: 500ms):")
	
	// Burst of requests
	for i := 1; i <= 8; i++ {
		if bucket.Allow() {
			fmt.Printf("✅ Request %d allowed at %s\n", i, time.Now().Format("15:04:05.000"))
		} else {
			fmt.Printf("❌ Request %d denied at %s\n", i, time.Now().Format("15:04:05.000"))
		}
		time.Sleep(200 * time.Millisecond)
	}

	// 🎯 DEMO 3: Sliding Window Rate Limiter
	fmt.Println("\n🌊 Sliding Window Rate Limiter")
	fmt.Println("==============================")

	window := NewSlidingWindow(5, 2*time.Second) // 5 requests per 2 seconds

	fmt.Println("Sliding window (5 requests per 2 seconds):")
	
	for i := 1; i <= 10; i++ {
		if window.Allow() {
			fmt.Printf("✅ Request %d allowed (count: %d) at %s\n", 
				i, window.RequestCount(), time.Now().Format("15:04:05.000"))
		} else {
			fmt.Printf("❌ Request %d denied (count: %d) at %s\n", 
				i, window.RequestCount(), time.Now().Format("15:04:05.000"))
		}
		time.Sleep(300 * time.Millisecond)
	}

	// 🎯 DEMO 4: Adaptive Rate Limiter
	fmt.Println("\n🎛️ Adaptive Rate Limiter")
	fmt.Println("========================")

	adaptive := NewAdaptiveRateLimiter(
		1*time.Second,   // initial rate
		200*time.Millisecond, // min rate (fastest)
		3*time.Second,   // max rate (slowest)
	)

	fmt.Println("Adaptive rate limiter (adjusts based on success/error rate):")
	
	// Simulate requests with varying success rates
	for i := 1; i <= 30; i++ {
		if adaptive.Allow() {
			// Simulate success/error
			if i < 10 || i > 20 {
				adaptive.RecordSuccess()
				fmt.Printf("✅ Request %d succeeded\n", i)
			} else {
				adaptive.RecordError()
				fmt.Printf("❌ Request %d failed\n", i)
			}
		} else {
			fmt.Printf("⏳ Request %d rate limited\n", i)
		}
		time.Sleep(100 * time.Millisecond)
	}

	// 🎯 DEMO 5: HTTP Rate Limiter (per client)
	fmt.Println("\n🌐 HTTP Rate Limiter (Per Client)")
	fmt.Println("=================================")

	httpLimiter := NewHTTPRateLimiter(2, 1*time.Second) // 2 requests, refill every second

	clients := []string{"client1", "client2", "client1", "client3", "client1", "client2"}
	
	fmt.Println("Per-client rate limiting:")
	for i, client := range clients {
		if httpLimiter.Allow(client) {
			fmt.Printf("✅ Request %d from %s allowed\n", i+1, client)
		} else {
			fmt.Printf("❌ Request %d from %s denied\n", i+1, client)
		}
		time.Sleep(200 * time.Millisecond)
	}

	// 🎯 DEMO 6: Rate Limiting with Context
	fmt.Println("\n⏰ Rate Limiting with Context")
	fmt.Println("=============================")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rateLimitedWork := func(ctx context.Context, id int) error {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()

		select {
		case <-ticker.C:
			fmt.Printf("🔄 Work %d completed\n", id)
			return nil
		case <-ctx.Done():
			fmt.Printf("⏰ Work %d cancelled: %v\n", id, ctx.Err())
			return ctx.Err()
		}
	}

	fmt.Println("Rate limited work with context timeout:")
	var wg sync.WaitGroup
	for i := 1; i <= 8; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			rateLimitedWork(ctx, id)
		}(i)
	}
	wg.Wait()

	// 🎯 DEMO 7: Backpressure Example
	fmt.Println("\n🔙 Backpressure Example")
	fmt.Println("=======================")

	jobs := make(chan int, 3) // Small buffer to demonstrate backpressure
	results := make(chan int, 3)

	// Worker with rate limiting
	go func() {
		ticker := time.NewTicker(800 * time.Millisecond)
		defer ticker.Stop()

		for job := range jobs {
			<-ticker.C // Rate limit
			result := job * 2
			results <- result
			fmt.Printf("🔄 Processed job %d -> %d\n", job, result)
		}
		close(results)
	}()

	// Producer with backpressure handling
	go func() {
		defer close(jobs)
		for i := 1; i <= 8; i++ {
			select {
			case jobs <- i:
				fmt.Printf("📤 Sent job %d\n", i)
			default:
				fmt.Printf("🚫 Job %d blocked (backpressure)\n", i)
				// In real scenario, might retry later or drop
				time.Sleep(200 * time.Millisecond)
				jobs <- i // Retry
				fmt.Printf("📤 Sent job %d (retry)\n", i)
			}
		}
	}()

	// Collect results
	for result := range results {
		fmt.Printf("📥 Received result: %d\n", result)
	}

	fmt.Println("\n✨ All rate limiting demos completed!")
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

🚦 RATE LIMITING ALGORITHMS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Token Bucket                                                         │
│ • Fixed capacity bucket with tokens                                     │
│ • Tokens refilled at constant rate                                      │
│ • Allows bursts up to bucket capacity                                   │
│ • Good for: API rate limiting, traffic shaping                          │
│                                                                         │
│ // Sliding Window                                                       │
│ • Track requests in time window                                         │
│ • More accurate than fixed window                                       │
│ • Higher memory usage                                                   │
│ • Good for: Precise rate limiting, analytics                            │
│                                                                         │
│ // Fixed Window                                                         │
│ • Count requests in fixed time periods                                  │
│ • Simple but can allow bursts at boundaries                             │
│ • Low memory usage                                                      │
│ • Good for: Simple rate limiting, logging                               │
└─────────────────────────────────────────────────────────────────────────┘

🪣 TOKEN BUCKET IMPLEMENTATION:
┌─────────────────────────────────────────────────────────────────────────┐
│ type TokenBucket struct {                                               │
│     tokens   chan struct{}  // Token storage                           │
│     ticker   *time.Ticker   // Token refill timer                      │
│     capacity int            // Maximum tokens                          │
│     rate     time.Duration  // Refill rate                             │
│ }                                                                       │
│                                                                         │
│ func (tb *TokenBucket) Allow() bool {                                   │
│     select {                                                            │
│     case <-tb.tokens:                                                   │
│         return true                                                     │
│     default:                                                            │
│         return false                                                    │
│     }                                                                   │
│ }                                                                       │
│                                                                         │
│ func (tb *TokenBucket) Wait() {                                         │
│     <-tb.tokens  // Block until token available                        │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

⏰ TIMING PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Simple ticker rate limiting                                          │
│ ticker := time.NewTicker(1 * time.Second)                               │
│ defer ticker.Stop()                                                     │
│ for range ticker.C {                                                    │
│     // Process one item per second                                      │
│ }                                                                       │
│                                                                         │
│ // Rate limiting with timeout                                           │
│ ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) │
│ defer cancel()                                                          │
│                                                                         │
│ ticker := time.NewTicker(500 * time.Millisecond)                       │
│ defer ticker.Stop()                                                     │
│                                                                         │
│ for {                                                                   │
│     select {                                                            │
│     case <-ticker.C:                                                    │
│         // Process item                                                 │
│     case <-ctx.Done():                                                  │
│         return ctx.Err()                                                │
│     }                                                                   │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

🌊 BACKPRESSURE HANDLING:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Non-blocking send with backpressure                                  │
│ select {                                                                │
│ case jobQueue <- job:                                                   │
│     // Job sent successfully                                            │
│ default:                                                                │
│     // Queue full, handle backpressure                                  │
│     // Options: drop, retry later, block, return error                  │
│ }                                                                       │
│                                                                         │
│ // Buffered channel for burst handling                                  │
│ jobs := make(chan Job, bufferSize)                                      │
│                                                                         │
│ // Rate limited worker                                                  │
│ ticker := time.NewTicker(rate)                                          │
│ for {                                                                   │
│     select {                                                            │
│     case job := <-jobs:                                                 │
│         <-ticker.C  // Wait for rate limit                              │
│         processJob(job)                                                 │
│     }                                                                   │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

📊 MONITORING AND METRICS:
┌─────────────────────────────────────────────────────────────────────────┐
│ type RateLimiterMetrics struct {                                        │
│     AllowedRequests  int64                                              │
│     DeniedRequests   int64                                              │
│     CurrentTokens    int                                                │
│     QueueLength      int                                                │
│     AverageWaitTime  time.Duration                                      │
│ }                                                                       │
│                                                                         │
│ // Metrics collection                                                   │
│ func (rl *RateLimiter) GetMetrics() RateLimiterMetrics {                │
│     return RateLimiterMetrics{                                          │
│         AllowedRequests: atomic.LoadInt64(&rl.allowed),                 │
│         DeniedRequests:  atomic.LoadInt64(&rl.denied),                  │
│         CurrentTokens:   len(rl.tokens),                                │
│         QueueLength:     len(rl.queue),                                 │
│     }                                                                   │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

💡 BEST PRACTICES:
• Choose algorithm based on requirements (burst vs smooth)
• Monitor rate limiter performance and adjust parameters
• Implement graceful degradation when limits are hit
• Use context for timeout and cancellation
• Consider per-client vs global rate limiting
• Log rate limiting events for analysis
• Test rate limiters under load

🚨 COMMON MISTAKES:
❌ Not handling backpressure properly
❌ Using fixed window without considering burst behavior
❌ Not monitoring rate limiter effectiveness
❌ Blocking operations in rate-limited paths
❌ Not considering memory usage of sliding windows
❌ Forgetting to stop tickers (memory leaks)

⚡ PERFORMANCE CONSIDERATIONS:
• Token bucket: O(1) for allow/deny decisions
• Sliding window: O(n) where n is requests in window
• Memory usage varies by algorithm
• Consider using atomic operations for counters
• Profile rate limiter overhead in hot paths

🎯 REAL-WORLD APPLICATIONS:
• API rate limiting (per user/IP)
• Database connection throttling
• External service call limiting
• Resource usage control
• Traffic shaping and QoS
• Batch processing rate control
• Circuit breaker patterns

🔧 ADVANCED PATTERNS:
• Hierarchical rate limiting (global + per-user)
• Adaptive rate limiting based on system load
• Distributed rate limiting with Redis
• Rate limiting with priority queues
• Exponential backoff with rate limiting

=============================================================================
*/