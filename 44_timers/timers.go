/*
=============================================================================
                           ⏲️ GO TIMERS TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Timers in Go provide a way to execute code after a specified duration.
Unlike tickers which repeat, timers fire once and can be stopped or reset.

🔑 KEY FEATURES:
• One-time delayed execution
• Cancellable and resettable
• Channel-based notification
• Efficient resource management
• Integration with select statements

💡 REAL-WORLD ANALOGY:
Timer = Kitchen Timer
- Set duration = How long to cook
- Fire once = Timer rings when done
- Reset = Set new cooking time
- Stop = Turn off timer early
- Channel = Timer bell notification

🎯 WHY USE TIMERS?
• Implement timeouts
• Delayed task execution
• Rate limiting and throttling
• Cleanup and maintenance tasks

=============================================================================
*/

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("⏲️ TIMERS TUTORIAL")
	fmt.Println("==================")

	// 🎯 DEMO 1: Basic Timer Usage
	fmt.Println("\n🎯 DEMO 1: Basic Timer")
	fmt.Println("======================")

	fmt.Println("Creating a 2-second timer...")
	timer := time.NewTimer(2 * time.Second)

	fmt.Printf("Waiting for timer at %s\n", time.Now().Format("15:04:05"))
	<-timer.C // Block until timer fires
	fmt.Printf("Timer fired at %s\n", time.Now().Format("15:04:05"))

	// 🎯 DEMO 2: Timer with Select
	fmt.Println("\n🎯 DEMO 2: Timer with Select")
	fmt.Println("============================")

	timer2 := time.NewTimer(1 * time.Second)
	defer timer2.Stop() // Always stop timers to free resources

	select {
	case <-timer2.C:
		fmt.Println("✅ Timer completed normally")
	case <-time.After(2 * time.Second):
		fmt.Println("❌ Timeout waiting for timer")
	}

	// 🎯 DEMO 3: Stopping a Timer
	fmt.Println("\n🎯 DEMO 3: Stopping a Timer")
	fmt.Println("============================")

	timer3 := time.NewTimer(3 * time.Second)
	
	go func() {
		time.Sleep(1 * time.Second)
		if timer3.Stop() {
			fmt.Println("🛑 Timer stopped before firing")
		} else {
			fmt.Println("⚠️ Timer had already fired")
		}
	}()

	select {
	case <-timer3.C:
		fmt.Println("Timer fired (this shouldn't happen)")
	case <-time.After(2 * time.Second):
		fmt.Println("✅ Confirmed timer was stopped")
	}

	// 🎯 DEMO 4: Resetting a Timer
	fmt.Println("\n🎯 DEMO 4: Resetting a Timer")
	fmt.Println("=============================")

	timer4 := time.NewTimer(3 * time.Second)
	defer timer4.Stop()

	fmt.Printf("Timer set for 3 seconds at %s\n", time.Now().Format("15:04:05"))

	// Reset timer to 1 second after 500ms
	go func() {
		time.Sleep(500 * time.Millisecond)
		if timer4.Reset(1 * time.Second) {
			fmt.Printf("🔄 Timer reset to 1 second at %s\n", time.Now().Format("15:04:05"))
		} else {
			fmt.Println("⚠️ Timer had already fired, reset anyway")
		}
	}()

	<-timer4.C
	fmt.Printf("✅ Timer fired at %s\n", time.Now().Format("15:04:05"))

	// 🎯 DEMO 5: Multiple Timers
	fmt.Println("\n🎯 DEMO 5: Multiple Timers")
	fmt.Println("==========================")

	timer5a := time.NewTimer(1 * time.Second)
	timer5b := time.NewTimer(2 * time.Second)
	timer5c := time.NewTimer(3 * time.Second)

	defer timer5a.Stop()
	defer timer5b.Stop()
	defer timer5c.Stop()

	fmt.Println("Starting three timers (1s, 2s, 3s)...")

	for i := 0; i < 3; i++ {
		select {
		case <-timer5a.C:
			fmt.Printf("⏰ Timer A (1s) fired at %s\n", time.Now().Format("15:04:05"))
		case <-timer5b.C:
			fmt.Printf("⏰ Timer B (2s) fired at %s\n", time.Now().Format("15:04:05"))
		case <-timer5c.C:
			fmt.Printf("⏰ Timer C (3s) fired at %s\n", time.Now().Format("15:04:05"))
		}
	}

	// 🎯 DEMO 6: Timer-based Timeout Pattern
	fmt.Println("\n🎯 DEMO 6: Timeout Pattern")
	fmt.Println("==========================")

	slowOperation := func() <-chan string {
		result := make(chan string, 1)
		go func() {
			time.Sleep(2 * time.Second) // Simulate slow work
			result <- "Operation completed"
		}()
		return result
	}

	timeout := time.NewTimer(1 * time.Second)
	defer timeout.Stop()

	fmt.Println("Starting slow operation with 1-second timeout...")

	select {
	case result := <-slowOperation():
		fmt.Printf("✅ %s\n", result)
	case <-timeout.C:
		fmt.Println("⏰ Operation timed out")
	}

	// 🎯 DEMO 7: Retry with Exponential Backoff
	fmt.Println("\n🎯 DEMO 7: Retry with Backoff")
	fmt.Println("=============================")

	retryOperation := func() error {
		// Simulate operation that fails first few times
		if time.Now().UnixNano()%3 != 0 {
			return fmt.Errorf("operation failed")
		}
		return nil
	}

	maxRetries := 5
	baseDelay := 100 * time.Millisecond

	for attempt := 1; attempt <= maxRetries; attempt++ {
		fmt.Printf("Attempt %d...\n", attempt)
		
		if err := retryOperation(); err == nil {
			fmt.Println("✅ Operation succeeded!")
			break
		}

		if attempt == maxRetries {
			fmt.Println("❌ All retries exhausted")
			break
		}

		// Exponential backoff
		delay := time.Duration(attempt) * baseDelay
		fmt.Printf("⏳ Retrying in %v\n", delay)
		
		timer := time.NewTimer(delay)
		<-timer.C
		timer.Stop()
	}

	// 🎯 DEMO 8: Debouncing with Timer
	fmt.Println("\n🎯 DEMO 8: Debouncing")
	fmt.Println("=====================")

	type Debouncer struct {
		timer *time.Timer
		delay time.Duration
	}

	NewDebouncer := func(delay time.Duration) *Debouncer {
		return &Debouncer{delay: delay}
	}

	debounce := func(d *Debouncer, fn func()) {
		if d.timer != nil {
			d.timer.Stop()
		}
		d.timer = time.NewTimer(d.delay)
		go func() {
			<-d.timer.C
			fn()
		}()
	}

	debouncer := NewDebouncer(500 * time.Millisecond)
	
	// Simulate rapid events
	events := []string{"event1", "event2", "event3", "event4", "event5"}
	
	fmt.Println("Sending rapid events (debounced to 500ms):")
	for i, event := range events {
		fmt.Printf("📨 Sending %s\n", event)
		
		debounce(debouncer, func() {
			fmt.Printf("🎯 Debounced action executed for final event\n")
		})
		
		if i < len(events)-1 {
			time.Sleep(100 * time.Millisecond) // Events come faster than debounce delay
		}
	}

	time.Sleep(1 * time.Second) // Wait for debounced action

	// 🎯 DEMO 9: Timer with Context
	fmt.Println("\n🎯 DEMO 9: Timer with Context")
	fmt.Println("=============================")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	timerWithContext := func(ctx context.Context, duration time.Duration) {
		timer := time.NewTimer(duration)
		defer timer.Stop()

		select {
		case <-timer.C:
			fmt.Println("✅ Timer completed")
		case <-ctx.Done():
			fmt.Printf("❌ Context cancelled: %v\n", ctx.Err())
		}
	}

	fmt.Println("Starting 3-second timer with 2-second context timeout...")
	timerWithContext(ctx, 3*time.Second)

	// 🎯 DEMO 10: Timer Pool for Efficiency
	fmt.Println("\n🎯 DEMO 10: Timer Pool")
	fmt.Println("======================")

	type TimerPool struct {
		pool chan *time.Timer
	}

	NewTimerPool := func(size int) *TimerPool {
		return &TimerPool{
			pool: make(chan *time.Timer, size),
		}
	}

	getTimer := func(tp *TimerPool, duration time.Duration) *time.Timer {
		select {
		case timer := <-tp.pool:
			timer.Reset(duration)
			return timer
		default:
			return time.NewTimer(duration)
		}
	}

	putTimer := func(tp *TimerPool, timer *time.Timer) {
		if !timer.Stop() {
			<-timer.C // Drain the channel if timer had fired
		}
		select {
		case tp.pool <- timer:
		default:
			// Pool is full, let timer be garbage collected
		}
	}

	timerPool := NewTimerPool(3)

	fmt.Println("Using timer pool for efficient timer reuse:")
	for i := 1; i <= 5; i++ {
		timer := getTimer(timerPool, 200*time.Millisecond)
		
		go func(id int) {
			<-timer.C
			fmt.Printf("⏰ Pooled timer %d fired\n", id)
			putTimer(timerPool, timer)
		}(i)
		
		time.Sleep(50 * time.Millisecond)
	}

	time.Sleep(1 * time.Second) // Wait for all timers

	// 🎯 DEMO 11: Cleanup Timer Pattern
	fmt.Println("\n🎯 DEMO 11: Cleanup Timer")
	fmt.Println("=========================")

	type Resource struct {
		id      int
		created time.Time
		timer   *time.Timer
	}

	resources := make(map[int]*Resource)
	
	createResource := func(id int) {
		resource := &Resource{
			id:      id,
			created: time.Now(),
		}
		
		// Set cleanup timer for 2 seconds
		resource.timer = time.NewTimer(2 * time.Second)
		go func() {
			<-resource.timer.C
			delete(resources, id)
			fmt.Printf("🗑️ Resource %d cleaned up after timeout\n", id)
		}()
		
		resources[id] = resource
		fmt.Printf("📦 Resource %d created\n", id)
	}

	accessResource := func(id int) {
		if resource, exists := resources[id]; exists {
			// Reset cleanup timer on access
			resource.timer.Reset(2 * time.Second)
			fmt.Printf("🔄 Resource %d accessed, cleanup timer reset\n", id)
		}
	}

	// Create resources
	createResource(1)
	createResource(2)
	
	time.Sleep(1 * time.Second)
	accessResource(1) // Reset timer for resource 1
	
	time.Sleep(1500 * time.Millisecond)
	fmt.Printf("📊 Resources remaining: %d\n", len(resources))
	
	time.Sleep(1 * time.Second)
	fmt.Printf("📊 Resources remaining: %d\n", len(resources))

	fmt.Println("\n✨ All timer demos completed!")
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

