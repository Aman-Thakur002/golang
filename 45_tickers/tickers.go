/*
=============================================================================
                           🎵 GO TICKERS TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Tickers in Go provide repeated execution at regular intervals.
Unlike timers which fire once, tickers continue firing until stopped.

🔑 KEY FEATURES:
• Repeated execution at fixed intervals
• Channel-based notifications
• Stoppable and resource-managed
• Non-blocking tick delivery
• Integration with select statements

💡 REAL-WORLD ANALOGY:
Ticker = Metronome
- Regular beats = Tick intervals
- Continuous operation = Keeps ticking until stopped
- Channel = Sound of each beat
- Stop = Turn off metronome
- Select = Listen for beats among other sounds

🎯 WHY USE TICKERS?
• Periodic tasks and maintenance
• Heartbeat mechanisms
• Regular data collection
• Animation and UI updates

=============================================================================
*/

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("🎵 TICKERS TUTORIAL")
	fmt.Println("===================")

	// 🎯 DEMO 1: Basic Ticker Usage
	fmt.Println("\n🎯 DEMO 1: Basic Ticker")
	fmt.Println("=======================")

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop() // Always stop tickers to prevent resource leaks

	fmt.Println("Ticker firing every 500ms for 3 seconds:")
	start := time.Now()

	for {
		select {
		case t := <-ticker.C:
			elapsed := time.Since(start)
			fmt.Printf("⏰ Tick at %s (elapsed: %v)\n", 
				t.Format("15:04:05.000"), elapsed.Round(time.Millisecond))
			
			if elapsed > 3*time.Second {
				fmt.Println("✅ Stopping ticker after 3 seconds")
				return
			}
		}
	}
}

func demoPeriodicTasks() {
	fmt.Println("\n🎯 DEMO 2: Periodic Tasks")
	fmt.Println("=========================")

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	counter := 0
	maxTicks := 5

	fmt.Printf("Running periodic task %d times:\n", maxTicks)

	for {
		select {
		case <-ticker.C:
			counter++
			fmt.Printf("📋 Task execution #%d at %s\n", 
				counter, time.Now().Format("15:04:05"))
			
			// Simulate some work
			time.Sleep(200 * time.Millisecond)
			
			if counter >= maxTicks {
				fmt.Println("✅ All periodic tasks completed")
				return
			}
		}
	}
}

func demoMultipleTickers() {
	fmt.Println("\n🎯 DEMO 3: Multiple Tickers")
	fmt.Println("===========================")

	fastTicker := time.NewTicker(300 * time.Millisecond)
	slowTicker := time.NewTicker(1 * time.Second)
	
	defer fastTicker.Stop()
	defer slowTicker.Stop()

	timeout := time.After(4 * time.Second)
	
	fmt.Println("Running fast (300ms) and slow (1s) tickers:")

	for {
		select {
		case <-fastTicker.C:
			fmt.Printf("🏃 Fast tick at %s\n", time.Now().Format("15:04:05.000"))
		case <-slowTicker.C:
			fmt.Printf("🐌 Slow tick at %s\n", time.Now().Format("15:04:05.000"))
		case <-timeout:
			fmt.Println("⏰ Timeout reached, stopping tickers")
			return
		}
	}
}

func demoTickerWithContext() {
	fmt.Println("\n🎯 DEMO 4: Ticker with Context")
	fmt.Println("==============================")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	fmt.Println("Ticker with context cancellation:")

	for {
		select {
		case <-ticker.C:
			fmt.Printf("🎯 Context tick at %s\n", time.Now().Format("15:04:05.000"))
		case <-ctx.Done():
			fmt.Printf("❌ Context cancelled: %v\n", ctx.Err())
			return
		}
	}
}

