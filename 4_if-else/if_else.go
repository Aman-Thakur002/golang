/*
=============================================================================
                           🔀 GO IF-ELSE TUTORIAL
=============================================================================

📚 CORE CONCEPT:
If-else statements control program flow based on conditions.
They let your program make decisions and execute different code paths.

🔑 KEY FEATURES:
• No parentheses required around conditions
• Braces are mandatory (even for single statements)
• Can declare variables in if statement
• Supports else if for multiple conditions

💡 REAL-WORLD ANALOGY:
If-else = Decision Tree
- Condition = Question you ask
- If block = What to do if answer is "yes"
- Else block = What to do if answer is "no"
- Else if = Additional questions to ask

🎯 WHY USE IF-ELSE?
• Control program flow based on conditions
• Handle different scenarios in your code
• Validate input and handle errors
• Implement business logic

=============================================================================
*/

package main

import "fmt"

func main() {
	fmt.Println("🔀 IF-ELSE LEARNING JOURNEY")
	fmt.Println("============================")

	fmt.Println("\n🎯 BASIC IF-ELSE")
	fmt.Println("=================")

	age := 17

	// 🎯 BASIC IF-ELSE: Simple condition checking
	if age >= 18 {  // 💡 NO PARENTHESES: Go doesn't require () around condition
		fmt.Println("🎉 Adult - You can vote!")
	} else {
		fmt.Println("👶 Minor - Wait a bit more!")
	}

	fmt.Println("\n🎯 ELSE IF CHAIN")
	fmt.Println("=================")

	// 🔄 ELSE IF: Multiple conditions (same as other languages)
	score := 85
	if score >= 90 {
		fmt.Println("🏆 Grade A - Excellent!")
	} else if score >= 80 {
		fmt.Println("🥈 Grade B - Good job!")
	} else if score >= 70 {
		fmt.Println("🥉 Grade C - Not bad!")
	} else if score >= 60 {
		fmt.Println("📚 Grade D - Need improvement!")
	} else {
		fmt.Println("❌ Grade F - Study harder!")
	}

	fmt.Println("\n🎯 IF WITH INITIALIZATION")
	fmt.Println("==========================")

	// 🎯 VARIABLE DECLARATION IN IF: Declare and use in same statement
	// Variable scope is limited to the if-else block
	if age := 11; age >= 18 {  // 💡 INITIALIZATION: age := 11; condition
		fmt.Println("🎉 Adult")
	} else if age >= 12 {
		fmt.Printf("👦 Age: %d - Pre-teen\n", age)
	} else {
		fmt.Printf("👶 Age: %d - Child\n", age)
	}
	// 💡 NOTE: 'age' variable from if statement is not accessible here

	fmt.Println("\n🎯 COMPLEX CONDITIONS")
	fmt.Println("=====================")

	// 🔗 LOGICAL OPERATORS: && (AND), || (OR), ! (NOT)
	temperature := 25
	isRaining := false

	if temperature > 20 && !isRaining {  // 💡 AND + NOT operators
		fmt.Println("🌞 Perfect weather for a walk!")
	} else if temperature > 20 && isRaining {
		fmt.Println("🌧️ Warm but rainy - take an umbrella!")
	} else if temperature <= 20 && !isRaining {
		fmt.Println("🧥 Cool weather - wear a jacket!")
	} else {
		fmt.Println("🏠 Stay inside - cold and rainy!")
	}

	fmt.Println("\n🎯 NESTED IF STATEMENTS")
	fmt.Println("========================")

	// 🏗️ NESTED IF: If statements inside other if statements
	hasLicense := true
	hasInsurance := true
	carWorking := true

	if hasLicense {
		fmt.Println("✅ You have a license")
		if hasInsurance {
			fmt.Println("✅ You have insurance")
			if carWorking {
				fmt.Println("🚗 You can drive! Have a safe trip!")
			} else {
				fmt.Println("🔧 Car needs repair - can't drive today")
			}
		} else {
			fmt.Println("❌ No insurance - driving is illegal!")
		}
	} else {
		fmt.Println("❌ No license - you cannot drive!")
	}

	fmt.Println("\n🎯 BOOLEAN EXPRESSIONS")
	fmt.Println("======================")

	// 🔍 BOOLEAN VARIABLES: Can use boolean variables directly
	isLoggedIn := true
	isAdmin := false

	if isLoggedIn {  // 💡 NO NEED for == true
		fmt.Println("👤 User is logged in")
		
		if isAdmin {
			fmt.Println("👑 Admin privileges granted")
		} else {
			fmt.Println("👤 Regular user privileges")
		}
	} else {
		fmt.Println("🚫 Please log in first")
	}

	fmt.Println("\n🎯 COMPARISON OPERATORS")
	fmt.Println("=======================")

	// 🔢 COMPARISON OPERATORS: ==, !=, <, >, <=, >=
	x, y := 10, 20

	if x == y {
		fmt.Printf("📊 %d equals %d\n", x, y)
	} else if x != y {
		fmt.Printf("📊 %d does not equal %d\n", x, y)
	}

	if x < y {
		fmt.Printf("📊 %d is less than %d\n", x, y)
	}

	if x <= y {
		fmt.Printf("📊 %d is less than or equal to %d\n", x, y)
	}

	// 💡 NOTE: Go doesn't have ternary operator (condition ? true : false)
	// Use if-else instead!
	result := ""
	if x > y {
		result = "x is greater"
	} else {
		result = "y is greater or equal"
	}
	fmt.Printf("📊 Result: %s\n", result)
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

🔀 IF-ELSE SYNTAX PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Basic if-else                                                        │
│ if condition {                                                          │
│     // code                                                             │
│ } else {                                                                │
│     // code                                                             │
│ }                                                                       │
│                                                                         │
│ // If with initialization                                               │
│ if variable := value; condition {                                       │
│     // code (variable is accessible here)                              │
│ } else {                                                                │
│     // code (variable is accessible here too)                          │
│ }                                                                       │
│ // variable is NOT accessible here                                      │
│                                                                         │
│ // Multiple conditions                                                  │
│ if condition1 {                                                         │
│     // code                                                             │
│ } else if condition2 {                                                  │
│     // code                                                             │
│ } else {                                                                │
│     // code                                                             │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

🔗 LOGICAL OPERATORS:
┌─────────────────┬─────────────────┬─────────────────────────────────────┐
│   Operator      │     Symbol      │           Description               │
├─────────────────┼─────────────────┼─────────────────────────────────────┤
│ AND             │ &&              │ Both conditions must be true        │
│ OR              │ ||              │ At least one condition must be true │
│ NOT             │ !               │ Inverts the boolean value           │
└─────────────────┴─────────────────┴─────────────────────────────────────┘

📊 COMPARISON OPERATORS:
┌─────────────────┬─────────────────┬─────────────────────────────────────┐
│   Operator      │     Symbol      │           Description               │
├─────────────────┼─────────────────┼─────────────────────────────────────┤
│ Equal           │ ==              │ Values are equal                    │
│ Not Equal       │ !=              │ Values are not equal                │
│ Less Than       │ <               │ Left value is smaller               │
│ Greater Than    │ >               │ Left value is larger                │
│ Less or Equal   │ <=              │ Left value is smaller or equal      │
│ Greater or Equal│ >=              │ Left value is larger or equal       │
└─────────────────┴─────────────────┴─────────────────────────────────────┘

⚡ GO-SPECIFIC FEATURES:
• No parentheses required around conditions
• Braces are mandatory (even for single statements)
• Variable initialization in if statement
• No ternary operator (use if-else instead)

🚨 GOTCHAS:
❌ No ternary operator (condition ? true : false)
❌ Braces are mandatory (can't omit for single statements)
❌ Variables declared in if are block-scoped
❌ Assignment (=) vs comparison (==) confusion

💡 BOOLEAN EVALUATION:
• Go uses short-circuit evaluation
• In &&: if first is false, second isn't evaluated
• In ||: if first is true, second isn't evaluated
• Use this for safe null checking: obj != nil && obj.field

🔧 BEST PRACTICES:
• Use meaningful variable names in conditions
• Keep conditions simple and readable
• Use early returns to reduce nesting
• Group related conditions with parentheses for clarity
• Prefer positive conditions over negative when possible

🎯 WHEN TO USE IF-ELSE:
✅ Simple boolean decisions
✅ Input validation
✅ Error handling
✅ Business logic implementation
✅ Flow control based on state

❌ When you have many discrete values (use switch)
❌ When logic becomes too complex (extract to functions)

=============================================================================
*/