⏲️ TIMER CREATION:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Create timer                                                         │
│ timer := time.NewTimer(duration)                                        │
│                                                                         │
│ // One-shot timer (convenience function)                                │
│ <-time.After(duration)  // Equivalent to NewTimer(duration).C          │
│                                                                         │
│ // Timer with channel                                                   │
│ timer := time.NewTimer(5 * time.Second)                                 │
│ <-timer.C  // Block until timer fires                                   │
└─────────────────────────────────────────────────────────────────────────┘

🔧 TIMER OPERATIONS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Stop timer                                                           │
│ stopped := timer.Stop()  // Returns true if timer was stopped           │
│                                                                         │
│ // Reset timer                                                          │
│ reset := timer.Reset(newDuration)  // Returns true if timer was active  │
│                                                                         │
│ // Always stop timers to free resources                                 │
│ defer timer.Stop()                                                      │
│                                                                         │
│ // Drain channel after stop (if needed)                                 │
│ if !timer.Stop() {                                                      │
│     <-timer.C  // Drain channel if timer had fired                      │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

🎯 TIMER PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Timeout pattern                                                      │
│ timer := time.NewTimer(timeout)                                         │
│ defer timer.Stop()                                                      │
│                                                                         │
│ select {                                                                │
│ case result := <-operation():                                           │
│     // Operation completed                                              │
│ case <-timer.C:                                                         │
│     // Timeout occurred                                                 │
│ }                                                                       │
│                                                                         │
│ // Retry with backoff                                                   │
│ for attempt := 1; attempt <= maxRetries; attempt++ {                    │
│     if err := operation(); err == nil {                                 │
│         break                                                           │
│     }                                                                   │
│     if attempt < maxRetries {                                           │
│         delay := time.Duration(attempt) * baseDelay                     │
│         timer := time.NewTimer(delay)                                   │
│         <-timer.C                                                       │
│         timer.Stop()                                                    │
│     }                                                                   │
│ }                                                                       │
│                                                                         │
│ // Debouncing                                                           │
│ var debounceTimer *time.Timer                                           │
│ func debounce(fn func(), delay time.Duration) {                         │
│     if debounceTimer != nil {                                           │
│         debounceTimer.Stop()                                            │
│     }                                                                   │
│     debounceTimer = time.NewTimer(delay)                                │
│     go func() {                                                         │
│         <-debounceTimer.C                                               │
│         fn()                                                            │
│     }()                                                                 │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

⚡ PERFORMANCE CONSIDERATIONS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Timer overhead                                                       │
│ • Each timer uses ~200 bytes of memory                                  │
│ • Timers are managed by Go runtime efficiently                          │
│ • Stopping timers frees resources immediately                           │
│                                                                         │
│ // Timer pool for high-frequency usage                                  │
│ var timerPool = sync.Pool{                                              │
│     New: func() interface{} {                                           │
│         return time.NewTimer(0)                                         │
│     },                                                                  │
│ }                                                                       │
│                                                                         │
│ func getPooledTimer(d time.Duration) *time.Timer {                      │
│     timer := timerPool.Get().(*time.Timer)                              │
│     timer.Reset(d)                                                      │
│     return timer                                                        │
│ }                                                                       │
│                                                                         │
│ func putPooledTimer(timer *time.Timer) {                                │
│     if !timer.Stop() {                                                  │
│         <-timer.C                                                       │
│     }                                                                   │
│     timerPool.Put(timer)                                                │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

🚨 COMMON MISTAKES:
┌─────────────────────────────────────────────────────────────────────────┐
│ // ❌ Not stopping timers                                                │
│ timer := time.NewTimer(5 * time.Second)                                 │
│ // Timer keeps running even if not needed                               │
│                                                                         │
│ // ✅ Always stop timers                                                 │
│ timer := time.NewTimer(5 * time.Second)                                 │
│ defer timer.Stop()                                                      │
│                                                                         │
│ // ❌ Race condition with Reset                                          │
│ if !timer.Stop() {                                                      │
│     // Timer might fire between Stop and Reset                          │
│ }                                                                       │
│ timer.Reset(newDuration)                                                │
│                                                                         │
│ // ✅ Proper Reset pattern                                               │
│ if !timer.Stop() {                                                      │
│     <-timer.C  // Drain channel                                         │
│ }                                                                       │
│ timer.Reset(newDuration)                                                │
│                                                                         │
│ // ❌ Using time.After in loops                                          │
│ for {                                                                   │
│     select {                                                            │
│     case <-time.After(1 * time.Second):  // Creates new timer each time │
│         // Process                                                      │
│     }                                                                   │
│ }                                                                       │
│                                                                         │
│ // ✅ Reuse timer in loops                                               │
│ timer := time.NewTimer(1 * time.Second)                                 │
│ defer timer.Stop()                                                      │
│ for {                                                                   │
│     timer.Reset(1 * time.Second)                                        │
│     select {                                                            │
│     case <-timer.C:                                                     │
│         // Process                                                      │
│     }                                                                   │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

💡 BEST PRACTICES:
• Always stop timers when done to free resources
• Use defer timer.Stop() for automatic cleanup
• Drain timer channel after stopping if timer might have fired
• Prefer time.NewTimer over time.After for reusable timers
• Use timer pools for high-frequency timer usage
• Consider context.WithTimeout for cancellable operations
• Be careful with timer Reset in concurrent scenarios

🎯 REAL-WORLD USE CASES:
• HTTP request timeouts
• Cache expiration
• Rate limiting and throttling
• Retry mechanisms with backoff
• Debouncing user input
• Session timeouts
• Cleanup and garbage collection
• Circuit breaker patterns
• Heartbeat mechanisms
• Scheduled maintenance tasks

🔄 TIMER VS TICKER:
┌─────────────────┬─────────────────┬─────────────────────────────────────┐
│   Aspect        │     Timer       │            Ticker                   │
├─────────────────┼─────────────────┼─────────────────────────────────────┤
│ Execution       │ Once            │ Repeatedly                          │
│ Use case        │ Timeouts        │ Periodic tasks                      │
│ Reset           │ Yes             │ No (stop and create new)            │
│ Resource usage  │ Lower           │ Higher (continuous)                 │
│ Precision       │ High            │ May drift over time                 │
└─────────────────┴─────────────────┴─────────────────────────────────────┘

🔧 ADVANCED PATTERNS:
• Hierarchical timers (timer of timers)
• Adaptive timeouts based on system load
• Timer wheels for efficient timer management
• Combining timers with context for cancellation
• Timer-based state machines
• Coordinated timer shutdown in applications

=============================================================================
*/