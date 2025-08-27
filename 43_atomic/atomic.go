/*
=============================================================================
                        ⚛️ GO ATOMIC OPERATIONS TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Atomic operations provide lock-free synchronization for simple operations
on shared variables. They're faster than mutexes for basic operations
and help build lock-free data structures.

🔑 KEY FEATURES:
• Lock-free synchronization
• Atomic read/write operations
• Compare-and-swap operations
• Memory ordering guarantees
• High performance for simple operations

💡 REAL-WORLD ANALOGY:
Atomic Operations = Bank ATM Transaction
- Atomic = Transaction either completes fully or not at all
- Compare-and-swap = Check balance, then withdraw if sufficient
- Memory barriers = Ensuring transaction order
- Lock-free = No need to lock entire bank

🎯 WHY USE ATOMIC OPERATIONS?
• High-performance counters and flags
• Lock-free data structures
• Reduce contention in hot paths
• Building blocks for complex synchronization

=============================================================================
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// 📊 ATOMIC COUNTER EXAMPLE
type AtomicCounter struct {
	value int64
}

func (c *AtomicCounter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

func (c *AtomicCounter) Decrement() {
	atomic.AddInt64(&c.value, -1)
}

func (c *AtomicCounter) Get() int64 {
	return atomic.LoadInt64(&c.value)
}

func (c *AtomicCounter) Set(value int64) {
	atomic.StoreInt64(&c.value, value)
}

func (c *AtomicCounter) CompareAndSwap(old, new int64) bool {
	return atomic.CompareAndSwapInt64(&c.value, old, new)
}

// 🔄 LOCK-FREE STACK EXAMPLE
type LockFreeStack struct {
	head unsafe.Pointer
}

type stackNode struct {
	value interface{}
	next  unsafe.Pointer
}

func NewLockFreeStack() *LockFreeStack {
	return &LockFreeStack{}
}

func (s *LockFreeStack) Push(value interface{}) {
	newNode := &stackNode{value: value}
	for {
		head := atomic.LoadPointer(&s.head)
		newNode.next = head
		if atomic.CompareAndSwapPointer(&s.head, head, unsafe.Pointer(newNode)) {
			break
		}
	}
}

func (s *LockFreeStack) Pop() interface{} {
	for {
		head := atomic.LoadPointer(&s.head)
		if head == nil {
			return nil
		}
		node := (*stackNode)(head)
		if atomic.CompareAndSwapPointer(&s.head, head, node.next) {
			return node.value
		}
	}
}

// 🚦 ATOMIC FLAG EXAMPLE
type AtomicFlag struct {
	flag int32
}

func (f *AtomicFlag) Set() {
	atomic.StoreInt32(&f.flag, 1)
}

func (f *AtomicFlag) Clear() {
	atomic.StoreInt32(&f.flag, 0)
}

func (f *AtomicFlag) IsSet() bool {
	return atomic.LoadInt32(&f.flag) != 0
}

func (f *AtomicFlag) TestAndSet() bool {
	return atomic.SwapInt32(&f.flag, 1) != 0
}

// 📈 PERFORMANCE COMPARISON
func compareCounterPerformance() {
	fmt.Println("📈 Performance Comparison")
	fmt.Println("========================")

	const numGoroutines = 10
	const numIncrements = 100000

	// Test atomic counter
	fmt.Println("Testing atomic counter...")
	atomicCounter := &AtomicCounter{}
	start := time.Now()

	var wg sync.WaitGroup
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < numIncrements; j++ {
				atomicCounter.Increment()
			}
		}()
	}
	wg.Wait()

	atomicTime := time.Since(start)
	fmt.Printf("Atomic counter: %v, Final value: %d\n", atomicTime, atomicCounter.Get())

	// Test mutex counter
	fmt.Println("Testing mutex counter...")
	var mutexCounter int64
	var mu sync.Mutex
	start = time.Now()

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < numIncrements; j++ {
				mu.Lock()
				mutexCounter++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()

	mutexTime := time.Since(start)
	fmt.Printf("Mutex counter: %v, Final value: %d\n", mutexTime, mutexCounter)
	fmt.Printf("Atomic is %.2fx faster\n", float64(mutexTime)/float64(atomicTime))
}

func main() {
	fmt.Println("⚛️ ATOMIC OPERATIONS TUTORIAL")
	fmt.Println("=============================")

	// 🎯 DEMO 1: Basic Atomic Operations
	fmt.Println("\n🎯 DEMO 1: Basic Atomic Operations")
	fmt.Println("==================================")

	var counter int64

	// Atomic increment
	atomic.AddInt64(&counter, 1)
	fmt.Printf("After increment: %d\n", atomic.LoadInt64(&counter))

	// Atomic add
	atomic.AddInt64(&counter, 5)
	fmt.Printf("After adding 5: %d\n", atomic.LoadInt64(&counter))

	// Atomic store
	atomic.StoreInt64(&counter, 100)
	fmt.Printf("After store 100: %d\n", atomic.LoadInt64(&counter))

	// Atomic swap
	old := atomic.SwapInt64(&counter, 200)
	fmt.Printf("Swapped %d with 200, new value: %d\n", old, atomic.LoadInt64(&counter))

	// Compare and swap
	swapped := atomic.CompareAndSwapInt64(&counter, 200, 300)
	fmt.Printf("CAS(200->300): %t, value: %d\n", swapped, atomic.LoadInt64(&counter))

	swapped = atomic.CompareAndSwapInt64(&counter, 200, 400)
	fmt.Printf("CAS(200->400): %t, value: %d\n", swapped, atomic.LoadInt64(&counter))

	// 🎯 DEMO 2: Atomic Counter
	fmt.Println("\n🎯 DEMO 2: Atomic Counter")
	fmt.Println("=========================")

	atomicCounter := &AtomicCounter{}
	
	// Concurrent increments
	var wg sync.WaitGroup
	numGoroutines := 5
	incrementsPerGoroutine := 1000

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < incrementsPerGoroutine; j++ {
				atomicCounter.Increment()
			}
			fmt.Printf("Goroutine %d finished, counter: %d\n", id, atomicCounter.Get())
		}(i)
	}

	wg.Wait()
	expectedValue := int64(numGoroutines * incrementsPerGoroutine)
	actualValue := atomicCounter.Get()
	fmt.Printf("Expected: %d, Actual: %d, Correct: %t\n", 
		expectedValue, actualValue, expectedValue == actualValue)

	// 🎯 DEMO 3: Atomic Flags
	fmt.Println("\n🎯 DEMO 3: Atomic Flags")
	fmt.Println("=======================")

	flag := &AtomicFlag{}
	
	fmt.Printf("Initial flag state: %t\n", flag.IsSet())
	
	flag.Set()
	fmt.Printf("After set: %t\n", flag.IsSet())
	
	wasSet := flag.TestAndSet()
	fmt.Printf("Test and set (was set: %t): %t\n", wasSet, flag.IsSet())
	
	flag.Clear()
	fmt.Printf("After clear: %t\n", flag.IsSet())
	
	wasSet = flag.TestAndSet()
	fmt.Printf("Test and set (was set: %t): %t\n", wasSet, flag.IsSet())

	// 🎯 DEMO 4: Different Atomic Types
	fmt.Println("\n🎯 DEMO 4: Different Atomic Types")
	fmt.Println("=================================")

	// int32
	var int32Val int32
	atomic.AddInt32(&int32Val, 42)
	fmt.Printf("int32 value: %d\n", atomic.LoadInt32(&int32Val))

	// uint64
	var uint64Val uint64
	atomic.AddUint64(&uint64Val, 123)
	fmt.Printf("uint64 value: %d\n", atomic.LoadUint64(&uint64Val))

	// uintptr (useful for pointers)
	var uintptrVal uintptr
	atomic.StoreUintptr(&uintptrVal, 0xDEADBEEF)
	fmt.Printf("uintptr value: 0x%X\n", atomic.LoadUintptr(&uintptrVal))

	// 🎯 DEMO 5: Compare-and-Swap Patterns
	fmt.Println("\n🎯 DEMO 5: Compare-and-Swap Patterns")
	fmt.Println("====================================")

	var value int64 = 10

	// Atomic increment using CAS (alternative to AddInt64)
	for {
		old := atomic.LoadInt64(&value)
		new := old + 1
		if atomic.CompareAndSwapInt64(&value, old, new) {
			fmt.Printf("CAS increment: %d -> %d\n", old, new)
			break
		}
		// If CAS failed, retry (another goroutine modified the value)
	}

	// Atomic maximum using CAS
	newMax := int64(15)
	for {
		old := atomic.LoadInt64(&value)
		if newMax <= old {
			fmt.Printf("Value %d is already >= %d\n", old, newMax)
			break
		}
		if atomic.CompareAndSwapInt64(&value, old, newMax) {
			fmt.Printf("Updated max: %d -> %d\n", old, newMax)
			break
		}
	}

	// 🎯 DEMO 6: Atomic Pointer Operations
	fmt.Println("\n🎯 DEMO 6: Atomic Pointer Operations")
	fmt.Println("====================================")

	type Data struct {
		Value int
		Name  string
	}

	var dataPtr unsafe.Pointer
	
	// Store pointer atomically
	data1 := &Data{Value: 42, Name: "first"}
	atomic.StorePointer(&dataPtr, unsafe.Pointer(data1))
	
	// Load pointer atomically
	loadedPtr := atomic.LoadPointer(&dataPtr)
	loadedData := (*Data)(loadedPtr)
	fmt.Printf("Loaded data: %+v\n", *loadedData)
	
	// Swap pointer atomically
	data2 := &Data{Value: 84, Name: "second"}
	oldPtr := atomic.SwapPointer(&dataPtr, unsafe.Pointer(data2))
	oldData := (*Data)(oldPtr)
	fmt.Printf("Swapped from: %+v\n", *oldData)
	fmt.Printf("Swapped to: %+v\n", *(*Data)(atomic.LoadPointer(&dataPtr)))

	// 🎯 DEMO 7: Atomic Value (type-safe)
	fmt.Println("\n🎯 DEMO 7: Atomic Value")
	fmt.Println("=======================")

	var atomicValue atomic.Value
	
	// Store and load different types (must be consistent)
	atomicValue.Store("Hello, Atomic!")
	loaded := atomicValue.Load()
	fmt.Printf("Loaded string: %s\n", loaded.(string))
	
	// Store struct
	config := struct {
		Host string
		Port int
	}{"localhost", 8080}
	
	atomicValue.Store(config)
	loadedConfig := atomicValue.Load().(struct {
		Host string
		Port int
	})
	fmt.Printf("Loaded config: %+v\n", loadedConfig)

	// 🎯 DEMO 8: Memory Ordering Example
	fmt.Println("\n🎯 DEMO 8: Memory Ordering")
	fmt.Println("==========================")

	var flag int32
	var data int32

	// Writer goroutine
	go func() {
		atomic.StoreInt32(&data, 42)        // Store data first
		atomic.StoreInt32(&flag, 1)         // Then set flag
	}()

	// Reader goroutine
	go func() {
		for atomic.LoadInt32(&flag) == 0 {  // Wait for flag
			runtime.Gosched()
		}
		value := atomic.LoadInt32(&data)    // Read data after flag is set
		fmt.Printf("Read data after flag: %d\n", value)
	}()

	time.Sleep(100 * time.Millisecond) // Give goroutines time to run

	// 🎯 DEMO 9: Performance Comparison
	compareCounterPerformance()

	// 🎯 DEMO 10: Practical Example - Connection Pool
	fmt.Println("\n🎯 DEMO 10: Connection Pool Example")
	fmt.Println("===================================")

	type ConnectionPool struct {
		activeConnections int64
		maxConnections    int64
	}

	pool := &ConnectionPool{maxConnections: 5}

	acquireConnection := func(pool *ConnectionPool) bool {
		for {
			current := atomic.LoadInt64(&pool.activeConnections)
			if current >= pool.maxConnections {
				return false // Pool exhausted
			}
			if atomic.CompareAndSwapInt64(&pool.activeConnections, current, current+1) {
				return true // Successfully acquired
			}
			// CAS failed, retry
		}
	}

	releaseConnection := func(pool *ConnectionPool) {
		atomic.AddInt64(&pool.activeConnections, -1)
	}

	// Simulate concurrent connection requests
	var wg sync.WaitGroup
	successCount := int64(0)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if acquireConnection(pool) {
				atomic.AddInt64(&successCount, 1)
				fmt.Printf("Connection %d acquired (active: %d)\n", 
					id, atomic.LoadInt64(&pool.activeConnections))
				time.Sleep(100 * time.Millisecond) // Simulate work
				releaseConnection(pool)
				fmt.Printf("Connection %d released (active: %d)\n", 
					id, atomic.LoadInt64(&pool.activeConnections))
			} else {
				fmt.Printf("Connection %d failed - pool exhausted\n", id)
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("Total successful connections: %d\n", atomic.LoadInt64(&successCount))

	fmt.Println("\n✨ All atomic operations demos completed!")
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

⚛️ ATOMIC OPERATION TYPES:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Load operations                                                      │
│ atomic.LoadInt32(&var)    // Atomically load int32                      │
│ atomic.LoadInt64(&var)    // Atomically load int64                      │
│ atomic.LoadUint32(&var)   // Atomically load uint32                     │
│ atomic.LoadUint64(&var)   // Atomically load uint64                     │
│ atomic.LoadUintptr(&var)  // Atomically load uintptr                    │
│ atomic.LoadPointer(&var)  // Atomically load unsafe.Pointer             │
│                                                                         │
│ // Store operations                                                     │
│ atomic.StoreInt32(&var, val)    // Atomically store int32               │
│ atomic.StoreInt64(&var, val)    // Atomically store int64               │
│ atomic.StoreUint32(&var, val)   // Atomically store uint32              │
│ atomic.StoreUint64(&var, val)   // Atomically store uint64              │
│ atomic.StoreUintptr(&var, val)  // Atomically store uintptr             │
│ atomic.StorePointer(&var, val)  // Atomically store unsafe.Pointer      │
└─────────────────────────────────────────────────────────────────────────┘

🔄 MODIFY OPERATIONS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Add operations                                                       │
│ atomic.AddInt32(&var, delta)    // Add delta to int32                   │
│ atomic.AddInt64(&var, delta)    // Add delta to int64                   │
│ atomic.AddUint32(&var, delta)   // Add delta to uint32                  │
│ atomic.AddUint64(&var, delta)   // Add delta to uint64                  │
│ atomic.AddUintptr(&var, delta)  // Add delta to uintptr                 │
│                                                                         │
│ // Swap operations                                                      │
│ old := atomic.SwapInt32(&var, new)    // Swap and return old value      │
│ old := atomic.SwapInt64(&var, new)    // Swap and return old value      │
│ old := atomic.SwapUint32(&var, new)   // Swap and return old value      │
│ old := atomic.SwapUint64(&var, new)   // Swap and return old value      │
│ old := atomic.SwapUintptr(&var, new)  // Swap and return old value      │
│ old := atomic.SwapPointer(&var, new)  // Swap and return old pointer    │
└─────────────────────────────────────────────────────────────────────────┘

🎯 COMPARE-AND-SWAP:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Compare-and-swap operations                                          │
│ swapped := atomic.CompareAndSwapInt32(&var, old, new)                   │
│ swapped := atomic.CompareAndSwapInt64(&var, old, new)                   │
│ swapped := atomic.CompareAndSwapUint32(&var, old, new)                  │
│ swapped := atomic.CompareAndSwapUint64(&var, old, new)                  │
│ swapped := atomic.CompareAndSwapUintptr(&var, old, new)                 │
│ swapped := atomic.CompareAndSwapPointer(&var, old, new)                 │
│                                                                         │
│ // CAS loop pattern                                                     │
│ for {                                                                   │
│     old := atomic.LoadInt64(&counter)                                   │
│     new := computeNewValue(old)                                         │
│     if atomic.CompareAndSwapInt64(&counter, old, new) {                 │
│         break  // Success                                               │
│     }                                                                   │
│     // Retry if another goroutine modified the value                    │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

📦 ATOMIC.VALUE:
┌─────────────────────────────────────────────────────────────────────────┐
│ var v atomic.Value                                                      │
│                                                                         │
│ // Store any type (must be consistent)                                  │
│ v.Store("hello")                                                        │
│ v.Store(42)        // Error: inconsistent type                          │
│                                                                         │
│ // Load with type assertion                                             │
│ val := v.Load().(string)                                                │
│                                                                         │
│ // Safe loading with type check                                         │
│ if val := v.Load(); val != nil {                                        │
│     if str, ok := val.(string); ok {                                    │
│         // Use str                                                      │
│     }                                                                   │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

🔒 LOCK-FREE PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Lock-free counter                                                    │
│ type Counter struct {                                                   │
│     value int64                                                         │
│ }                                                                       │
│                                                                         │
│ func (c *Counter) Inc() int64 {                                         │
│     return atomic.AddInt64(&c.value, 1)                                 │
│ }                                                                       │
│                                                                         │
│ func (c *Counter) Get() int64 {                                         │
│     return atomic.LoadInt64(&c.value)                                   │
│ }                                                                       │
│                                                                         │
│ // Lock-free flag                                                       │
│ type Flag struct {                                                      │
│     state int32                                                         │
│ }                                                                       │
│                                                                         │
│ func (f *Flag) Set() {                                                  │
│     atomic.StoreInt32(&f.state, 1)                                      │
│ }                                                                       │
│                                                                         │
│ func (f *Flag) IsSet() bool {                                           │
│     return atomic.LoadInt32(&f.state) != 0                              │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

⚡ PERFORMANCE CHARACTERISTICS:
┌─────────────────┬─────────────────┬─────────────────┬─────────────────┐
│   Operation     │   Atomic        │     Mutex       │     Channel     │
├─────────────────┼─────────────────┼─────────────────┼─────────────────┤
│ Simple counter  │ ~1-2 ns         │ ~20-30 ns       │ ~100+ ns        │
│ Memory usage    │ Variable size   │ 8+ bytes        │ 96+ bytes       │
│ Contention      │ Lock-free       │ Blocking        │ Blocking        │
│ Complexity      │ Simple          │ Medium          │ High            │
└─────────────────┴─────────────────┴─────────────────┴─────────────────┘

💡 BEST PRACTICES:
• Use atomic operations for simple shared state
• Prefer atomic operations over mutexes for counters and flags
• Use CAS loops for complex atomic updates
• Be careful with memory ordering in complex scenarios
• Profile to ensure atomic operations provide benefit
• Use atomic.Value for type-safe atomic references

🚨 COMMON MISTAKES:
❌ Mixing atomic and non-atomic access to same variable
❌ Using atomic operations on non-aligned memory
❌ Assuming atomic operations provide ordering for other variables
❌ Using atomic.Value with inconsistent types
❌ Not understanding ABA problem in CAS loops
❌ Overusing atomic operations where mutexes would be clearer

🎯 WHEN TO USE ATOMIC OPERATIONS:
✅ Simple counters and statistics
✅ Flags and boolean state
✅ Reference counting
✅ Lock-free data structures
✅ High-performance scenarios with low contention

❌ Complex shared state
❌ Multiple related variables that must be updated together
❌ When code clarity is more important than performance
❌ Scenarios with high contention (consider channels/mutexes)

🔧 MEMORY ORDERING:
• Go's atomic operations provide sequential consistency
• Atomic operations act as memory barriers
• Use atomic operations to establish happens-before relationships
• Be careful when mixing atomic and non-atomic operations

⚛️ ADVANCED PATTERNS:
• ABA problem mitigation with versioning
• Lock-free queues and stacks
• Reference counting for memory management
• Atomic snapshots of multiple variables
• Wait-free algorithms

=============================================================================
*/

import "unsafe" // Add this import at the top