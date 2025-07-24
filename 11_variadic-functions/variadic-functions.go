/*
=============================================================================
                        🔢 GO VARIADIC FUNCTIONS TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Variadic functions can accept any number of parameters of the same type.
Think of them as "flexible functions" that adapt to different argument counts.

🔑 KEY FEATURES:
• Accept zero or more arguments of same type
• Use ... (ellipsis) syntax
• Arguments become a slice inside function
• Can mix regular and variadic parameters

💡 REAL-WORLD ANALOGY:
Variadic Function = Restaurant Buffet
- Regular function = Fixed meal (appetizer, main, dessert)
- Variadic function = Buffet (take as much as you want)
- ...args = "Help yourself to any amount"

🎯 WHY USE VARIADIC FUNCTIONS?
• Flexible APIs (fmt.Printf, append, etc.)
• Reduce function overloading
• Clean, readable function calls
• Handle unknown number of inputs

=============================================================================
*/

package main

import "fmt"
  // functions which can accept any number of parameters

  // 🔢 BASIC VARIADIC FUNCTION: Accepts any number of integers
  func sum(nums ...int) int {  // ...int means "zero or more int arguments"
	total := 0
	// nums becomes a []int slice inside the function
	for _, num := range nums{  // Iterate over all passed arguments
		total = total + num
	}

	return total
  }

// 🎯 MIXED TYPES VARIADIC: Accepts any number of any type
// to receive anytype of parameters
func fun1(para ...interface{}) interface{}{  // or use "any" keyword
	fmt.Print("Mixed values: ")
	for i, item := range para {  // para is []interface{} slice
		if i > 0 {
			fmt.Print(", ")  // Add comma between items
		}
		fmt.Printf("%v", item)  // %v prints any type
	}
	fmt.Println()
	return len(para) // return count of parameters
}

// 🎯 MIXED REGULAR + VARIADIC: Regular params first, variadic last
func greetPeople(greeting string, names ...string) {
	for _, name := range names {
		fmt.Printf("%s, %s!\n", greeting, name)
	}
}

// 🔧 PRACTICAL EXAMPLE: Logger with different levels
func log(level string, messages ...string) {
	fmt.Printf("[%s] ", level)
	for i, msg := range messages {
		if i > 0 {
			fmt.Print(" | ")
		}
		fmt.Print(msg)
	}
	fmt.Println()
}

func main() {
	fmt.Println("🔢 VARIADIC FUNCTIONS LEARNING JOURNEY")
	fmt.Println("======================================")

	fmt.Println("\n🎯 BASIC VARIADIC USAGE")
	fmt.Println("=======================")

	// 🔢 CALLING WITH DIFFERENT ARGUMENT COUNTS
	fmt.Println("Sum of no numbers:", sum())           // 0 arguments
	fmt.Println("Sum of one number:", sum(5))          // 1 argument  
	fmt.Println("Sum of multiple:", sum(1, 2, 3, 4))   // 4 arguments

	fmt.Println("\n🎯 SLICE EXPANSION")
	fmt.Println("==================")

	// 📋 SLICE EXPANSION: Use ... to pass slice as variadic args
	nums := []int{2,4,23,45,13,2,23}
	result := sum(nums ...)  // ... expands slice into individual arguments
	fmt.Println("Sum from slice:", result)

	fmt.Println("\n🎯 MIXED TYPES VARIADIC")
	fmt.Println("=======================")

	// Test fun1 with mixed types
	count := fun1(1, "hello", 3.14, true, []int{1,2,3})
	fmt.Printf("Total parameters: %v\n", count)

	fmt.Println("\n🎯 REGULAR + VARIADIC PARAMETERS")
	fmt.Println("=================================")

	// 👋 MIXED PARAMETERS: greeting is regular, names is variadic
	greetPeople("Hello", "Alice", "Bob", "Charlie")
	greetPeople("Hi")  // No variadic args - that's OK!

	fmt.Println("\n🎯 PRACTICAL EXAMPLE: LOGGER")
	fmt.Println("=============================")

	// 📝 LOGGER EXAMPLE: Different message counts
	log("INFO", "Application started")
	log("ERROR", "Database connection failed", "Retrying in 5 seconds")
	log("DEBUG", "User login", "Session created", "Redirecting to dashboard")
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

🔢 VARIADIC FUNCTION SYNTAX:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Basic variadic function                                              │
│ func name(args ...type) returnType {                                    │
│     // args is a slice of type []type                                   │
│ }                                                                       │
│                                                                         │
│ // Mixed regular and variadic                                           │
│ func name(regular type, args ...type) returnType {                      │
│     // regular param first, variadic last                              │
│ }                                                                       │
│                                                                         │
│ // Calling with slice expansion                                         │
│ slice := []int{1, 2, 3}                                                 │
│ result := name(slice...)  // ... expands slice                         │
└─────────────────────────────────────────────────────────────────────────┘

⚡ VARIADIC RULES:
• Only last parameter can be variadic
• Inside function, variadic parameter becomes a slice
• Can pass zero or more arguments
• Use ... to expand slice into arguments

🎯 COMMON PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Standard library examples                                            │
│ fmt.Printf(format, args...)     // Formatted printing                   │
│ append(slice, elements...)      // Add multiple elements                │
│ max(numbers...)                 // Find maximum of many numbers         │
│                                                                         │
│ // Custom examples                                                      │
│ sum(numbers...)                 // Add multiple numbers                 │
│ concat(strings...)              // Join multiple strings                │
│ log(level, messages...)         // Log multiple messages               │
└─────────────────────────────────────────────────────────────────────────┘

🔍 VARIADIC vs SLICE PARAMETER:
┌─────────────────┬─────────────────┬─────────────────────────────────────┐
│    Aspect       │   Variadic      │           Slice                     │
├─────────────────┼─────────────────┼─────────────────────────────────────┤
│ Declaration     │ func(args...)   │ func(slice []type)                  │
│ Calling         │ f(1, 2, 3)      │ f([]int{1, 2, 3})                   │
│ Flexibility     │ Variable args   │ Must pass slice                     │
│ Expansion       │ f(slice...)     │ f(slice)                            │
└─────────────────┴─────────────────┴─────────────────────────────────────┘

🚨 GOTCHAS:
❌ Only last parameter can be variadic
❌ Can't have multiple variadic parameters
❌ Variadic parameter is always a slice (even with 0 args)
❌ Must use ... when expanding slice to variadic

💡 MEMORY CONSIDERATIONS:
• Variadic arguments create a new slice each call
• For performance-critical code, consider slice parameters
• Empty variadic call creates empty slice (not nil)

🔧 BEST PRACTICES:
• Use variadic for user-friendly APIs
• Combine with regular parameters for context
• Document expected argument types clearly
• Consider performance implications for hot paths

🎯 WHEN TO USE VARIADIC:
✅ Building flexible APIs (like fmt.Printf)
✅ Mathematical operations (sum, max, min)
✅ Collection operations (append, merge)
✅ Logging and debugging functions

❌ When you always need exact number of args
❌ When performance is critical (use slices)
❌ When arguments have different meanings

=============================================================================
*/