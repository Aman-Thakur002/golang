/*
=============================================================================
                        🚨 GO PANIC & RECOVER TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Panic is Go's way of handling unrecoverable errors. Recover allows you to
regain control during a panic. Think of panic as throwing an exception
and recover as catching it.

🔑 KEY FEATURES:
• Panic stops normal execution and unwinds the stack
• Recover can only be called inside deferred functions
• Used for truly exceptional situations
• Prefer errors for normal error handling

💡 REAL-WORLD ANALOGY:
Panic/Recover = Emergency Procedures
- Panic = Fire alarm (something seriously wrong)
- Recover = Emergency response team (handle the crisis)
- Defer = Emergency exits (cleanup before evacuation)

🎯 WHEN TO USE?
• Unrecoverable programming errors
• Library initialization failures
• Critical system resource exhaustion
• NOT for normal error handling (use error type instead)

=============================================================================
*/

package main

import (
	"fmt"
	"runtime"
)

// 🚨 BASIC PANIC: Demonstrates simple panic
func basicPanic() {
	fmt.Println("📝 Before panic")
	panic("Something went terribly wrong!")
	fmt.Println("📝 This line never executes")
}

// 🛡️ PANIC WITH RECOVERY: Demonstrates recover
func panicWithRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("🛡️ Recovered from panic: %v\n", r)
		}
	}()
	
	fmt.Println("📝 Before panic")
	panic("Recoverable panic!")
	fmt.Println("📝 This line never executes")
}

// 🔢 DIVIDE BY ZERO: Practical example
func safeDivide(a, b float64) (result float64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("division error: %v", r)
			result = 0
		}
	}()
	
	if b == 0 {
		panic("division by zero")
	}
	
	result = a / b
	return result, nil
}

// 📋 ARRAY ACCESS: Index out of bounds
func safeArrayAccess(arr []int, index int) (value int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("array access error: %v", r)
			value = -1
		}
	}()
	
	value = arr[index] // This might panic if index is out of bounds
	return value, nil
}

// 🏗️ NESTED PANIC: Multiple panic levels
func nestedPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("🏗️ Outer recover: %v\n", r)
			// You can panic again from a recover
			// panic("Panic from recover!")
		}
	}()
	
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("🏗️ Inner recover: %v\n", r)
				panic("Re-panic from inner function")
			}
		}()
		
		panic("Original panic")
	}()
}

// 📊 STACK TRACE: Get stack information during panic
func panicWithStackTrace() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("📊 Panic: %v\n", r)
			
			// Print stack trace
			buf := make([]byte, 1024)
			n := runtime.Stack(buf, false)
			fmt.Printf("📊 Stack trace:\n%s", buf[:n])
		}
	}()
	
	innerFunction()
}

func innerFunction() {
	panic("Panic from inner function")
}

// 🎯 CUSTOM PANIC TYPE: Using structured panic data
type CustomError struct {
	Code    int
	Message string
	Context map[string]interface{}
}

func (e CustomError) String() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func customPanicExample() {
	defer func() {
		if r := recover(); r != nil {
			if customErr, ok := r.(CustomError); ok {
				fmt.Printf("🎯 Custom error recovered:\n")
				fmt.Printf("   Code: %d\n", customErr.Code)
				fmt.Printf("   Message: %s\n", customErr.Message)
				fmt.Printf("   Context: %v\n", customErr.Context)
			} else {
				fmt.Printf("🎯 Unknown panic type: %v\n", r)
			}
		}
	}()
	
	panic(CustomError{
		Code:    500,
		Message: "Internal server error",
		Context: map[string]interface{}{
			"user_id": 12345,
			"action":  "process_payment",
		},
	})
}

// 🔄 PANIC IN GOROUTINE: Goroutine panics don't affect main
func panicInGoroutine() {
	fmt.Println("🔄 Starting goroutine that will panic")
	
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("🔄 Goroutine recovered: %v\n", r)
			}
		}()
		
		panic("Panic in goroutine")
	}()
	
	// Give goroutine time to panic and recover
	fmt.Println("🔄 Main function continues normally")
}

// ✅ PROPER ERROR HANDLING: When NOT to use panic
func properErrorHandling(filename string) error {
	// ✅ GOOD: Return error instead of panic
	if filename == "" {
		return fmt.Errorf("filename cannot be empty")
	}
	
	// ❌ BAD: Don't panic for expected errors
	// if filename == "" {
	//     panic("filename cannot be empty")
	// }
	
	fmt.Printf("✅ Processing file: %s\n", filename)
	return nil
}

