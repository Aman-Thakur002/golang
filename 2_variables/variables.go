/*
=============================================================================
                           📦 GO VARIABLES TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Variables are named storage locations that hold values.
Think of them as labeled boxes where you can store and retrieve data.

🔑 KEY FEATURES:
• Static typing (type is known at compile time)
• Type inference (Go can figure out the type)
• Zero values (variables have default values)
• Multiple declaration syntaxes

💡 REAL-WORLD ANALOGY:
Variable = Labeled Storage Box
- Name = Label on the box
- Type = What kind of items the box can hold
- Value = What's currently inside the box
- Declaration = Getting a new box and labeling it

🎯 WHY UNDERSTAND VARIABLES?
• Foundation of all programming
• Store and manipulate data
• Make programs dynamic and interactive
• Essential for any meaningful computation

=============================================================================
*/

package main

import "fmt"

func main() {
	fmt.Println("📦 VARIABLES LEARNING JOURNEY")
	fmt.Println("==============================")

	fmt.Println("\n🎯 VARIABLE DECLARATION METHODS")
	fmt.Println("================================")

	// 📝 METHOD 1: Full declaration with type and value
	var name1 string = "Aman"  // var name type = value
	fmt.Printf("📝 Method 1 (full): %s\n", name1)

	// 🎯 METHOD 2: Type inference - Go figures out the type
	var name2 = "Aman"  // Go infers this is a string
	fmt.Printf("🎯 Method 2 (inferred): %s\n", name2)

	// 🚀 METHOD 3: Shorthand syntax (most common)
	name3 := "Aman"  // := means declare and assign
	fmt.Printf("🚀 Method 3 (shorthand): %s\n", name3)

	fmt.Println("\n🔢 DIFFERENT DATA TYPES")
	fmt.Println("========================")

	// 🔢 NUMERIC TYPES
	var age int = 25
	var height float64 = 5.9
	var temperature float32 = 36.5
	
	fmt.Printf("🔢 Age (int): %d\n", age)
	fmt.Printf("🔢 Height (float64): %.1f\n", height)
	fmt.Printf("🔢 Temperature (float32): %.1f\n", temperature)

	// ✅ BOOLEAN TYPE
	var isStudent bool = true
	var isWorking bool = false
	
	fmt.Printf("✅ Is student: %t\n", isStudent)
	fmt.Printf("✅ Is working: %t\n", isWorking)

	// 📝 STRING TYPE
	var city string = "New York"
	var country string = "USA"
	
	fmt.Printf("📝 City: %s\n", city)
	fmt.Printf("📝 Country: %s\n", country)

	fmt.Println("\n🎯 ZERO VALUES")
	fmt.Println("===============")

	// 🔄 ZERO VALUES: Default values when not initialized
	var defaultInt int        // 0
	var defaultFloat float64  // 0.0
	var defaultBool bool      // false
	var defaultString string  // ""
	
	fmt.Printf("🔄 Default int: %d\n", defaultInt)
	fmt.Printf("🔄 Default float: %.1f\n", defaultFloat)
	fmt.Printf("🔄 Default bool: %t\n", defaultBool)
	fmt.Printf("🔄 Default string: '%s'\n", defaultString)

	fmt.Println("\n🎯 MULTIPLE VARIABLE DECLARATION")
	fmt.Println("=================================")

	// 📋 MULTIPLE VARIABLES: Declare several at once
	var (
		firstName string = "John"
		lastName  string = "Doe"
		userAge   int    = 30
	)
	
	fmt.Printf("📋 Full name: %s %s\n", firstName, lastName)
	fmt.Printf("📋 Age: %d\n", userAge)

	// 🚀 MULTIPLE SHORTHAND: Multiple variables with :=
	x, y, z := 10, 20, 30
	fmt.Printf("🚀 Multiple shorthand: x=%d, y=%d, z=%d\n", x, y, z)

	// 🔄 MULTIPLE ASSIGNMENT: Change multiple variables
	a, b := 1, 2
	fmt.Printf("🔄 Before swap: a=%d, b=%d\n", a, b)
	a, b = b, a  // Swap values
	fmt.Printf("🔄 After swap: a=%d, b=%d\n", a, b)

	fmt.Println("\n🎯 VARIABLE SCOPE")
	fmt.Println("==================")

	// 🏠 PACKAGE SCOPE: Variables declared outside functions
	// (We can't demonstrate this in main, but it's important to know)
	
	// 🏠 FUNCTION SCOPE: Variables declared inside functions
	functionVar := "I'm inside main function"
	fmt.Printf("🏠 Function scope: %s\n", functionVar)

	// 🏠 BLOCK SCOPE: Variables declared inside blocks
	if true {
		blockVar := "I'm inside this if block"
		fmt.Printf("🏠 Block scope: %s\n", blockVar)
		// blockVar is only accessible within this if block
	}
	// blockVar is not accessible here

	fmt.Println("\n🎯 CONSTANTS")
	fmt.Println("=============")

	// 🔒 CONSTANTS: Values that cannot be changed
	const pi float64 = 3.14159
	const greeting string = "Hello"
	const maxUsers int = 100
	
	fmt.Printf("🔒 Pi: %.5f\n", pi)
	fmt.Printf("🔒 Greeting: %s\n", greeting)
	fmt.Printf("🔒 Max users: %d\n", maxUsers)

	// 🔒 MULTIPLE CONSTANTS
	const (
		StatusActive   = "active"
		StatusInactive = "inactive"
		StatusPending  = "pending"
	)
	
	fmt.Printf("🔒 Status options: %s, %s, %s\n", StatusActive, StatusInactive, StatusPending)

	fmt.Println("\n🎯 TYPE CONVERSION")
	fmt.Println("===================")

	// 🔄 TYPE CONVERSION: Converting between types
	var intValue int = 42
	var floatValue float64 = float64(intValue)  // Convert int to float64
	var stringValue string = fmt.Sprintf("%d", intValue)  // Convert int to string
	
	fmt.Printf("🔄 Original int: %d\n", intValue)
	fmt.Printf("🔄 Converted to float: %.1f\n", floatValue)
	fmt.Printf("🔄 Converted to string: '%s'\n", stringValue)

	// 🎯 FINAL EXAMPLE: Using variables together
	fmt.Println("\n🎯 PUTTING IT ALL TOGETHER")
	fmt.Println("===========================")

	userName := "Alice"
	userScore := 95
	isPassing := userScore >= 60
	
	fmt.Printf("👤 Student: %s\n", userName)
	fmt.Printf("📊 Score: %d\n", userScore)
	fmt.Printf("✅ Passing: %t\n", isPassing)
	
	// String concatenation and formatting
	message := fmt.Sprintf("Congratulations %s! You scored %d points.", userName, userScore)
	fmt.Printf("🎉 %s\n", message)

	// Using the classic greeting
	fmt.Println("Hello, " + name3 + "!") // String concatenation
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

📦 VARIABLE DECLARATION SYNTAX:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Method 1: Full declaration                                           │
│ var name type = value                                                   │
│                                                                         │
│ // Method 2: Type inference                                             │
│ var name = value                                                        │
│                                                                         │
│ // Method 3: Shorthand (inside functions only)                         │
│ name := value                                                           │
│                                                                         │
│ // Multiple variables                                                   │
│ var a, b, c int = 1, 2, 3                                              │
│ x, y := 10, 20                                                          │
└─────────────────────────────────────────────────────────────────────────┘

🔢 BASIC DATA TYPES:
┌─────────────────┬─────────────────┬─────────────────────────────────────┐
│     Type        │   Zero Value    │           Examples                  │
├─────────────────┼─────────────────┼─────────────────────────────────────┤
│ int             │ 0               │ 42, -17, 1000                       │
│ float64         │ 0.0             │ 3.14, -2.5, 1.0                    │
│ bool            │ false           │ true, false                         │
│ string          │ ""              │ "Hello", "Go", ""                   │
│ byte            │ 0               │ 'A', 65, 0                         │
│ rune            │ 0               │ 'A', '世', 65                       │
└─────────────────┴─────────────────┴─────────────────────────────────────┘

🏠 VARIABLE SCOPE:
• Package scope: Declared outside functions, accessible throughout package
• Function scope: Declared inside functions, accessible within function
• Block scope: Declared inside {}, accessible within that block

🔒 CONSTANTS vs VARIABLES:
┌─────────────────┬─────────────────┬─────────────────────────────────────┐
│    Aspect       │   Variables     │            Constants                │
├─────────────────┼─────────────────┼─────────────────────────────────────┤
│ Mutability      │ Can change      │ Cannot change                       │
│ Declaration     │ var, :=         │ const                               │
│ Scope           │ Any scope       │ Any scope                           │
│ Evaluation      │ Runtime         │ Compile time                        │
└─────────────────┴─────────────────┴─────────────────────────────────────┘

⚡ IMPORTANT RULES:
• := can only be used inside functions
• Variables must be used (unused variables cause compile error)
• Go is statically typed (type is fixed at compile time)
• Zero values make variables safe to use immediately

🚨 GOTCHAS:
❌ := creates new variables (can shadow existing ones)
❌ Unused variables cause compilation errors
❌ Type conversion is explicit (no automatic conversion)
❌ Short variable declaration requires at least one new variable

💡 NAMING CONVENTIONS:
• Use camelCase for variable names
• Start with lowercase for private (package-level)
• Start with uppercase for public (exported)
• Use descriptive names (userName, not u)

🔧 BEST PRACTICES:
• Use := for most variable declarations inside functions
• Use var for zero values or when type is important
• Use const for values that never change
• Keep variable scope as narrow as possible
• Use meaningful names that describe the data

🎯 WHEN TO USE EACH METHOD:
• var name type = value: When type clarity is important
• var name = value: When you want explicit var keyword
• name := value: Most common, concise and clear
• const: For values that never change

=============================================================================
*/