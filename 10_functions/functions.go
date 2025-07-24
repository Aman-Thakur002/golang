/*
=============================================================================
                           🚀 GO FUNCTIONS TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Functions are reusable blocks of code that perform specific tasks.
Go functions are first-class citizens - they can be passed around like variables!

🔑 KEY FEATURES:
• Multiple return values (unique to Go!)
• Functions as parameters (higher-order functions)
• Functions returning functions (closures)
• Named return values
• Variadic functions (...args)

💡 REAL-WORLD ANALOGY:
Function = Kitchen Recipe
- Parameters = Ingredients you need
- Return values = What you get after cooking
- Function body = Cooking instructions

🎯 WHY USE FUNCTIONS?
• Code reusability and organization
• Break complex problems into smaller pieces
• Enable functional programming patterns
• Make code testable and maintainable

=============================================================================
*/

package main

import "fmt"
 // func add(a, b int) this means all parameters are f int type

// 🔢 BASIC FUNCTION: Takes parameters, returns single value
func add(a int, b int) int { // (parameters), return type
	return a + b  // Simple addition operation
}

// 🎯 MULTIPLE RETURN VALUES: Go's unique feature!
// Very common pattern in Go for returning result + error
func getLanguages()(string, string, bool){
	return "golang","javascript", true  // Return 3 values at once
	// Common pattern: (result, error) or (value1, value2, success)
}

// 🔄 HIGHER-ORDER FUNCTION: Function that takes another function as parameter
// This enables functional programming patterns
func processIt(f1 func(a int) int){
   // f1 is a function parameter that takes int and returns int
   // You can call f1(someValue) inside this function
}

// 🏭 FUNCTION FACTORY: Function that returns another function
// This creates closures - functions that "remember" their environment
func func2() func(a int) int {
	return func(a int) int {  // Anonymous function (lambda)
		return 4  // This inner function has access to outer function's variables
	}
}

func main() {
	// 🔢 CALLING BASIC FUNCTION
	result := add(2, 4)  // Pass arguments, get single return value
	fmt.Println(result)
	
	// 🎯 HANDLING MULTIPLE RETURN VALUES
	l1, l2, l3 := getLanguages()  // Unpack all 3 return values
	fmt.Println(l1, l2, l3)
	
	// 💡 ALTERNATIVE: Ignore some return values with _
	// lang1, _, success := getLanguages()  // Ignore middle value
	
	// 🏭 USING FUNCTION FACTORY
	// myFunc := func2()  // Get a function from func2
	// result2 := myFunc(10)  // Call the returned function
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

🔧 FUNCTION SYNTAX PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Basic function                                                       │
│ func name(param type) returnType { }                                    │
│                                                                         │
│ // Multiple parameters of same type                                     │
│ func add(a, b, c int) int { }                                           │
│                                                                         │
│ // Multiple return values                                               │
│ func divide(a, b int) (int, error) { }                                  │
│                                                                         │
│ // Named return values                                                  │
│ func calculate(x int) (result int, err error) { }                       │
│                                                                         │
│ // Variadic function                                                    │
│ func sum(nums ...int) int { }                                           │
│                                                                         │
│ // Function as parameter                                                │
│ func apply(f func(int) int, value int) int { }                          │
└─────────────────────────────────────────────────────────────────────────┘

🎯 MULTIPLE RETURN VALUES PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Common Go pattern: (result, error)                                   │
│ func readFile(name string) ([]byte, error)                              │
│                                                                         │
│ // Usage:                                                               │
│ data, err := readFile("config.txt")                                     │
│ if err != nil {                                                         │
│     // handle error                                                     │
│ }                                                                       │
│ // use data                                                             │
└─────────────────────────────────────────────────────────────────────────┘

🔄 FUNCTION TYPES & VARIABLES:
• Functions are first-class citizens in Go
• Can assign functions to variables
• Can pass functions as arguments
• Can return functions from functions

💡 COMMON PATTERNS:
• Error handling: (result, error)
• Optional values: (value, bool)
• Callbacks and event handlers
• Middleware functions
• Factory functions

🚨 GOTCHAS:
❌ Unused return values cause compile error (use _ to ignore)
❌ Functions are passed by value (copy)
❌ No function overloading (same name, different parameters)
❌ No default parameter values

🔧 BEST PRACTICES:
• Keep functions small and focused (single responsibility)
• Use multiple return values for error handling
• Use descriptive function names
• Return errors as last return value
• Use variadic functions for flexible APIs

🎯 WHEN TO USE EACH PATTERN:
• Single return: Simple calculations, getters
• Multiple returns: Operations that can fail, complex calculations
• Function parameters: Callbacks, strategy pattern, functional programming
• Function returns: Factory pattern, configuration, closures

=============================================================================
*/