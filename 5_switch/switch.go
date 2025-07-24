/*
=============================================================================
                           🔀 GO SWITCH TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Switch statements provide a clean way to execute different code blocks
based on different conditions. Much cleaner than long if-else chains!

🔑 KEY FEATURES:
• No automatic fallthrough (unlike C/Java)
• Multiple values in one case
• Type switches for interface{}
• Switch without expression (like if-else)

💡 REAL-WORLD ANALOGY:
Switch = Traffic Control System
- Expression = Incoming traffic direction
- Cases = Different traffic light patterns
- Default = What to do when no pattern matches
- No fallthrough = Each direction gets exactly one action

🎯 WHY USE SWITCH?
• Cleaner than long if-else chains
• More readable for multiple conditions
• Efficient compilation (often jump tables)
• Type switching for interfaces

=============================================================================
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("🔀 SWITCH LEARNING JOURNEY")
	fmt.Println("===========================")

	fmt.Println("\n🎯 BASIC SWITCH")
	fmt.Println("================")

	// 🎯 SIMPLE SWITCH: No need for break keyword!
	i := 8
	switch i {
	case 1:
		fmt.Println("🥇 One")
	case 2:
		fmt.Println("🥈 Two")
	case 3:
		fmt.Println("🥉 Three")
	case 4:
		fmt.Println("4️⃣ Four")
	case 5:
		fmt.Println("5️⃣ Five")
	default:  // 💡 DEFAULT: Executes when no case matches
		fmt.Println("🔢 Number is greater than five")
	}

	fmt.Println("\n🎯 MULTIPLE CONDITIONS")
	fmt.Println("======================")

	// 🔄 MULTIPLE CONDITION SWITCH: One case, multiple values
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:  // 💡 COMMA-SEPARATED: Multiple values in one case
		fmt.Println("🎉 It's weekend! Time to relax!")
	case time.Monday:
		fmt.Println("😴 Monday blues...")
	case time.Friday:
		fmt.Println("🎊 TGIF! Weekend is coming!")
	default:
		fmt.Println("💼 It's a work day")
	}

	fmt.Println("\n🎯 TYPE SWITCH")
	fmt.Println("===============")

	// 🔍 TYPE SWITCH: Check the type of interface{} value
	whoAmI := func(i interface{}) {  // interface{} can hold any type
		switch i.(type) {  // 💡 TYPE ASSERTION: i.(type) gets the actual type
		case int:
			fmt.Println("🔢 I'm an Integer!")
		case bool:
			fmt.Println("✅ I'm a Boolean!")
		case string:
			fmt.Println("📝 I'm a String!")
		case float64:
			fmt.Println("🔢 I'm a Float!")
		default:
			fmt.Println("❓ I'm some other type")
		}
	}

	// 🧪 TEST TYPE SWITCH with different types
	whoAmI(1)           // int
	whoAmI(true)        // bool
	whoAmI("hello")     // string
	whoAmI(3.14)        // float64
	whoAmI([]int{1, 2}) // slice (default case)

	fmt.Println("\n🎯 SWITCH WITH EXPRESSIONS")
	fmt.Println("===========================")

	// 🎯 SWITCH WITH EXPRESSIONS: Each case can be an expression
	age := 25
	switch {  // 💡 NO EXPRESSION: Acts like if-else chain
	case age < 13:
		fmt.Println("👶 Child")
	case age < 20:
		fmt.Println("👦 Teenager")
	case age < 60:
		fmt.Println("👨 Adult")
	default:
		fmt.Println("👴 Senior")
	}

	fmt.Println("\n🎯 SWITCH WITH INITIALIZATION")
	fmt.Println("==============================")

	// 🎯 SWITCH WITH INITIALIZATION: Declare variable in switch statement
	switch day := time.Now().Weekday(); day {  // 💡 INITIALIZATION: day := ... ; day
	case time.Monday:
		fmt.Printf("📅 Today is %v - Start of work week\n", day)
	case time.Wednesday:
		fmt.Printf("📅 Today is %v - Hump day!\n", day)
	case time.Friday:
		fmt.Printf("📅 Today is %v - Almost weekend!\n", day)
	default:
		fmt.Printf("📅 Today is %v\n", day)
	}

	fmt.Println("\n🎯 ADVANCED TYPE SWITCH")
	fmt.Println("========================")

	// 🔍 ADVANCED TYPE SWITCH: Capture the value with type assertion
	describeValue := func(x interface{}) {
		switch v := x.(type) {  // 💡 v gets the actual typed value
		case int:
			fmt.Printf("🔢 Integer: %d (doubled: %d)\n", v, v*2)
		case string:
			fmt.Printf("📝 String: '%s' (length: %d)\n", v, len(v))
		case bool:
			fmt.Printf("✅ Boolean: %t (negated: %t)\n", v, !v)
		default:
			fmt.Printf("❓ Unknown type: %T with value: %v\n", v, v)
		}
	}

	describeValue(42)
	describeValue("Hello Go!")
	describeValue(true)
	describeValue(3.14159)
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

🔀 SWITCH SYNTAX PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Basic switch                                                         │
│ switch expression {                                                     │
│ case value1:                                                            │
│     // code                                                             │
│ case value2, value3:  // multiple values                               │
│     // code                                                             │
│ default:                                                                │
│     // code                                                             │
│ }                                                                       │
│                                                                         │
│ // Switch without expression (like if-else)                             │
│ switch {                                                                │
│ case condition1:                                                        │
│     // code                                                             │
│ case condition2:                                                        │
│     // code                                                             │
│ }                                                                       │
│                                                                         │
│ // Type switch                                                          │
│ switch v := x.(type) {                                                  │
│ case int:                                                               │
│     // v is int                                                         │
│ case string:                                                            │
│     // v is string                                                      │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

🔄 SWITCH vs IF-ELSE COMPARISON:
┌─────────────────┬─────────────────┬─────────────────────────────────────┐
│    Aspect       │     Switch      │              If-Else                │
├─────────────────┼─────────────────┼─────────────────────────────────────┤
│ Readability     │ Better          │ Gets messy with many conditions     │
│ Performance     │ Often faster    │ Sequential evaluation               │
│ Fallthrough     │ No (explicit)   │ N/A                                 │
│ Multiple Values │ Easy            │ Requires || operators               │
│ Type Checking   │ Built-in        │ Requires type assertions            │
└─────────────────┴─────────────────┴─────────────────────────────────────┘

⚡ SWITCH CHARACTERISTICS:
• No automatic fallthrough (unlike C/Java)
• Cases don't need to be constants (can be expressions)
• Default case is optional
• Can switch on any comparable type
• Empty switch{} acts like if-else chain

🎯 TYPE SWITCH DETAILS:
• Works only with interface{} or interface types
• x.(type) syntax only valid in switch statements
• Can capture typed value: switch v := x.(type)
• Useful for handling different types in generic code

🚨 GOTCHAS:
❌ No automatic fallthrough (use fallthrough keyword if needed)
❌ Type switch only works with interfaces
❌ Can't use := in case expressions (use switch init; expr)
❌ Default case can be anywhere, not just at the end

💡 PERFORMANCE NOTES:
• Compiler may optimize switch to jump table
• Generally faster than equivalent if-else chain
• Type switches are efficient for interface handling

🔧 BEST PRACTICES:
• Use switch for 3+ conditions instead of if-else
• Put most common cases first
• Use type switches for interface{} handling
• Keep case logic simple (extract to functions if complex)
• Use meaningful variable names in type switches

🎯 WHEN TO USE SWITCH:
✅ Multiple discrete values to check
✅ Type checking on interfaces
✅ State machines
✅ Command parsing
✅ Replacing long if-else chains

❌ When you need complex boolean logic
❌ When conditions are ranges (use if-else)
❌ When you need fallthrough behavior (rare in Go)

=============================================================================
*/