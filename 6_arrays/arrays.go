/*
=============================================================================
                           📊 GO ARRAYS TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Arrays are fixed-size sequences of elements of the same type.
Think of them as "numbered boxes" where each box holds one item.

🔑 KEY FEATURES:
• Fixed size (determined at compile time)
• All elements must be same type
• Zero-indexed (first element is at index 0)
• Value type (copied when assigned/passed)

💡 REAL-WORLD ANALOGY:
Array = Apartment building with numbered units
- Size = Number of apartments (fixed when built)
- Index = Apartment number (0, 1, 2, 3...)
- Element = Person living in each apartment
- Type = All apartments house the same type (e.g., all 1-bedroom)

🎯 WHY USE ARRAYS?
• When you know exact size needed
• Memory-efficient (no overhead)
• Fast access by index
• Foundation for slices

⚠️ NOTE: In Go, slices are more commonly used than arrays!

=============================================================================
*/

package main

import "fmt"

func main() {
	fmt.Println("📊 ARRAYS LEARNING JOURNEY")
	fmt.Println("===========================")

	// 🏗️ DECLARE ARRAY: var name [size]type
	var nums [4]int  // Array of 4 integers, initialized to zero values
	
	// 📝 SET INDIVIDUAL ELEMENTS: array[index] = value
	nums[0] = 1  // Set first element
	
	// 🔍 ACCESS ELEMENTS: array[index]
	fmt.Println("📋 First element:", nums[0])
	fmt.Println("📋 Full array:", nums)        // [1 0 0 0] - unset elements are zero
	fmt.Println("📏 Array length:", len(nums)) // len() gives array size

	fmt.Println("\n🎯 ZERO VALUES DEMONSTRATION")
	fmt.Println("=============================")

	// 🔢 INTEGER ARRAY: Zero value is 0
	var boolArray [4]bool
	fmt.Println("📋 Bool array (zero values):", boolArray)  // [false false false false]

	// 📝 STRING ARRAY: Zero value is empty string ""
	var names [3]string
	fmt.Println("📋 String array (zero values):", names)  // ["" "" ""]

	fmt.Println("\n🎯 ARRAY INITIALIZATION")
	fmt.Println("========================")
 
    // 🎯 ARRAY LITERAL: Initialize with values
	numsArray := [3]int{1, 2, 4}  // [size]type{values}
	fmt.Println("📋 Initialized array:", numsArray)

	// 💡 AUTO-SIZE: Let Go count the elements
	autoArray := [...]int{10, 20, 30, 40}  // [...] means "count for me"
	fmt.Printf("📋 Auto-sized array: %v (length: %d)\n", autoArray, len(autoArray))

	// 🎯 PARTIAL INITIALIZATION: Specify some indices
	sparseArray := [5]int{1: 100, 3: 300}  // index:value syntax
	fmt.Println("📋 Sparse array:", sparseArray)  // [0 100 0 300 0]

	fmt.Println("\n🏢 2D ARRAYS")
	fmt.Println("=============")

	// 🏢 2D ARRAY: Array of arrays
	num2dArray := [2][2]int{{1, 2}, {3, 4}}  // 2x2 matrix
	fmt.Println("📋 2D array:", num2dArray)
	fmt.Println("📋 Element [0][1]:", num2dArray[0][1])  // Access element at row 0, col 1

	// 🎯 LARGER 2D ARRAY EXAMPLE
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("📋 3x3 Matrix:")
	for i := 0; i < len(matrix); i++ {
		fmt.Printf("   Row %d: %v\n", i, matrix[i])
	}

	fmt.Println("\n🔄 ARRAY ITERATION")
	fmt.Println("===================")

	// 🔄 ITERATE WITH TRADITIONAL FOR LOOP
	fmt.Println("Traditional loop:")
	for i := 0; i < len(numsArray); i++ {
		fmt.Printf("   Index %d: %d\n", i, numsArray[i])
	}

	// 🔄 ITERATE WITH RANGE (more idiomatic)
	fmt.Println("Range loop:")
	for index, value := range numsArray {
		fmt.Printf("   Index %d: %d\n", index, value)
	}
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

📊 ARRAY DECLARATION PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Basic declaration                                                    │
│ var arr [5]int              // Zero-initialized array                   │
│                                                                         │
│ // Array literal                                                        │
│ arr := [3]int{1, 2, 3}      // Explicit size                           │
│ arr := [...]int{1, 2, 3}    // Auto-size (compiler counts)             │
│                                                                         │
│ // Partial initialization                                               │
│ arr := [5]int{1: 10, 3: 30} // Only indices 1 and 3 set               │
│                                                                         │
│ // 2D arrays                                                            │
│ arr := [2][3]int{{1,2,3}, {4,5,6}} // 2 rows, 3 columns               │
└─────────────────────────────────────────────────────────────────────────┘

🔍 ARRAY vs SLICE COMPARISON:
┌─────────────────┬─────────────────┬─────────────────────────────────────┐
│    Aspect       │     Array       │              Slice                  │
├─────────────────┼─────────────────┼─────────────────────────────────────┤
│ Size            │ Fixed           │ Dynamic                             │
│ Declaration     │ [5]int          │ []int                               │
│ Memory          │ Value type      │ Reference type                      │
│ Passing         │ Copies data     │ Copies reference                    │
│ Performance     │ Faster          │ Slightly slower (indirection)       │
│ Usage           │ Rare            │ Very common                         │
│ Zero Value      │ Array of zeros  │ nil                                 │
└─────────────────┴─────────────────┴─────────────────────────────────────┘

⚡ ARRAY CHARACTERISTICS:
• Size is part of the type: [3]int and [4]int are different types
• Arrays are value types (assignment copies all elements)
• Passing to functions copies the entire array
• Cannot change size after declaration
• All elements initialized to zero value of the type

🎯 ZERO VALUES BY TYPE:
• int, float: 0
• bool: false
• string: ""
• pointer: nil
• struct: struct with all fields set to their zero values

🚨 GOTCHAS:
❌ Arrays of different sizes are different types
❌ Passing large arrays to functions is expensive (copies all data)
❌ Cannot append to arrays (use slices instead)
❌ Array size must be known at compile time

💡 MEMORY LAYOUT:
Arrays store elements in contiguous memory locations:
[1][2][3][4] ← All elements stored next to each other

🔧 BEST PRACTICES:
• Use slices instead of arrays in most cases
• Arrays are good for fixed-size data (e.g., RGB values, coordinates)
• Use arrays when you need value semantics
• Consider arrays for small, fixed collections
• Use [...] for auto-sizing when possible

🎯 WHEN TO USE ARRAYS:
✅ Fixed-size data (coordinates, RGB values)
✅ When you need value semantics
✅ Performance-critical code with known size
✅ Interfacing with C code

❌ When size varies at runtime
❌ When you need to append/remove elements
❌ Most general-purpose collections (use slices)

=============================================================================
*/