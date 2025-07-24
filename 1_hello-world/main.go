/*
=============================================================================
                           👋 GO HELLO WORLD TUTORIAL
=============================================================================

📚 CORE CONCEPT:
This is your first Go program! Every Go journey starts here.
Understanding this simple program teaches you the fundamental structure of Go.

🔑 KEY COMPONENTS:
• Package declaration (every Go file needs one)
• Import statements (bring in external functionality)
• Main function (entry point of the program)
• Function calls (executing code)

💡 REAL-WORLD ANALOGY:
Go Program = Recipe
- Package = Recipe category (desserts, main dishes)
- Import = Ingredients you need from the store
- Main function = "Start cooking here"
- Function calls = Individual cooking steps

🎯 WHY START WITH HELLO WORLD?
• Verify your Go installation works
• Understand basic program structure
• Learn the minimal code needed to run
• Foundation for all future programs

=============================================================================
*/

package main  // 📦 PACKAGE DECLARATION: Every Go file belongs to a package
              // "main" is special - it creates an executable program

import "fmt" // 📥 IMPORT: Brings in the "fmt" package for formatting
             // fmt = format, used for input/output operations
             // This is from Go's standard library

// 🚀 MAIN FUNCTION: The entry point of every Go program
// When you run the program, execution starts here
func main() {  // func = function keyword, main = function name
   // 🎯 PRINTLN: Print line - outputs text and adds a newline
   fmt.Println("Hello, World!") // Print "Hello, World!" to the console
   
   // 💡 Let's explore more fmt functions
   fmt.Print("This prints without newline. ")
   fmt.Print("See? Same line!\n")  // \n = manual newline
   
   // 🎨 FORMATTED PRINTING: Printf allows formatting
   name := "Go Developer"
   age := 25
   fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
   // %s = string placeholder, %d = integer placeholder
   
   // 🌟 MULTIPLE WAYS TO SAY HELLO
   fmt.Println("🌍 Hello, World!")
   fmt.Println("🇺🇸 Hello, World!")
   fmt.Println("🇪🇸 ¡Hola, Mundo!")
   fmt.Println("🇫🇷 Bonjour, le Monde!")
   fmt.Println("🇩🇪 Hallo, Welt!")
   fmt.Println("🇯🇵 こんにちは、世界！")
   
   // 🎉 CONGRATULATIONS MESSAGE
   fmt.Println("\n🎉 Congratulations!")
   fmt.Println("✅ You've successfully run your first Go program!")
   fmt.Println("🚀 Ready to explore more Go features!")
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

🏗️ GO PROGRAM STRUCTURE:
┌─────────────────────────────────────────────────────────────────────────┐
│ package main           // 1. Package declaration (required)             │
│                                                                         │
│ import "fmt"           // 2. Import statements (optional)               │
│                                                                         │
│ func main() {          // 3. Main function (required for executables)  │
│     // Your code here  // 4. Program logic                             │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

📦 PACKAGE SYSTEM:
• Every Go file must belong to a package
• "main" package creates executable programs
• Other package names create libraries
• Package name should match directory name (usually)

📥 IMPORT STATEMENT:
• Brings external packages into your program
• "fmt" is from Go's standard library
• Can import multiple packages: import ("fmt", "time")
• Unused imports cause compilation errors

🚀 MAIN FUNCTION:
• Entry point for executable programs
• Must be in "main" package
• Takes no parameters, returns nothing
• Program execution starts and ends here

🎨 FMT PACKAGE FUNCTIONS:
┌─────────────────┬─────────────────────────────────────────────────────────┐
│   Function      │                Description                              │
├─────────────────┼─────────────────────────────────────────────────────────┤
│ fmt.Print()     │ Prints without newline                                  │
│ fmt.Println()   │ Prints with newline                                     │
│ fmt.Printf()    │ Prints with formatting (like C's printf)               │
│ fmt.Sprintf()   │ Returns formatted string (doesn't print)               │
└─────────────────┴─────────────────────────────────────────────────────────┘

🎯 PRINTF FORMAT SPECIFIERS:
• %s = string
• %d = decimal integer
• %f = floating point
• %t = boolean
• %v = any value (default format)
• %T = type of value

⚡ COMPILATION & EXECUTION:
1. Save file as main.go (or any name ending in .go)
2. Run: go run main.go (compiles and runs immediately)
3. Or: go build main.go (creates executable file)
4. Then: ./main (runs the executable)

🚨 COMMON MISTAKES:
❌ Forgetting package declaration
❌ Wrong package name (must be "main" for executables)
❌ Unused imports (Go won't compile)
❌ Missing main function in main package
❌ Typos in function names (Go is case-sensitive)

💡 GO PHILOSOPHY:
• Simplicity: Minimal syntax, clear structure
• Readability: Code should be easy to understand
• Efficiency: Fast compilation and execution
• Reliability: Strong typing and error handling

🔧 BEST PRACTICES:
• Use meaningful package names
• Import only what you need
• Keep main function simple (delegate to other functions)
• Use gofmt to format your code consistently
• Add comments to explain complex logic

🎯 NEXT STEPS:
After mastering Hello World, you're ready to learn:
• Variables and data types
• Control structures (if, for, switch)
• Functions and methods
• Data structures (arrays, slices, maps)
• Concurrency with goroutines

=============================================================================
*/