/*
=============================================================================
                           🎯 GO SELECT STATEMENT TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Select is like a switch statement for channels. It lets you wait on multiple
channel operations simultaneously and execute the first one that's ready.

🔑 KEY FEATURES:
• Non-blocking channel operations
• Multiple channel communication
• Default case for immediate execution
• Random selection when multiple cases are ready

💡 REAL-WORLD ANALOGY:
Select = Restaurant Waiter
- Multiple tables (channels) need attention
- Waiter serves whichever table is ready first
- If no table is ready, waiter can do other tasks (default case)

🎯 WHY USE SELECT?
• Handle multiple channels efficiently
• Implement timeouts and non-blocking operations
• Create responsive concurrent programs
• Coordinate between multiple goroutines

=============================================================================
*/

package main

import (
	"fmt"
	"time"
)

// 📨 WORKER FUNCTIONS: Simulate different tasks
func worker1(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "Worker 1 finished"
}

func worker2(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Worker 2 finished"
}

func worker3(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "Worker 3 finished"
}

func main() {
	fmt.Println("🎯 SELECT STATEMENT DEMO")
	fmt.Println("========================")

	// 🎯 DEMO 1: Basic Select with Multiple Channels
	fmt.Println("\n🎯 DEMO 1: Multiple Channel Select")
	fmt.Println("==================================")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go worker1(ch1)
	go worker2(ch2)

	// 🎯 SELECT: Wait for first available channel
	select {
	case msg1 := <-ch1:
		fmt.Println("📥 Received from ch1:", msg1)
	case msg2 := <-ch2:
		fmt.Println("📥 Received from ch2:", msg2)
	}
	// Worker 2 finishes first (1 second), so we'll get that message

	// 🎯 DEMO 2: Select with Default Case
	fmt.Println("\n🎯 DEMO 2: Non-blocking Select")
	fmt.Println("==============================")

	ch3 := make(chan string)
	
	select {
	case msg := <-ch3:
		fmt.Println("📥 Received:", msg)
	default:
		fmt.Println("🚫 No data available, doing other work...")
	}
	// Default case executes immediately since ch3 has no data

	// 🎯 DEMO 3: Select with Timeout
	fmt.Println("\n🎯 DEMO 3: Timeout Pattern")
	fmt.Println("==========================")

	ch4 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		ch4 <- "Slow operation completed"
	}()

	select {
	case msg := <-ch4:
		fmt.Println("📥 Received:", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("⏰ Timeout! Operation took too long")
	}
	// Timeout occurs after 1 second, before the 2-second operation completes

	// 🎯 DEMO 4: Select in Loop (Fan-in Pattern)
	fmt.Println("\n🎯 DEMO 4: Fan-in Pattern")
	fmt.Println("=========================")

	ch5 := make(chan string)
	ch6 := make(chan string)
	ch7 := make(chan string)

	go worker1(ch5)
	go worker2(ch6)
	go worker3(ch7)

	// 📥 COLLECT ALL RESULTS: Wait for all workers
	completed := 0
	for completed < 3 {
		select {
		case msg := <-ch5:
			fmt.Println("📥 Worker 1:", msg)
			completed++
		case msg := <-ch6:
			fmt.Println("📥 Worker 2:", msg)
			completed++
		case msg := <-ch7:
			fmt.Println("📥 Worker 3:", msg)
			completed++
		}
	}

	fmt.Println("\n✨ All select demos completed!")
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

🎯 SELECT STATEMENT SYNTAX:
┌─────────────────────────────────────────────────────────────────────────┐
│ select {                                                                │
│ case <-ch1:                    // Receive from ch1                      │
│     // Handle ch1 data                                                  │
│ case data := <-ch2:            // Receive from ch2 with assignment      │
│     // Handle ch2 data                                                  │
│ case ch3 <- value:             // Send to ch3                           │
│     // Handle successful send                                           │
│ default:                       // Optional: executes if no case ready   │
│     // Handle no channels ready                                         │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

⚡ SELECT BEHAVIOR:
• Blocks until one case can proceed
• If multiple cases are ready, chooses randomly
• Default case makes select non-blocking
• Empty select{} blocks forever

🎯 COMMON PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Timeout pattern                                                      │
│ select {                                                                │
│ case result := <-ch:                                                    │
│     // Handle result                                                    │
│ case <-time.After(5 * time.Second):                                     │
│     // Handle timeout                                                   │
│ }                                                                       │
│                                                                         │
│ // Non-blocking receive                                                 │
│ select {                                                                │
│ case data := <-ch:                                                      │
│     // Handle data                                                      │
│ default:                                                                │
│     // Channel empty, do something else                                 │
│ }                                                                       │
│                                                                         │
│ // Fan-in (multiple producers, one consumer)                            │
│ for {                                                                   │
│     select {                                                            │
│     case msg1 := <-ch1:                                                 │
│         // Handle from source 1                                        │
│     case msg2 := <-ch2:                                                 │
│         // Handle from source 2                                        │
│     }                                                                   │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

🚨 GOTCHAS:
❌ Select without default blocks until a case is ready
❌ Empty select{} blocks forever (infinite wait)
❌ Select chooses randomly among ready cases
❌ Nil channels in select are ignored

💡 BEST PRACTICES:
• Use timeouts to prevent indefinite blocking
• Use default case for non-blocking operations
• Combine with for loops for continuous monitoring
• Handle all possible channel states (open, closed, nil)

🎯 REAL-WORLD USES:
• HTTP request timeouts
• Worker pool coordination
• Event handling systems
• Graceful shutdown patterns
• Load balancing between services

=============================================================================
*/