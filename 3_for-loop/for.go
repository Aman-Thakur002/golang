/*
=============================================================================
                           🔄 GO FOR LOOPS TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Go has only ONE loop construct: the for loop!
But it's flexible enough to handle all looping scenarios you need.

🔑 KEY FEATURES:
• Only loop keyword in Go (no while, do-while)
• Three components: init; condition; post
• Can omit any component for different behaviors
• Range-based iteration for collections

💡 REAL-WORLD ANALOGY:
For Loop = Assembly Line Process
- Init = Set up the workspace
- Condition = "Keep working while materials available"
- Post = Move to next item
- Body = Do the actual work

🎯 WHY MASTER FOR LOOPS?
• Essential for processing collections
• Automate repetitive tasks
• Control program flow precisely
• Foundation for algorithms

=============================================================================
*/

package main

import "fmt"

func main() {
	fmt.Println("🔄 FOR LOOPS LEARNING JOURNEY")
	fmt.Println("==============================")

	fmt.Println("\n🎯 WHILE-STYLE LOOP")
	fmt.Println("====================")

	// 🔄 WHILE LOOP FASHION: Only condition, no init or post
	k := 1
	for k <= 3 {  // 💡 WHILE-STYLE: Just the condition
		fmt.Printf("🔢 While-style: %d\n", k)
		k = k + 1  // Manual increment
	}

	fmt.Println("\n🎯 TRADITIONAL FOR LOOP")
	fmt.Println("=======================")

	// 🎯 TRADITIONAL FOR LOOP: init; condition; post
	for i := 0; i < 3; i++ {  // 💡 CLASSIC: init; condition; post
		fmt.Printf("🔢 Traditional: %d\n", i)
	}

	fmt.Println("\n🎯 RANGE-BASED LOOP (Modern)")
	fmt.Println("=============================")

	// 🚀 MODERN WAY: Range-based looping (Go 1.22+)
	for y := range 3 {  // 💡 RANGE: Iterates from 0 to 2
		fmt.Printf("🔢 Range-based: %d\n", y)
	}

	fmt.Println("\n🎯 INFINITE LOOP")
	fmt.Println("=================")

	// ♾️ INFINITE LOOP: No condition means loop forever
	counter := 0
	for {  // 💡 INFINITE: No condition = true forever
		fmt.Printf("♾️ Infinite loop iteration: %d\n", counter)
		counter++
		if counter >= 3 {
			break  // 🛑 BREAK: Exit the loop
		}
	}

	fmt.Println("\n🎯 LOOP WITH CONTINUE")
	fmt.Println("======================")

	// ⏭️ CONTINUE: Skip current iteration
	fmt.Println("🔢 Even numbers only:")
	for i := 0; i < 10; i++ {
		if i%2 != 0 {  // If odd number
			continue  // 💡 CONTINUE: Skip to next iteration
		}
		fmt.Printf("   Even: %d\n", i)
	}

	fmt.Println("\n🎯 NESTED LOOPS")
	fmt.Println("================")

	// 🏗️ NESTED LOOPS: Loop inside another loop
	fmt.Println("📊 Multiplication table (3x3):")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d×%d=%d  ", i, j, i*j)
		}
		fmt.Println()  // New line after each row
	}

	fmt.Println("\n🎯 LOOP WITH SLICE")
	fmt.Println("===================")

	// 📋 ITERATING OVER SLICE: Traditional way
	numbers := []int{10, 20, 30, 40, 50}
	
	fmt.Println("Traditional slice iteration:")
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("   Index %d: %d\n", i, numbers[i])
	}

	fmt.Println("Range-based slice iteration:")
	for index, value := range numbers {  // 💡 RANGE: index, value
		fmt.Printf("   Index %d: %d\n", index, value)
	}

	fmt.Println("\n🎯 LOOP CONTROL STATEMENTS")
	fmt.Println("===========================")

	// 🎮 LOOP CONTROL: break and continue in action
	fmt.Println("🔍 Finding first number divisible by 7:")
	for i := 1; i <= 50; i++ {
		if i%7 == 0 {
			fmt.Printf("   Found: %d\n", i)
			break  // 🛑 EXIT: Stop the loop immediately
		}
		if i%10 == 0 {
			fmt.Printf("   Checkpoint: %d\n", i)
		}
	}

	fmt.Println("\n🎯 LABELED BREAK (Advanced)")
	fmt.Println("============================")

	// 🏷️ LABELED BREAK: Break out of nested loops
	fmt.Println("🔍 Finding pair that sums to 15:")
	outer:  // 💡 LABEL: Name for the outer loop
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			if i+j == 15 {
				fmt.Printf("   Found pair: %d + %d = 15\n", i, j)
				break outer  // 🛑 LABELED BREAK: Exit outer loop
			}
		}
	}

	fmt.Println("\n🎯 LOOP PATTERNS")
	fmt.Println("=================")

	// 🎨 COMMON PATTERNS: Different ways to use loops
	
	// Pattern 1: Countdown
	fmt.Println("🚀 Countdown:")
	for i := 5; i > 0; i-- {
		fmt.Printf("   %d...\n", i)
	}
	fmt.Println("   🎉 Blast off!")

	// Pattern 2: Step by 2
	fmt.Println("🔢 Counting by 2s:")
	for i := 0; i <= 10; i += 2 {
		fmt.Printf("   %d ", i)
	}
	fmt.Println()

	// Pattern 3: Reverse iteration
	fmt.Println("🔄 Reverse slice iteration:")
	for i := len(numbers) - 1; i >= 0; i-- {
		fmt.Printf("   %d ", numbers[i])
	}
	fmt.Println()
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

