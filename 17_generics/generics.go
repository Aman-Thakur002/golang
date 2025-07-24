/*
=============================================================================
                           🧬 GO GENERICS TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Generics allow you to write functions and types that work with multiple types
while maintaining type safety. Think of them as "type templates" or "blueprints".

🔑 KEY FEATURES:
• Type parameters in square brackets [T any]
• Type constraints to limit allowed types
• Compile-time type safety (no runtime type assertions)
• Code reuse without sacrificing performance

💡 REAL-WORLD ANALOGY:
Generic = Universal Tool
- Regular function = Screwdriver (works only with screws)
- Generic function = Multi-tool (works with screws, bolts, nuts)
- Type parameter = Adjustable setting for different sizes
- Constraints = "Only works with metal fasteners" (not wood, plastic)

🎯 WHY USE GENERICS?
• Write code once, use with many types (DRY principle)
• Type safety at compile time (catch errors early)
• Better performance than interface{} + type assertions
• More expressive and maintainable code

=============================================================================
*/

/*
 -> A generic function declares one or more type parameters in square brackets immediately after the function name. Each type parameter can have an optional constraint (an interface) that restricts which concrete types are allowed.

 Map is a generic function with two type parameters: T and U.

The constraint any means "no restrictions" (any type is allowed).

At call‑site you can use it for, say, []int → []string or []float64 → []float64.

 -> Benefits
DRY Code: write algorithms once instead of duplicating for each type.

Type Safety: catches mismatches at compile time, unlike interface{} + type assertions.

Expressiveness: constraints let you specify exactly what capabilities ("≤", "+", methods, etc.) you need.

-> When to Use
Collections–style utilities (Map, Filter, Reduce).

Data structures (linked lists, trees, caches) that work over many element types.

Algorithms (search, sort, set) that apply across types.

Anywhere you'd otherwise accept interface{} and cast internally.

Generics let you write clean, reusable, and type‑safe code without the boilerplate of duplicating logic for each type. Once you've internalized the syntax ([type parameters] + [constraints]), you can start refactoring existing code to share more logic generically!

*/

package main

import "fmt"

// 🚫 PROBLEM: Without generics, we need separate functions for each type
// func printSlice(items ...int) { // variadic function
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func printStringSlice(items []string) { // variadic function
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }   -> we repeated function 2 times just for different data-types, to solve this we use generic functions

//-------------- generic function ----------------

// 🎯 SOLUTION 1: Union type constraint (specific types only)
// func printSlice[T int | string](items []T) {  // -> now we are allowing integer and string 
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// 🎯 SOLUTION 2: Built-in constraint (more flexible)
//-> we use comparable too
func printSlice[T comparable](items []T) {  // 📝 T must be comparable (==, !=)
	for _, item := range items {
		fmt.Println(item)  // Works with any comparable type
	}
}

// 🔧 GENERIC FUNCTION: Find element in slice
func findElement[T comparable](slice []T, target T) int {
	for i, item := range slice {
		if item == target {  // 💡 Requires comparable constraint
			return i
		}
	}
	return -1  // Not found
}

// 🎯 GENERIC FUNCTION: Map transformation
func mapSlice[T any, U any](slice []T, transform func(T) U) []U {
	result := make([]U, len(slice))
	for i, item := range slice {
		result[i] = transform(item)  // Apply transformation
	}
	return result
}

//-----------generic struct --------------

// 🏗️ GENERIC STRUCT: Stack that works with any type
type stack[T any] struct {  // T can be any type
	elements []T  // Slice of type T
}

// 🔧 METHODS ON GENERIC STRUCT
func (s *stack[T]) push(item T) {
	s.elements = append(s.elements, item)
}

func (s *stack[T]) pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T  // Zero value of type T
		return zero, false
	}
	
	index := len(s.elements) - 1
	item := s.elements[index]
	s.elements = s.elements[:index]
	return item, true
}

func (s *stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}

// 🎯 GENERIC INTERFACE: Constraint for numeric types
type Numeric interface {
	int | int8 | int16 | int32 | int64 |
	uint | uint8 | uint16 | uint32 | uint64 |
	float32 | float64
}

// 🔢 GENERIC FUNCTION: Sum with numeric constraint
func sum[T Numeric](numbers []T) T {
	var total T  // Zero value of numeric type
	for _, num := range numbers {
		total += num  // Addition works because of Numeric constraint
	}
	return total
}