func demoHeartbeat() {
	fmt.Println("\n🎯 DEMO 5: Heartbeat Pattern")
	fmt.Println("============================")

	type Service struct {
		name      string
		heartbeat *time.Ticker
		quit      chan bool
		wg        sync.WaitGroup
	}

	NewService := func(name string, interval time.Duration) *Service {
		return &Service{
			name:      name,
			heartbeat: time.NewTicker(interval),
			quit:      make(chan bool),
		}
	}

	Start := func(s *Service) {
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			defer s.heartbeat.Stop()

			fmt.Printf("💓 Service %s started with heartbeat\n", s.name)

			for {
				select {
				case <-s.heartbeat.C:
					fmt.Printf("💓 %s heartbeat at %s\n", 
						s.name, time.Now().Format("15:04:05"))
				case <-s.quit:
					fmt.Printf("🛑 Service %s stopping\n", s.name)
					return
				}
			}
		}()
	}

	Stop := func(s *Service) {
		close(s.quit)
		s.wg.Wait()
		fmt.Printf("✅ Service %s stopped\n", s.name)
	}

	// Create and start services
	service1 := NewService("Database", 1*time.Second)
	service2 := NewService("Cache", 2*time.Second)

	Start(service1)
	Start(service2)

	// Let them run for a while
	time.Sleep(5 * time.Second)

	// Stop services
	Stop(service1)
	Stop(service2)
}

func demoDataCollection() {
	fmt.Println("\n🎯 DEMO 6: Data Collection")
	fmt.Println("==========================")

	type Metrics struct {
		mu       sync.Mutex
		requests int
		errors   int
	}

	metrics := &Metrics{}

	// Simulate incoming requests
	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()

		for i := 0; i < 30; i++ {
			<-ticker.C
			metrics.mu.Lock()
			metrics.requests++
			if i%7 == 0 { // Simulate some errors
				metrics.errors++
			}
			metrics.mu.Unlock()
		}
	}()

	// Collect metrics every second
	collector := time.NewTicker(1 * time.Second)
	defer collector.Stop()

	timeout := time.After(4 * time.Second)

	fmt.Println("Collecting metrics every second:")

	for {
		select {
		case <-collector.C:
			metrics.mu.Lock()
			requests := metrics.requests
			errors := metrics.errors
			metrics.mu.Unlock()

			errorRate := float64(errors) / float64(requests) * 100
			fmt.Printf("📊 Requests: %d, Errors: %d (%.1f%%)\n", 
				requests, errors, errorRate)

		case <-timeout:
			fmt.Println("📊 Data collection completed")
			return
		}
	}
}

func demoRateLimiting() {
	fmt.Println("\n🎯 DEMO 7: Rate Limiting with Ticker")
	fmt.Println("====================================")

	// Rate limiter using ticker
	rateLimiter := time.NewTicker(200 * time.Millisecond) // 5 requests per second
	defer rateLimiter.Stop()

	requests := []string{"req1", "req2", "req3", "req4", "req5", "req6", "req7"}

	fmt.Println("Processing requests with rate limiting (5/sec):")

	for i, req := range requests {
		<-rateLimiter.C // Wait for rate limiter
		fmt.Printf("🚦 Processing %s at %s\n", req, time.Now().Format("15:04:05.000"))
		
		// Simulate request processing
		go func(id string, num int) {
			time.Sleep(50 * time.Millisecond)
			fmt.Printf("   ✅ %s completed (%d/%d)\n", id, num+1, len(requests))
		}(req, i)
	}

	time.Sleep(500 * time.Millisecond) // Wait for last requests to complete
}

func demoTickerPool() {
	fmt.Println("\n🎯 DEMO 8: Ticker Pool Pattern")
	fmt.Println("==============================")

	type TickerPool struct {
		tickers map[string]*time.Ticker
		mu      sync.RWMutex
	}

	NewTickerPool := func() *TickerPool {
		return &TickerPool{
			tickers: make(map[string]*time.Ticker),
		}
	}

	AddTicker := func(tp *TickerPool, name string, interval time.Duration, fn func()) {
		tp.mu.Lock()
		defer tp.mu.Unlock()

		if _, exists := tp.tickers[name]; exists {
			return // Ticker already exists
		}

		ticker := time.NewTicker(interval)
		tp.tickers[name] = ticker

		go func() {
			for range ticker.C {
				fn()
			}
		}()

		fmt.Printf("➕ Added ticker '%s' with interval %v\n", name, interval)
	}

	RemoveTicker := func(tp *TickerPool, name string) {
		tp.mu.Lock()
		defer tp.mu.Unlock()

		if ticker, exists := tp.tickers[name]; exists {
			ticker.Stop()
			delete(tp.tickers, name)
			fmt.Printf("➖ Removed ticker '%s'\n", name)
		}
	}

	StopAll := func(tp *TickerPool) {
		tp.mu.Lock()
		defer tp.mu.Unlock()

		for name, ticker := range tp.tickers {
			ticker.Stop()
			fmt.Printf("🛑 Stopped ticker '%s'\n", name)
		}
		tp.tickers = make(map[string]*time.Ticker)
	}

	// Use ticker pool
	pool := NewTickerPool()

	AddTicker(pool, "logger", 1*time.Second, func() {
		fmt.Printf("📝 Log entry at %s\n", time.Now().Format("15:04:05"))
	})

	AddTicker(pool, "monitor", 2*time.Second, func() {
		fmt.Printf("📊 System check at %s\n", time.Now().Format("15:04:05"))
	})

	time.Sleep(5 * time.Second)

	RemoveTicker(pool, "logger")
	time.Sleep(3 * time.Second)

	StopAll(pool)
}

