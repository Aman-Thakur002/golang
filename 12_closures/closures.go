/*
=============================================================================
                           🔒 GO CLOSURES TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Closures are functions that "remember" variables from their surrounding scope.
Think of them as functions with "memory" - they capture and keep access to variables.

🔑 KEY FEATURES:
• Functions that reference variables from outer scope
• Variables remain alive even after outer function returns
• Each closure has its own copy of captured variables
• Enable powerful functional programming patterns

💡 REAL-WORLD ANALOGY:
Closure = Personal Assistant with Memory
- Function = Assistant who can do tasks
- Captured variables = Personal notebook the assistant keeps
- Multiple closures = Multiple assistants, each with their own notebook
- Memory persists = Notebook survives even when you're not around

🎯 WHY USE CLOSURES?
• Create specialized functions with built-in state
• Implement counters, accumulators, and generators
• Build function factories and configuration
• Enable elegant functional programming patterns

=============================================================================
*/

package main

import "fmt"

// 🏭 BASIC CLOSURE: Function factory that creates counter functions
func counter() func() int {
	var count int = 0  // 📝 CAPTURED VARIABLE: This will be "remembered"
	
	// 🔒 CLOSURE: This inner function captures 'count' from outer scope
	return func() int {
		count += 1     // 💡 ACCESSES outer variable - this is the closure!
		return count   // Each call increments and returns the same 'count'
	}
}

// 🎯 CLOSURE WITH PARAMETERS: More flexible closure factory
func makeAdder(x int) func(int) int {
	// 📝 CAPTURED: x is captured from outer function
	return func(y int) int {
		return x + y  // 🔒 CLOSURE: Uses captured 'x' plus parameter 'y'
	}
}

// 🔧 PRACTICAL EXAMPLE: Configuration closure
func makeMultiplier(factor int) func(int) int {
	return func(value int) int {
		return value * factor  // 📝 'factor' is captured and remembered
	}
}

// 🎯 ADVANCED: Closure with multiple captured variables
func makeCalculator(operation string) func(int, int) int {
	return func(a, b int) int {
		switch operation {  // 📝 'operation' is captured
		case "add":
			return a + b
		case "multiply":
			return a * b
		case "subtract":
			return a - b
		case "divide":
			if b != 0 {
				return a / b
			}
			return 0
		default:
			return 0
		}
	}
}