func main() {
	fmt.Println("🧬 GENERICS LEARNING JOURNEY")
	fmt.Println("============================")

	fmt.Println("\n🎯 GENERIC STRUCT USAGE")
	fmt.Println("=======================")
    
	// 🏗️ CREATE GENERIC STACK: Specify type in square brackets
	myStack := stack[string] {
		elements: []string{"Java"},
	}
     fmt.Println("Initial stack:", myStack)

	// 🔧 USE STACK METHODS
	myStack.push("Go")
	myStack.push("Python")
	fmt.Println("After pushes:", myStack.elements)

	item, ok := myStack.pop()
	if ok {
		fmt.Println("Popped:", item)
		fmt.Println("Stack now:", myStack.elements)
	}

	fmt.Println("\n🎯 GENERIC FUNCTION USAGE")
	fmt.Println("=========================")

	// 📋 PRINT DIFFERENT TYPES: Same function, different types
	nums := []int{1, 3, 9}
	stringsSlice := []string{"Golang", "Javascript"}
	
	fmt.Println("Numbers:")
	printSlice(nums)        // T inferred as int
	fmt.Println("Strings:")
	printSlice(stringsSlice) // T inferred as string

	fmt.Println("\n🎯 FIND ELEMENT EXAMPLE")
	fmt.Println("=======================")

	// 🔍 FIND IN DIFFERENT TYPES
	numIndex := findElement(nums, 3)
	strIndex := findElement(stringsSlice, "Go")
	
	fmt.Printf("Number 3 found at index: %d\n", numIndex)
	fmt.Printf("String 'Go' found at index: %d\n", strIndex)

	fmt.Println("\n🎯 MAP TRANSFORMATION")
	fmt.Println("=====================")

	// 🔄 TRANSFORM: int to string
	intToString := mapSlice(nums, func(n int) string {
		return fmt.Sprintf("num_%d", n)
	})
	fmt.Println("Transformed to strings:", intToString)

	// 🔄 TRANSFORM: string to length
	stringToLength := mapSlice(stringsSlice, func(s string) int {
		return len(s)
	})
	fmt.Println("String lengths:", stringToLength)

	fmt.Println("\n🎯 NUMERIC CONSTRAINT")
	fmt.Println("=====================")

	// 🔢 SUM DIFFERENT NUMERIC TYPES
	intSum := sum([]int{1, 2, 3, 4, 5})
	floatSum := sum([]float64{1.1, 2.2, 3.3})
	
	fmt.Printf("Sum of ints: %d\n", intSum)
	fmt.Printf("Sum of floats: %.1f\n", floatSum)

	fmt.Println("\n🎯 INTEGER STACK EXAMPLE")
	fmt.Println("========================")

	// 🔢 INTEGER STACK
	intStack := stack[int]{elements: []int{}}
	intStack.push(10)
	intStack.push(20)
	intStack.push(30)
	
	fmt.Println("Integer stack:", intStack.elements)
	
	for !intStack.isEmpty() {
		if val, ok := intStack.pop(); ok {
			fmt.Printf("Popped: %d\n", val)
		}
	}
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

🧬 GENERIC SYNTAX:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Generic function                                                     │
│ func name[T constraint](param T) T { }                                  │
│                                                                         │
│ // Multiple type parameters                                             │
│ func name[T, U any](a T, b U) (T, U) { }                               │
│                                                                         │
│ // Generic struct                                                       │
│ type MyStruct[T any] struct {                                           │
│     field T                                                             │
│ }                                                                       │
│                                                                         │
│ // Generic interface (constraint)                                       │
│ type MyConstraint interface {                                           │
│     int | string | float64                                              │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

🎯 BUILT-IN CONSTRAINTS:
┌─────────────────┬─────────────────────────────────────────────────────────┐
│   Constraint    │                Description                              │
├─────────────────┼─────────────────────────────────────────────────────────┤
│ any             │ No restrictions (same as interface{})                   │
│ comparable      │ Types that support == and != operators                 │
│ Ordered         │ Types that support <, <=, >, >= (golang.org/x/exp)     │
└─────────────────┴─────────────────────────────────────────────────────────┘

🔍 CONSTRAINT TYPES:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Union constraint (specific types)                                    │
│ func process[T int | string | float64](value T) { }                     │
│                                                                         │
│ // Interface constraint (method requirements)                           │
│ type Stringer interface {                                               │
│     String() string                                                     │
│ }                                                                       │
│ func print[T Stringer](value T) { }                                     │
│                                                                         │
│ // Underlying type constraint                                           │
│ type MyInt int                                                          │
│ func process[T ~int](value T) { }  // Accepts int and MyInt             │
└─────────────────────────────────────────────────────────────────────────┘

⚡ TYPE INFERENCE:
• Go can often infer type parameters from arguments
• Explicit specification: func[int](value)
• Inferred: func(value) where value is int

🎯 COMMON PATTERNS:
• Container types: Stack[T], Queue[T], Map[K,V]
• Utility functions: Map, Filter, Reduce
• Algorithms: Sort, Search, Min, Max
• Optional/Result types: Option[T], Result[T,E]

🚨 GOTCHAS:
❌ Can't use type parameters in const declarations
❌ Method type parameters not allowed (only function/struct)
❌ Type inference doesn't work in all cases
❌ Constraints must be interfaces

💡 PERFORMANCE:
• Generics are compile-time feature (no runtime overhead)
• Each instantiation creates specialized code
• Better performance than interface{} + type assertions
• May increase binary size with many instantiations

🔧 BEST PRACTICES:
• Use meaningful constraint names (not just T, U)
• Prefer built-in constraints when possible
• Keep type parameter lists short
• Use generics for algorithms, not just type storage
• Document constraints clearly

🎯 WHEN TO USE GENERICS:
✅ Data structures (containers, collections)
✅ Algorithms that work across types
✅ Utility functions (map, filter, reduce)
✅ Type-safe APIs without interface{}

❌ Simple functions with single type
❌ When interface{} is actually needed
❌ Over-engineering simple problems

=============================================================================
*/