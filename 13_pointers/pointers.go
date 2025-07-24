/*
=============================================================================
                           👉 GO POINTERS TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Pointers store memory addresses instead of values directly.
Think of them as "directions to a house" rather than the house itself.

🔑 KEY FEATURES:
• & operator gets address of a variable (address-of)
• * operator accesses value at address (dereference)
• Enable pass-by-reference behavior
• Allow functions to modify original variables

💡 REAL-WORLD ANALOGY:
Pointer = House Address
- Variable = Actual house
- Pointer = Address written on paper (123 Main St)
- & operator = "What's the address of this house?"
- * operator = "Go to this address and see what's there"

🎯 WHY USE POINTERS?
• Modify variables in functions (pass by reference)
• Avoid copying large data structures
• Share data between different parts of program
• Enable efficient memory usage

=============================================================================
*/

package main

import "fmt"

// 📋 PASS BY VALUE: Function receives a copy of the variable
// pass by value , so num passed in this function is just a copy 
func changeNum(num int) {
	num = 5  // 📝 Changes only the local copy, not original
	fmt.Println("In change num", num)  // Shows 5
}

// 👉 PASS BY REFERENCE: Function receives pointer to original variable
//pass by refernece 
func fun1(num *int){  // *int means "pointer to int"
   *num = 5  // 🎯 DEREFERENCE: Access and modify value at the address
   // *num means "the value that num points to"
}

// 🔧 PRACTICAL EXAMPLE: Swap two variables
func swap(a, b *int) {
	temp := *a  // Get value at address a
	*a = *b     // Set value at address a to value at address b
	*b = temp   // Set value at address b to temp
}

// 📊 WORKING WITH STRUCT POINTERS
type Person struct {
	name string
	age  int
}

func updateAge(p *Person, newAge int) {
	p.age = newAge  // 💡 Go automatically dereferences struct pointers
	// No need to write (*p).age - Go does it for you!
}

