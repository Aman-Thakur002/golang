/*
=============================================================================
                           ⏰ GO DEFER STATEMENT TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Defer schedules a function call to be executed when the surrounding function
returns. It's like setting a reminder to clean up before leaving.

🔑 KEY FEATURES:
• Executes in LIFO order (Last In, First Out)
• Runs even if function panics
• Arguments evaluated immediately, execution deferred
• Perfect for cleanup operations

💡 REAL-WORLD ANALOGY:
Defer = Leaving a Room Checklist
- Turn off lights (defer cleanup)
- Lock the door (defer security)
- Set alarm (defer monitoring)
- All happen automatically when you leave

🎯 WHY USE DEFER?
• Guaranteed cleanup (files, connections, locks)
• Keep cleanup code near setup code
• Exception-safe resource management
• Simplify error handling

=============================================================================
*/

package main

import (
	"fmt"
	"os"
	"sync"
)

// 📁 FILE OPERATIONS: Demonstrate defer with file handling
func readFileWithDefer(filename string) error {
	fmt.Printf("📁 Opening file: %s\n", filename)
	
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	
	// 🔒 DEFER: Ensure file is closed even if error occurs
	defer func() {
		fmt.Printf("🔒 Closing file: %s\n", filename)
		file.Close()
	}()
	
	// Simulate file operations
	fmt.Printf("📖 Reading from file: %s\n", filename)
	
	// File will be automatically closed when function returns
	return nil
}

// 🔢 COUNTER WITH MUTEX: Demonstrate defer with locks
func safeCounter() {
	var mu sync.Mutex
	var count int
	
	increment := func() {
		mu.Lock()
		defer mu.Unlock() // 🔓 DEFER: Ensure unlock even if panic
		
		count++
		fmt.Printf("🔢 Count: %d\n", count)
		
		// Mutex will be unlocked automatically
	}
	
	// Call increment multiple times
	for i := 0; i < 3; i++ {
		increment()
	}
}

// 📚 DEFER ORDER: Demonstrate LIFO execution
func deferOrder() {
	fmt.Println("📚 Demonstrating defer order (LIFO)")
	
	defer fmt.Println("🥉 Third: This runs third (first deferred)")
	defer fmt.Println("🥈 Second: This runs second (second deferred)")
	defer fmt.Println("🥇 First: This runs first (last deferred)")
	
	fmt.Println("📝 Main function body executes first")
	// When function returns:
	// 1. "First: This runs first" (last deferred)
	// 2. "Second: This runs second" (second deferred)  
	// 3. "Third: This runs third" (first deferred)
}

// 🎯 DEFER WITH PARAMETERS: Arguments evaluated immediately
func deferWithParams() {
	fmt.Println("🎯 Defer with parameters")
	
	x := 10
	defer fmt.Printf("🔢 Deferred: x = %d\n", x) // x evaluated as 10 immediately
	
	x = 20
	fmt.Printf("📝 Current: x = %d\n", x)
	
	// When function returns, deferred call prints x = 10 (not 20)
}

// 🎯 DEFER WITH CLOSURE: Capture variables by reference
func deferWithClosure() {
	fmt.Println("🎯 Defer with closure")
	
	x := 10
	defer func() {
		fmt.Printf("🔢 Deferred closure: x = %d\n", x) // x captured by reference
	}()
	
	x = 20
	fmt.Printf("📝 Current: x = %d\n", x)
	
	// When function returns, deferred closure prints x = 20
}

// 🚨 DEFER WITH PANIC: Defer runs even during panic
func deferWithPanic() {
	defer fmt.Println("🚨 Cleanup: This runs even during panic!")
	
	fmt.Println("📝 Before panic")
	panic("Something went wrong!")
	fmt.Println("📝 This line never executes")
}

// 🔄 DEFER IN LOOP: Common mistake demonstration
func deferInLoop() {
	fmt.Println("🔄 Defer in loop (potential issue)")
	
	// ❌ PROBLEMATIC: All defers execute when function returns
	for i := 0; i < 3; i++ {
		defer fmt.Printf("🔄 Loop defer: %d\n", i)
	}
	
	fmt.Println("📝 Loop completed")
	// All defers execute in reverse order: 2, 1, 0
}

// ✅ BETTER APPROACH: Use anonymous function
func deferInLoopFixed() {
	fmt.Println("✅ Defer in loop (fixed)")
	
	for i := 0; i < 3; i++ {
		func(index int) {
			defer fmt.Printf("✅ Fixed defer: %d\n", index)
			// Defer executes when anonymous function returns
		}(i)
	}
}