🔄 FOR LOOP VARIATIONS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Traditional for loop                                                 │
│ for init; condition; post {                                             │
│     // body                                                             │
│ }                                                                       │
│                                                                         │
│ // While-style loop                                                     │
│ for condition {                                                         │
│     // body                                                             │
│ }                                                                       │
│                                                                         │
│ // Infinite loop                                                        │
│ for {                                                                   │
│     // body (use break to exit)                                        │
│ }                                                                       │
│                                                                         │
│ // Range-based loop                                                     │
│ for index, value := range collection {                                  │
│     // body                                                             │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

🎮 LOOP CONTROL STATEMENTS:
┌─────────────────┬─────────────────────────────────────────────────────────┐
│   Statement     │                Description                              │
├─────────────────┼─────────────────────────────────────────────────────────┤
│ break           │ Exit the current loop immediately                       │
│ continue        │ Skip rest of current iteration, go to next              │
│ break label     │ Exit the labeled loop (useful for nested loops)        │
│ continue label  │ Continue the labeled loop                               │
└─────────────────┴─────────────────────────────────────────────────────────┘

⚡ LOOP COMPONENTS:
• Init: Executed once before loop starts (e.g., i := 0)
• Condition: Checked before each iteration (e.g., i < 10)
• Post: Executed after each iteration (e.g., i++)
• Body: Code that runs in each iteration

🔍 RANGE LOOP DETAILS:
• Arrays/Slices: for i, v := range arr (index, value)
• Maps: for k, v := range m (key, value)
• Strings: for i, r := range s (byte index, rune)
• Channels: for v := range ch (value only)
• Numbers: for i := range n (0 to n-1, Go 1.22+)

🚨 GOTCHAS:
❌ Range variable is reused (be careful with goroutines)
❌ Modifying slice during range can cause issues
❌ String range iterates by runes, not bytes
❌ Map iteration order is random

💡 PERFORMANCE TIPS:
• Range loops are generally as fast as traditional loops
• Pre-calculate slice length if used multiple times
• Use break/continue to avoid unnecessary iterations
• Consider loop unrolling for performance-critical code

🔧 BEST PRACTICES:
• Use range loops for collections when possible
• Use meaningful variable names (not just i, j)
• Keep loop bodies simple (extract complex logic to functions)
• Use break/continue for early exit conditions
• Prefer range over traditional loops for readability

🎯 WHEN TO USE EACH TYPE:
• Traditional for: When you need precise control over iteration
• While-style: When you don't know iteration count in advance
• Range: When iterating over collections
• Infinite: For servers, event loops, or until-condition patterns

=============================================================================
*/