func main() {
	fmt.Println("🚨 PANIC & RECOVER TUTORIAL")
	fmt.Println("============================")

	// 🎯 DEMO 1: Basic Panic with Recovery
	fmt.Println("\n🎯 DEMO 1: Basic Recovery")
	fmt.Println("=========================")
	
	panicWithRecover()
	fmt.Println("✅ Program continues after recovery")

	// 🎯 DEMO 2: Safe Division
	fmt.Println("\n🎯 DEMO 2: Safe Division")
	fmt.Println("========================")
	
	result, err := safeDivide(10, 2)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Printf("✅ Result: %.2f\n", result)
	}
	
	result, err = safeDivide(10, 0)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Printf("✅ Result: %.2f\n", result)
	}

	// 🎯 DEMO 3: Safe Array Access
	fmt.Println("\n🎯 DEMO 3: Safe Array Access")
	fmt.Println("============================")
	
	arr := []int{1, 2, 3, 4, 5}
	
	value, err := safeArrayAccess(arr, 2)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Printf("✅ Value at index 2: %d\n", value)
	}
	
	value, err = safeArrayAccess(arr, 10)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Printf("✅ Value at index 10: %d\n", value)
	}

	// 🎯 DEMO 4: Nested Panic
	fmt.Println("\n🎯 DEMO 4: Nested Panic")
	fmt.Println("=======================")
	
	nestedPanic()

	// 🎯 DEMO 5: Custom Panic Type
	fmt.Println("\n🎯 DEMO 5: Custom Panic Type")
	fmt.Println("============================")
	
	customPanicExample()

	// 🎯 DEMO 6: Stack Trace
	fmt.Println("\n🎯 DEMO 6: Stack Trace")
	fmt.Println("======================")
	
	panicWithStackTrace()

	// 🎯 DEMO 7: Proper Error Handling
	fmt.Println("\n🎯 DEMO 7: Proper Error Handling")
	fmt.Println("=================================")
	
	if err := properErrorHandling("config.txt"); err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	}
	
	if err := properErrorHandling(""); err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	}

	fmt.Println("\n✨ All panic/recover demos completed!")
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

🚨 PANIC BEHAVIOR:
┌─────────────────────────────────────────────────────────────────────────┐
│ 1. Panic stops normal execution of current function                     │
│ 2. Deferred functions still execute (in LIFO order)                     │
│ 3. Stack unwinds to calling function                                    │
│ 4. Process repeats until recover() is called or program terminates      │
│ 5. If not recovered, program prints stack trace and exits               │
└─────────────────────────────────────────────────────────────────────────┘

🛡️ RECOVER RULES:
┌─────────────────────────────────────────────────────────────────────────┐
│ // ✅ Correct: recover in deferred function                              │
│ defer func() {                                                          │
│     if r := recover(); r != nil {                                       │
│         // Handle panic                                                 │
│     }                                                                   │
│ }()                                                                     │
│                                                                         │
│ // ❌ Wrong: recover outside defer                                       │
│ if r := recover(); r != nil {  // Always returns nil                    │
│     // This never executes                                             │
│ }                                                                       │
│                                                                         │
│ // ❌ Wrong: recover in normal function call                             │
│ func handlePanic() {                                                    │
│     recover() // Returns nil, doesn't work                             │
│ }                                                                       │
│ defer handlePanic()                                                     │
└─────────────────────────────────────────────────────────────────────────┘

🎯 PANIC vs ERROR:
┌─────────────────┬─────────────────┬─────────────────────────────────────┐
│   Situation     │   Use Error     │           Use Panic                 │
├─────────────────┼─────────────────┼─────────────────────────────────────┤
│ File not found  │ ✅ Expected     │ ❌ Too harsh                        │
│ Network timeout │ ✅ Recoverable  │ ❌ Normal occurrence                │
│ Invalid input   │ ✅ User error   │ ❌ Should validate                  │
│ Out of memory   │ ❌ Hard to fix  │ ✅ System failure                   │
│ Array bounds    │ ❌ Preventable  │ ✅ Programming error                │
│ Nil pointer     │ ❌ Preventable  │ ✅ Programming error                │
└─────────────────┴─────────────────┴─────────────────────────────────────┘

🔄 PANIC SOURCES:
• Runtime errors (nil pointer, index out of bounds)
• Explicit panic() calls
• Type assertions on wrong types
• Closing closed channels
• Writing to closed channels

💡 BEST PRACTICES:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Use errors for expected failures                                     │
│ func ReadFile(name string) ([]byte, error) {                            │
│     if name == "" {                                                     │
│         return nil, errors.New("filename required")                     │
│     }                                                                   │
│     // ...                                                              │
│ }                                                                       │
│                                                                         │
│ // Use panic for programming errors                                     │
│ func ProcessArray(arr []int, index int) int {                           │
│     if index < 0 || index >= len(arr) {                                 │
│         panic("index out of bounds")                                    │
│     }                                                                   │
│     return arr[index]                                                   │
│ }                                                                       │
│                                                                         │
│ // Always recover in libraries                                          │
│ func LibraryFunction() (result int, err error) {                        │
│     defer func() {                                                      │
│         if r := recover(); r != nil {                                   │
│             err = fmt.Errorf("library error: %v", r)                    │
│         }                                                               │
│     }()                                                                 │
│     // risky operations                                                 │
│     return result, nil                                                  │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

🚨 GOTCHAS:
❌ Recover only works in deferred functions
❌ Panic in goroutine doesn't affect main (unless recovered)
❌ Don't use panic for flow control
❌ Don't ignore recovered panics without logging
❌ Panic can carry any type, not just strings

⚡ PERFORMANCE IMPACT:
• Panic/recover is expensive (avoid in hot paths)
• Stack unwinding takes time
• Use errors for performance-critical code
• Panic should be exceptional, not routine

🎯 WHEN TO RECOVER:
• Web servers (don't crash on single request)
• Libraries (convert panics to errors)
• Worker pools (isolate worker failures)
• Plugin systems (isolate plugin failures)

=============================================================================
*/