func main() {
	fmt.Println("⏰ DEFER STATEMENT TUTORIAL")
	fmt.Println("===========================")

	// 🎯 DEMO 1: Basic Defer Usage
	fmt.Println("\n🎯 DEMO 1: Basic Defer")
	fmt.Println("======================")
	
	defer fmt.Println("🎯 This will print last!")
	fmt.Println("📝 This prints first")
	fmt.Println("📝 This prints second")
	// "This will print last!" executes when main() returns

	// 🎯 DEMO 2: File Operations
	fmt.Println("\n🎯 DEMO 2: File Operations")
	fmt.Println("==========================")
	
	// Create a temporary file for demonstration
	tempFile := "/tmp/demo.txt"
	file, _ := os.Create(tempFile)
	file.WriteString("Hello, defer!")
	file.Close()
	
	readFileWithDefer(tempFile)
	os.Remove(tempFile) // Cleanup

	// 🎯 DEMO 3: Mutex with Defer
	fmt.Println("\n🎯 DEMO 3: Mutex with Defer")
	fmt.Println("===========================")
	
	safeCounter()

	// 🎯 DEMO 4: Defer Order (LIFO)
	fmt.Println("\n🎯 DEMO 4: Defer Order")
	fmt.Println("======================")
	
	deferOrder()

	// 🎯 DEMO 5: Parameters vs Closure
	fmt.Println("\n🎯 DEMO 5: Parameters vs Closure")
	fmt.Println("=================================")
	
	deferWithParams()
	fmt.Println()
	deferWithClosure()

	// 🎯 DEMO 6: Defer in Loop
	fmt.Println("\n🎯 DEMO 6: Defer in Loop")
	fmt.Println("========================")
	
	deferInLoop()
	fmt.Println()
	deferInLoopFixed()

	// 🎯 DEMO 7: Defer with Panic (commented to prevent crash)
	fmt.Println("\n🎯 DEMO 7: Defer with Panic")
	fmt.Println("===========================")
	
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("🚨 Recovered from panic: %v\n", r)
			}
		}()
		deferWithPanic()
	}()

	fmt.Println("\n✨ All defer demos completed!")
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

⏰ DEFER EXECUTION ORDER:
┌─────────────────────────────────────────────────────────────────────────┐
│ func example() {                                                        │
│     defer fmt.Println("First")   // Executes: 3rd (LIFO)               │
│     defer fmt.Println("Second")  // Executes: 2nd                       │
│     defer fmt.Println("Third")   // Executes: 1st (last deferred)       │
│     fmt.Println("Main")          // Executes: 1st (immediately)         │
│ }                                                                       │
│                                                                         │
│ Output:                                                                 │
│ Main                                                                    │
│ Third                                                                   │
│ Second                                                                  │
│ First                                                                   │
└─────────────────────────────────────────────────────────────────────────┘

🎯 DEFER PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Resource cleanup                                                     │
│ file, err := os.Open("file.txt")                                        │
│ if err != nil { return err }                                            │
│ defer file.Close()                                                      │
│                                                                         │
│ // Mutex unlock                                                         │
│ mu.Lock()                                                               │
│ defer mu.Unlock()                                                       │
│                                                                         │
│ // Timer measurement                                                    │
│ start := time.Now()                                                     │
│ defer func() {                                                          │
│     fmt.Printf("Duration: %v\n", time.Since(start))                     │
│ }()                                                                     │
│                                                                         │
│ // Panic recovery                                                       │
│ defer func() {                                                          │
│     if r := recover(); r != nil {                                       │
│         fmt.Printf("Recovered: %v\n", r)                                │
│     }                                                                   │
│ }()                                                                     │
└─────────────────────────────────────────────────────────────────────────┘

⚡ DEFER BEHAVIOR:
• Arguments evaluated immediately when defer is encountered
• Function call deferred until surrounding function returns
• Executes even if function panics
• LIFO order (stack-like behavior)

🚨 COMMON MISTAKES:
┌─────────────────────────────────────────────────────────────────────────┐
│ // ❌ Wrong: defer in loop                                               │
│ for i := 0; i < 10; i++ {                                               │
│     file, _ := os.Open(fmt.Sprintf("file%d.txt", i))                    │
│     defer file.Close() // All files stay open until function returns    │
│ }                                                                       │
│                                                                         │
│ // ✅ Correct: use anonymous function                                    │
│ for i := 0; i < 10; i++ {                                               │
│     func() {                                                            │
│         file, _ := os.Open(fmt.Sprintf("file%d.txt", i))                │
│         defer file.Close() // File closed each iteration                │
│         // use file                                                     │
│     }()                                                                 │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

💡 BEST PRACTICES:
• Use defer immediately after resource acquisition
• Keep deferred functions simple and fast
• Use anonymous functions to capture variables by reference
• Avoid defer in tight loops
• Always check errors in deferred functions when necessary

🎯 REAL-WORLD USES:
• File/database connection cleanup
• Mutex unlocking
• Timer measurement
• Panic recovery
• Logging function entry/exit
• Transaction rollback

⚡ PERFORMANCE NOTES:
• Defer has small overhead (~50ns per call)
• Avoid in performance-critical tight loops
• Consider manual cleanup for high-frequency operations
• Defer overhead is usually negligible compared to I/O operations

=============================================================================
*/