func main() {
	fmt.Println("👉 POINTERS LEARNING JOURNEY")
	fmt.Println("============================")

	fmt.Println("\n🎯 PASS BY VALUE vs PASS BY REFERENCE")
	fmt.Println("======================================")

	num := 1
	fmt.Println("Original value:", num)  // 1

	// 📋 PASS BY VALUE: Original unchanged
	changeNum(num)  // Passes copy of num
	fmt.Println("After change num:", num)  // Still 1 - original unchanged

	fmt.Println("\n🎯 WORKING WITH ADDRESSES")
	fmt.Println("=========================")

	// 📍 ADDRESS OPERATOR (&): Get memory address
	fmt.Println("Memory address:", &num)  // Shows memory address like 0xc000014098
	
	// 👉 PASS BY REFERENCE: Original gets modified
	fun1(&num)  // Pass address of num (not the value)
	fmt.Println("After change num (pass by reference):", num)  // Now 5!

	fmt.Println("\n🎯 POINTER VARIABLES")
	fmt.Println("====================")

	// 📝 DECLARE POINTER VARIABLE
	var ptr *int        // ptr is a pointer to int (initially nil)
	fmt.Println("Nil pointer:", ptr)  // <nil>

	// 🎯 ASSIGN ADDRESS TO POINTER
	x := 42
	ptr = &x           // ptr now points to x
	fmt.Println("Pointer value (address):", ptr)    // Memory address
	fmt.Println("Dereferenced value:", *ptr)        // 42 (value at address)

	// 🔄 MODIFY THROUGH POINTER
	*ptr = 100         // Change value at address ptr points to
	fmt.Println("x after pointer modification:", x)  // 100

	fmt.Println("\n🎯 PRACTICAL EXAMPLE: SWAP")
	fmt.Println("===========================")

	a, b := 10, 20
	fmt.Printf("Before swap: a=%d, b=%d\n", a, b)
	
	swap(&a, &b)  // Pass addresses of a and b
	fmt.Printf("After swap: a=%d, b=%d\n", a, b)

	fmt.Println("\n🎯 STRUCT POINTERS")
	fmt.Println("==================")

	// 👤 STRUCT WITH POINTER
	person := Person{name: "Alice", age: 25}
	fmt.Printf("Before: %+v\n", person)

	updateAge(&person, 30)  // Pass pointer to struct
	fmt.Printf("After: %+v\n", person)

	fmt.Println("\n🎯 POINTER ARITHMETIC (NOT ALLOWED)")
	fmt.Println("===================================")

	// 🚫 Go doesn't allow pointer arithmetic like C/C++
	// ptr++     // ❌ This would cause compile error
	// ptr + 1   // ❌ This would cause compile error
	fmt.Println("✅ Go pointers are safe - no arithmetic allowed!")
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

👉 POINTER SYNTAX:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Declaration                                                          │
│ var ptr *int           // ptr is pointer to int                        │
│                                                                         │
│ // Get address                                                          │
│ ptr = &variable        // & gets address of variable                    │
│                                                                         │
│ // Dereference                                                          │
│ value := *ptr          // * gets value at address                      │
│ *ptr = newValue        // * sets value at address                      │
└─────────────────────────────────────────────────────────────────────────┘

🔍 POINTER OPERATORS:
┌─────────────────┬─────────────────┬─────────────────────────────────────┐
│   Operator      │      Name       │           Description               │
├─────────────────┼─────────────────┼─────────────────────────────────────┤
│ &variable       │ Address-of      │ Gets memory address of variable     │
│ *pointer        │ Dereference     │ Gets/sets value at pointer address  │
│ *Type           │ Pointer type    │ Declares pointer to Type            │
└─────────────────┴─────────────────┴─────────────────────────────────────┘

📊 PASS BY VALUE vs PASS BY REFERENCE:
┌─────────────────┬─────────────────┬─────────────────────────────────────┐
│    Aspect       │  Pass by Value  │        Pass by Reference            │
├─────────────────┼─────────────────┼─────────────────────────────────────┤
│ Parameter       │ func(x int)     │ func(x *int)                        │
│ Function call   │ f(variable)     │ f(&variable)                        │
│ Inside function │ x = newValue    │ *x = newValue                       │
│ Original change │ No              │ Yes                                 │
│ Memory usage    │ Copy created    │ Only address passed                 │
└─────────────────┴─────────────────┴─────────────────────────────────────┘

⚡ POINTER CHARACTERISTICS:
• Zero value of pointer is nil
• Dereferencing nil pointer causes panic
• Go automatically dereferences struct pointers
• No pointer arithmetic (unlike C/C++)
• Garbage collector handles memory management

🎯 COMMON PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Check for nil before dereferencing                                   │
│ if ptr != nil {                                                         │
│     value := *ptr                                                       │
│ }                                                                       │
│                                                                         │
│ // Function that may modify parameter                                   │
│ func modify(data *SomeStruct) {                                         │
│     if data != nil {                                                    │
│         data.field = newValue                                           │
│     }                                                                   │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

🚨 GOTCHAS:
❌ Dereferencing nil pointer = panic
❌ Forgetting & when passing to pointer parameter
❌ Forgetting * when accessing pointer value
❌ Pointer arithmetic not allowed (good thing!)

💡 MEMORY SAFETY:
• Go pointers are "safe" - no arithmetic
• Garbage collector prevents memory leaks
• Cannot create invalid pointers
• Runtime checks prevent most pointer errors

🔧 BEST PRACTICES:
• Use pointers for large structs to avoid copying
• Always check for nil before dereferencing
• Use pointers when function needs to modify parameter
• Prefer value types when possible (simpler)
• Use pointer receivers for methods that modify

🎯 WHEN TO USE POINTERS:
✅ Function needs to modify parameter
✅ Avoiding expensive copies of large data
✅ Sharing data between functions
✅ Optional values (nil = not present)
✅ Implementing linked data structures

❌ For simple values (int, bool, small structs)
❌ When value semantics are clearer
❌ When you don't need to modify original

=============================================================================
*/