func demoAdaptiveTicker() {
	fmt.Println("\n🎯 DEMO 9: Adaptive Ticker")
	fmt.Println("==========================")

	type AdaptiveTicker struct {
		ticker   *time.Ticker
		interval time.Duration
		mu       sync.Mutex
	}

	NewAdaptiveTicker := func(initialInterval time.Duration) *AdaptiveTicker {
		return &AdaptiveTicker{
			ticker:   time.NewTicker(initialInterval),
			interval: initialInterval,
		}
	}

	AdjustInterval := func(at *AdaptiveTicker, newInterval time.Duration) {
		at.mu.Lock()
		defer at.mu.Unlock()

		if newInterval != at.interval {
			at.ticker.Stop()
			at.ticker = time.NewTicker(newInterval)
			at.interval = newInterval
			fmt.Printf("🔄 Adjusted ticker interval to %v\n", newInterval)
		}
	}

	Stop := func(at *AdaptiveTicker) {
		at.mu.Lock()
		defer at.mu.Unlock()
		at.ticker.Stop()
	}

	C := func(at *AdaptiveTicker) <-chan time.Time {
		at.mu.Lock()
		defer at.mu.Unlock()
		return at.ticker.C
	}

	// Use adaptive ticker
	adaptiveTicker := NewAdaptiveTicker(1 * time.Second)
	defer Stop(adaptiveTicker)

	load := 0
	fmt.Println("Adaptive ticker adjusting based on system load:")

	timeout := time.After(8 * time.Second)

	for {
		select {
		case <-C(adaptiveTicker):
			load = (load + 1) % 10 // Simulate varying load
			fmt.Printf("⚡ Tick (load: %d) at %s\n", load, time.Now().Format("15:04:05"))

			// Adjust interval based on load
			if load > 7 {
				AdjustInterval(adaptiveTicker, 2*time.Second) // Slow down under high load
			} else if load < 3 {
				AdjustInterval(adaptiveTicker, 500*time.Millisecond) // Speed up under low load
			} else {
				AdjustInterval(adaptiveTicker, 1*time.Second) // Normal interval
			}

		case <-timeout:
			fmt.Println("⏰ Adaptive ticker demo completed")
			return
		}
	}
}

func demoTickerCleanup() {
	fmt.Println("\n🎯 DEMO 10: Proper Ticker Cleanup")
	fmt.Println("=================================")

	// Demonstrate proper cleanup patterns
	fmt.Println("Testing ticker cleanup patterns:")

	// Pattern 1: defer cleanup
	func() {
		fmt.Println("1. Using defer for cleanup:")
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()

		count := 0
		for range ticker.C {
			count++
			fmt.Printf("   Tick %d\n", count)
			if count >= 3 {
				break
			}
		}
		fmt.Println("   ✅ Ticker stopped via defer")
	}()

	// Pattern 2: explicit cleanup with done channel
	func() {
		fmt.Println("2. Using done channel for cleanup:")
		ticker := time.NewTicker(100 * time.Millisecond)
		done := make(chan bool)

		go func() {
			count := 0
			for {
				select {
				case <-ticker.C:
					count++
					fmt.Printf("   Background tick %d\n", count)
				case <-done:
					ticker.Stop()
					fmt.Println("   ✅ Background ticker stopped")
					return
				}
			}
		}()

		time.Sleep(350 * time.Millisecond)
		close(done)
		time.Sleep(100 * time.Millisecond) // Wait for cleanup
	}()

	// Pattern 3: context-based cleanup
	func() {
		fmt.Println("3. Using context for cleanup:")
		ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
		defer cancel()

		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()

		count := 0
		for {
			select {
			case <-ticker.C:
				count++
				fmt.Printf("   Context tick %d\n", count)
			case <-ctx.Done():
				fmt.Println("   ✅ Ticker stopped via context")
				return
			}
		}
	}()
}