func main() {
	fmt.Println("🔒 CLOSURES LEARNING JOURNEY")
	fmt.Println("============================")

	fmt.Println("\n🎯 BASIC CLOSURE: COUNTER")
	fmt.Println("=========================")

	// 🏭 CREATE CLOSURE: Each call to counter() creates a new closure
	increment := counter()  // increment "remembers" its own count variable
	fmt.Println("First call:", increment())   // 1 - count starts at 0, becomes 1
	fmt.Println("Second call:", increment())  // 2 - same count variable, now becomes 2

	fmt.Println("\n🎯 MULTIPLE INDEPENDENT CLOSURES")
	fmt.Println("=================================")

	// 🔄 INDEPENDENT CLOSURES: Each has its own captured variables
	counter1 := counter()  // counter1 has its own 'count'
	counter2 := counter()  // counter2 has its own separate 'count'

	fmt.Println("Counter1 first:", counter1())   // 1
	fmt.Println("Counter1 second:", counter1())  // 2
	fmt.Println("Counter2 first:", counter2())   // 1 (independent!)
	fmt.Println("Counter1 third:", counter1())   // 3

	fmt.Println("\n🎯 CLOSURE WITH PARAMETERS")
	fmt.Println("===========================")

	// 🎯 ADDER CLOSURES: Each captures different 'x' value
	add5 := makeAdder(5)   // Captures x=5
	add10 := makeAdder(10) // Captures x=10

	fmt.Println("add5(3):", add5(3))   // 5 + 3 = 8
	fmt.Println("add10(3):", add10(3)) // 10 + 3 = 13

	fmt.Println("\n🎯 PRACTICAL EXAMPLE: MULTIPLIERS")
	fmt.Println("==================================")

	// 🔧 MULTIPLIER CLOSURES: Useful for scaling operations
	double := makeMultiplier(2)
	triple := makeMultiplier(3)
	
	fmt.Println("Double 7:", double(7))   // 7 * 2 = 14
	fmt.Println("Triple 7:", triple(7))   // 7 * 3 = 21

	fmt.Println("\n🎯 ADVANCED: CALCULATOR CLOSURES")
	fmt.Println("=================================")

	// 🎯 CALCULATOR CLOSURES: Each captures different operation
	adder := makeCalculator("add")
	multiplier := makeCalculator("multiply")
	divider := makeCalculator("divide")
	
	fmt.Println("Add 5 + 3:", adder(5, 3))        // 8
	fmt.Println("Multiply 5 * 3:", multiplier(5, 3)) // 15
	fmt.Println("Divide 6 / 2:", divider(6, 2))    // 3
	fmt.Println("Divide 5 / 0:", divider(5, 0))    // 0 (safe division)
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

🔒 CLOSURE MECHANICS:
┌─────────────────────────────────────────────────────────────────────────┐
│ func outer() func() int {                                               │
│     captured := 42        // This variable will be captured            │
│     return func() int {                                                 │
│         captured++        // Inner function accesses outer variable    │
│         return captured   // This creates a closure                    │
│     }                                                                   │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

🎯 CLOSURE CHARACTERISTICS:
• Inner function + captured variables = Closure
• Captured variables survive outer function's return
• Each closure instance has independent captured variables
• Variables are captured by reference, not value

🔍 CLOSURE vs REGULAR FUNCTION:
┌─────────────────┬─────────────────┬─────────────────────────────────────┐
│    Aspect       │   Regular Func  │            Closure                  │
├─────────────────┼─────────────────┼────────────────���────────────────────┤
│ Variable Access │ Parameters only │ Parameters + captured variables     │
│ State           │ Stateless       │ Can maintain state                  │
│ Memory          │ No extra memory │ Keeps captured variables alive      │
│ Independence    │ All identical   │ Each instance independent           │
└─────────────────┴─────────────────┴─────────────────────────────────────┘

💡 COMMON CLOSURE PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Counter pattern                                                      │
│ func makeCounter() func() int  //returns incrementing function          |
│                                                                         │
│ // Configuration pattern                                                │
│ func makeHandler(config Config) http.HandlerFunc {}                     │
│                                                                         │
│ // Accumulator pattern                                                  │
│ func makeAccumulator() func(int) int { }                                │
│                                                                         │
│ // Filter/Transform pattern                                             │
│ func makeFilter(criteria func(T) bool) func([]T) []T { }                │
└─────────────────────────────────────────────────────────────────────────┘

🚨 GOTCHAS:
❌ Captured variables are shared by reference
❌ Loop variables in closures (common trap)
❌ Memory leaks if closures hold large objects
❌ Goroutine closures capturing loop variables

💡 MEMORY IMPLICATIONS:
• Closures keep captured variables alive
• Can prevent garbage collection
• Each closure instance uses separate memory
• Be careful with large captured objects

🔧 BEST PRACTICES:
• Use closures for stateful function factories
• Prefer closures over global variables for state
• Be mindful of memory usage with captured variables
• Use closures for configuration and specialization

🎯 WHEN TO USE CLOSURES:
✅ Creating specialized functions (adders, multipliers)
✅ Maintaining state without global variables
✅ Event handlers with context
✅ Functional programming patterns
✅ Configuration and factory patterns

❌ When simple functions suffice
❌ When memory usage is critical
❌ When state management is complex (use structs)

🔄 LOOP VARIABLE CLOSURE TRAP:
// ❌ Wrong - all closures capture same variable
for i := 0; i < 3; i++ {
    go func() { fmt.Println(i) }() // All print 3!
}

// ✅ Correct - capture loop variable properly
for i := 0; i < 3; i++ {
    go func(val int) { fmt.Println(val) }(i)
}

=============================================================================
*/