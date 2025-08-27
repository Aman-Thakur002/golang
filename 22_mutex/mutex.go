/*
=============================================================================
                           🔒 GO MUTEX TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Mutex (Mutual Exclusion) prevents race conditions by ensuring only one
goroutine can access shared data at a time. It's like a lock for your data.

🔑 KEY FEATURES:
• Prevents race conditions
• Ensures thread-safe access to shared resources
• Two types: Mutex and RWMutex
• Lock/Unlock operations

💡 REAL-WORLD ANALOGY:
Mutex = Bathroom Lock
- Only one person can use the bathroom at a time
- Others must wait until it's unlocked
- Prevents conflicts and ensures privacy

🎯 WHY USE MUTEX?
• Protect shared variables from concurrent access
• Prevent data corruption in multi-goroutine programs
• Ensure atomic operations on critical sections

=============================================================================
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

// 📊 SHARED DATA: This will be accessed by multiple goroutines
var (
	counter int
	mutex   sync.Mutex    // 🔒 Mutex to protect counter
	rwMutex sync.RWMutex  // 📖 Read-Write mutex for advanced scenarios
)

// 🔢 UNSAFE COUNTER: Without mutex protection
func unsafeIncrement() {
	for i := 0; i < 1000; i++ {
		counter++ // ❌ RACE CONDITION: Multiple goroutines modifying same variable
	}
}

// 🔒 SAFE COUNTER: With mutex protection
func safeIncrement() {
	for i := 0; i < 1000; i++ {
		mutex.Lock()   // 🔒 LOCK: Only one goroutine can proceed
		counter++      // ✅ SAFE: Protected by mutex
		mutex.Unlock() // 🔓 UNLOCK: Allow other goroutines to proceed
	}
}

// 📖 READ-WRITE MUTEX EXAMPLE
type SafeMap struct {
	mu   sync.RWMutex
	data map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()         // 🔒 WRITE LOCK: Exclusive access
	defer sm.mu.Unlock() // 🔓 Ensure unlock even if panic occurs
	sm.data[key] = value
}

func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.RLock()         // 📖 READ LOCK: Multiple readers allowed
	defer sm.mu.RUnlock() // 🔓 Unlock read lock
	value, exists := sm.data[key]
	return value, exists
}

func main() {
	fmt.Println("🔒 MUTEX TUTORIAL")
	fmt.Println("=================")

	// 🎯 DEMO 1: Race Condition (Unsafe)
	fmt.Println("\n🎯 DEMO 1: Race Condition")
	fmt.Println("=========================")

	counter = 0 // Reset counter
	var wg sync.WaitGroup

	// Start 5 goroutines without mutex protection
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			unsafeIncrement()
			fmt.Printf("🔴 Goroutine %d finished\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Printf("🔴 Unsafe counter result: %d (should be 5000)\n", counter)
	// Result will likely be less than 5000 due to race conditions

	// 🎯 DEMO 2: Mutex Protection (Safe)
	fmt.Println("\n🎯 DEMO 2: Mutex Protection")
	fmt.Println("===========================")

	counter = 0 // Reset counter

	// Start 5 goroutines with mutex protection
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			safeIncrement()
			fmt.Printf("✅ Goroutine %d finished\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Printf("✅ Safe counter result: %d (exactly 5000)\n", counter)

	// 🎯 DEMO 3: Read-Write Mutex
	fmt.Println("\n🎯 DEMO 3: Read-Write Mutex")
	fmt.Println("===========================")

	safeMap := NewSafeMap()

	// 📝 WRITERS: Set values concurrently
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", id)
			safeMap.Set(key, id*10)
			fmt.Printf("📝 Writer %d: Set %s = %d\n", id, key, id*10)
		}(i)
	}

	// Give writers time to start
	time.Sleep(100 * time.Millisecond)

	// 📖 READERS: Read values concurrently
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", id%3) // Read existing keys
			if value, exists := safeMap.Get(key); exists {
				fmt.Printf("📖 Reader %d: Got %s = %d\n", id, key, value)
			} else {
				fmt.Printf("📖 Reader %d: Key %s not found\n", id, key)
			}
		}(i)
	}

	wg.Wait()

	// 🎯 DEMO 4: Defer with Mutex (Best Practice)
	fmt.Println("\n🎯 DEMO 4: Defer Pattern")
	fmt.Println("========================")

	var criticalSection = func(id int) {
		mutex.Lock()
		defer mutex.Unlock() // 🔓 DEFER: Ensures unlock even if panic occurs

		fmt.Printf("🔒 Goroutine %d: Entered critical section\n", id)
		time.Sleep(100 * time.Millisecond) // Simulate work
		fmt.Printf("🔓 Goroutine %d: Leaving critical section\n", id)
	}

	// Start multiple goroutines
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			criticalSection(id)
		}(i)
	}

	wg.Wait()

	fmt.Println("\n✨ All mutex demos completed!")
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

🔒 MUTEX TYPES:
┌─────────────────┬─────────────────┬─────────────────────────────────────┐
│   Mutex Type    │   Operations    │           Use Case                  │
├─────────────────┼─────────────────┼─────────────────────────────────────┤
│ sync.Mutex      │ Lock/Unlock     │ Exclusive access to shared data     │
│ sync.RWMutex    │ RLock/RUnlock   │ Multiple readers, exclusive writers │
│                 │ Lock/Unlock     │                                     │
└─────────────────┴─────────────────┴─────────────────────────────────────┘

🔒 MUTEX OPERATIONS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Basic Mutex                                                          │
│ var mu sync.Mutex                                                       │
│ mu.Lock()     // Acquire exclusive lock                                 │
│ // critical section                                                     │
│ mu.Unlock()   // Release lock                                           │
│                                                                         │
│ // Read-Write Mutex                                                     │
│ var rwMu sync.RWMutex                                                   │
│ rwMu.RLock()    // Acquire read lock (multiple allowed)                 │
│ // read operation                                                       │
│ rwMu.RUnlock()  // Release read lock                                    │
│                                                                         │
│ rwMu.Lock()     // Acquire write lock (exclusive)                       │
│ // write operation                                                      │
│ rwMu.Unlock()   // Release write lock                                   │
└─────────────────────────────────────────────────────────────────────────┘

⚡ RACE CONDITIONS:
• Occur when multiple goroutines access shared data simultaneously
• Can cause data corruption, inconsistent state, crashes
• Detected with: go run -race program.go

🔧 BEST PRACTICES:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Always use defer for unlock                                          │
│ mu.Lock()                                                               │
│ defer mu.Unlock()                                                       │
│                                                                         │
│ // Keep critical sections small                                         │
│ mu.Lock()                                                               │
│ // minimal code here                                                    │
│ mu.Unlock()                                                             │
│                                                                         │
│ // Use RWMutex for read-heavy workloads                                 │
│ rwMu.RLock()                                                            │
│ value := sharedData                                                     │
│ rwMu.RUnlock()                                                          │
└─────────────────────────────────────────────────────────────────────────┘

🚨 COMMON MISTAKES:
❌ Forgetting to unlock (causes deadlock)
❌ Unlocking without locking
❌ Copying mutex (use pointers)
❌ Nested locking same mutex (deadlock)
❌ Long-running operations in critical section

💡 ALTERNATIVES TO MUTEX:
• Channels: "Don't communicate by sharing memory; share memory by communicating"
• sync.Once: For one-time initialization
• sync/atomic: For simple atomic operations
• Context: For cancellation and timeouts

🎯 WHEN TO USE EACH:
• Mutex: Simple exclusive access to shared data
• RWMutex: Many readers, few writers
• Channels: Communication between goroutines
• Atomic: Simple counters, flags

⚡ PERFORMANCE TIPS:
• RWMutex is slower than Mutex for write-heavy workloads
• Keep critical sections as short as possible
• Consider lock-free alternatives for high-performance scenarios
• Use sync.Pool for object reuse to reduce lock contention

=============================================================================
*/