func init() {
	// Run all demos
	go func() {
		time.Sleep(100 * time.Millisecond)
		demoPeriodicTasks()
		demoMultipleTickers()
		demoTickerWithContext()
		demoHeartbeat()
		demoDataCollection()
		demoRateLimiting()
		demoTickerPool()
		demoAdaptiveTicker()
		demoTickerCleanup()
		fmt.Println("\n✨ All ticker demos completed!")
	}()
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

🎵 TICKER CREATION:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Create ticker                                                        │
│ ticker := time.NewTicker(interval)                                      │
│                                                                         │
│ // Always stop ticker to prevent resource leaks                         │
│ defer ticker.Stop()                                                     │
│                                                                         │
│ // Receive ticks                                                        │
│ for range ticker.C {                                                    │
│     // Process tick                                                     │
│ }                                                                       │
│                                                                         │
│ // Or with select                                                       │
│ select {                                                                │
│ case <-ticker.C:                                                        │
│     // Handle tick                                                      │
│ case <-done:                                                            │
│     return                                                              │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

⏰ TICKER OPERATIONS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Stop ticker                                                          │
│ ticker.Stop()  // Stops ticker and frees resources                      │
│                                                                         │
│ // Reset ticker (not available - must stop and create new)              │
│ ticker.Stop()                                                           │
│ ticker = time.NewTicker(newInterval)                                    │
│                                                                         │
│ // Check if ticker channel has pending ticks                            │
│ select {                                                                │
│ case <-ticker.C:                                                        │
│     // Process tick                                                     │
│ default:                                                                │
│     // No pending tick                                                  │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

🎯 COMMON PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Periodic task execution                                              │
│ ticker := time.NewTicker(1 * time.Second)                               │
│ defer ticker.Stop()                                                     │
│                                                                         │
│ for {                                                                   │
│     select {                                                            │
│     case <-ticker.C:                                                    │
│         performPeriodicTask()                                           │
│     case <-ctx.Done():                                                  │
│         return                                                          │
│     }                                                                   │
│ }                                                                       │
│                                                                         │
│ // Heartbeat pattern                                                    │
│ heartbeat := time.NewTicker(30 * time.Second)                           │
│ defer heartbeat.Stop()                                                  │
│                                                                         │
│ for {                                                                   │
│     select {                                                            │
│     case <-heartbeat.C:                                                 │
│         sendHeartbeat()                                                 │
│     case <-shutdown:                                                    │
│         return                                                          │
│     }                                                                   │
│ }                                                                       │
│                                                                         │
│ // Rate limiting                                                        │
│ rateLimiter := time.NewTicker(time.Second / requestsPerSecond)          │
│ defer rateLimiter.Stop()                                                │
│                                                                         │
│ for _, request := range requests {                                      │
│     <-rateLimiter.C  // Wait for rate limit                             │
│     processRequest(request)                                             │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

🔄 TICKER MANAGEMENT:
┌─────────────────────────────────────────────────────────────────────────┐
│ type TickerManager struct {                                             │
│     tickers map[string]*time.Ticker                                     │
│     mu      sync.RWMutex                                                │
│ }                                                                       │
│                                                                         │
│ func (tm *TickerManager) Add(name string, interval time.Duration,       │
│                              fn func()) {                               │
│     tm.mu.Lock()                                                        │
│     defer tm.mu.Unlock()                                                │
│                                                                         │
│     if _, exists := tm.tickers[name]; exists {                          │
│         return  // Already exists                                       │
│     }                                                                   │
│                                                                         │
│     ticker := time.NewTicker(interval)                                  │
│     tm.tickers[name] = ticker                                           │
│                                                                         │
│     go func() {                                                         │
│         for range ticker.C {                                            │
│             fn()                                                        │
│         }                                                               │
│     }()                                                                 │
│ }                                                                       │
│                                                                         │
│ func (tm *TickerManager) Remove(name string) {                          │
│     tm.mu.Lock()                                                        │
│     defer tm.mu.Unlock()                                                │
│                                                                         │
│     if ticker, exists := tm.tickers[name]; exists {                     │
│         ticker.Stop()                                                   │
│         delete(tm.tickers, name)                                        │
│     }                                                                   │
│ }                                                                       │
│                                                                         │
│ func (tm *TickerManager) StopAll() {                                    │
│     tm.mu.Lock()                                                        │
│     defer tm.mu.Unlock()                                                │
│                                                                         │
│     for _, ticker := range tm.tickers {                                 │
│         ticker.Stop()                                                   │
│     }                                                                   │
│     tm.tickers = make(map[string]*time.Ticker)                          │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

⚡ PERFORMANCE CONSIDERATIONS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Ticker overhead                                                      │
│ • Each ticker uses ~200 bytes of memory                                 │
│ • Tickers are managed efficiently by Go runtime                         │
│ • Stopping tickers immediately frees resources                          │
│                                                                         │
│ // High-frequency tickers                                               │
│ • Be careful with very short intervals (< 1ms)                          │
│ • Consider batch processing for high-frequency events                   │
│ • Monitor CPU usage with high-frequency tickers                         │
│                                                                         │
│ // Memory considerations                                                 │
│ • Ticker channels are unbuffered                                        │
│ • Slow consumers can cause ticks to be dropped                          │
│ • Consider buffered processing for bursty workloads                     │
└─────────────────────────────────────────────────────────────────────────┘

🚨 COMMON MISTAKES:
┌─────────────────────────────────────────────────────────────────────────┐
│ // ❌ Not stopping tickers                                               │
│ ticker := time.NewTicker(1 * time.Second)                               │
│ // Ticker keeps running and consuming resources                          │
│                                                                         │
│ // ✅ Always stop tickers                                                │
│ ticker := time.NewTicker(1 * time.Second)                               │
│ defer ticker.Stop()                                                     │
│                                                                         │
│ // ❌ Blocking ticker processing                                         │
│ for range ticker.C {                                                    │
│     time.Sleep(2 * time.Second)  // Blocks ticker                       │
│ }                                                                       │
│                                                                         │
│ // ✅ Non-blocking ticker processing                                     │
│ for range ticker.C {                                                    │
│     go func() {                                                         │
│         time.Sleep(2 * time.Second)  // Process in goroutine            │
│     }()                                                                 │
│ }                                                                       │
│                                                                         │
│ // ❌ Creating tickers in loops                                          │
│ for {                                                                   │
│     ticker := time.NewTicker(1 * time.Second)  // Memory leak           │
│     <-ticker.C                                                          │
│ }                                                                       │
│                                                                         │
│ // ✅ Reuse ticker outside loop                                          │
│ ticker := time.NewTicker(1 * time.Second)                               │
│ defer ticker.Stop()                                                     │
│ for range ticker.C {                                                    │
│     // Process tick                                                     │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

💡 BEST PRACTICES:
• Always stop tickers when done to prevent resource leaks
• Use defer ticker.Stop() for automatic cleanup
• Process ticks in separate goroutines for long-running tasks
• Use select with context for cancellable ticker operations
• Consider ticker pools for managing multiple periodic tasks
• Monitor ticker performance in production
• Use appropriate intervals based on system capabilities

🎯 REAL-WORLD USE CASES:
• System monitoring and health checks
• Periodic data synchronization
• Cache cleanup and maintenance
• Heartbeat and keep-alive mechanisms
• Metrics collection and reporting
• Background job scheduling
• Rate limiting and throttling
• Animation and UI updates
• Log rotation and archival
• Database connection pool maintenance

🔄 TICKER VS TIMER:
┌─────────────────┬─────────────────┬─────────────────────────────────────┐
│   Aspect        │     Ticker      │            Timer                    │
├─────────────────┼─────────────────┼─────────────────────────────────────┤
│ Execution       │ Repeatedly      │ Once                                │
│ Use case        │ Periodic tasks  │ Timeouts, delays                    │
│ Reset           │ No (recreate)   │ Yes                                 │
│ Resource usage  │ Higher          │ Lower                               │
│ Precision       │ May drift       │ High precision                      │
│ Channel buffer  │ Unbuffered      │ Buffered (size 1)                   │
└─────────────────┴─────────────────┴─────────────────────────────────────┘

🔧 ADVANCED PATTERNS:
• Adaptive tickers that adjust interval based on load
• Hierarchical tickers for different priority levels
• Ticker synchronization across multiple services
• Jittered tickers to avoid thundering herd problems
• Ticker-based state machines
• Coordinated ticker shutdown in distributed systems

=============================================================================
*/