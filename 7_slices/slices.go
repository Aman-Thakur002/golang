/*
=============================================================================
                           🍰 GO SLICES TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Slices are Go's dynamic arrays - they can grow and shrink during runtime.
Think of them as "flexible arrays" that automatically resize when needed.

🔑 KEY DIFFERENCES FROM ARRAYS:
• Arrays: Fixed size, value type
• Slices: Dynamic size, reference type (points to underlying array)

💡 REAL-WORLD ANALOGY:
Array = Fixed parking lot (10 spaces, can't change)
Slice = Expandable parking lot (starts small, adds spaces when needed)

🎯 SLICE INTERNALS:
Slice has 3 components:
• Pointer: Points to underlying array
• Length: Current number of elements
• Capacity: Maximum elements before reallocation

=============================================================================
*/

package main

import (
	"fmt"
	"slices"
)

// no need to tell the size
// -dynamic array
// - most used construct in go

func main(){
// 🚨 NIL SLICE: Uninitialized slice is nil	
 var nums []int  // Declaration without initialization
 fmt.Println("Slice",nums)
 fmt.Println(nums == nil)  // true - nil slice

 // 🏗️ CREATING NON-NIL SLICE: Using make()
  nums1 := make([]int,0,5)  // make(type, initial length, capacity)
  nums1 = append(nums1, 1)  // append() adds elements to slice
  nums1 = append(nums1, 2)  // slice grows automatically
 fmt.Println("Not null slice 1",nums1)
 fmt.Println("Length : ",len(nums1))    // len() = current number of elements
 fmt.Println("capacity : ",cap(nums1))  // cap() = maximum before reallocation

 // 🎯 ALTERNATIVE: Empty slice literal (not nil)
 nums2 := []int{}  // Empty slice (length=0, but not nil)
 nums2 = append(nums2, 4)
 fmt.Println("Not null slice 2",nums2)

 // 📋 COPYING SLICES: copy() function
 copiedNums := make([]int, len(nums2))  // Create destination slice
 copy(copiedNums, nums2)                // copy(destination, source)
 fmt.Println("Copied Slice",copiedNums)

 // ✂️ SLICE OPERATOR: Extract portions of arrays/slices
 nums3 := [3]int{2,3,4}  // This is an array (fixed size)
 fmt.Println("Sliced array of nums3",nums3[0:2])  // [start:end) - end is exclusive
 fmt.Println("Sliced array of nums3",nums3[:3])   // [:end] - from beginning
 fmt.Println("Sliced array of nums3",nums3[0:])   // [start:] - to end

 // 📦 SLICES PACKAGE: Utility functions for slices
 s1 := []int{1,2}
 s2 := []int{1,2}
 s3 := []int{1,3}
 fmt.Println(slices.Equal(s1,s2))  // true - same elements
 fmt.Println(slices.Equal(s1,s3))  // false - different elements

 // 🏢 2D SLICES: Slice of slices
 array2D := [][]int{{1,2,4},{3,4}}  // Each inner slice can have different length
 fmt.Println(array2D)
 
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

🔍 SLICE vs ARRAY COMPARISON:
┌─────────────────┬─────────────────┬─────────────────────────────────────┐
│    Aspect       │     Array       │              Slice                  │
├─────────────────┼─────────────────┼─────────────────────────────────────┤
│ Size            │ Fixed           │ Dynamic                             │
│ Declaration     │ [5]int          │ []int                               │
│ Memory          │ Value type      │ Reference type                      │
│ Passing         │ Copies data     │ Copies reference                    │
│ Performance     │ Faster          │ Slightly slower (indirection)       │
│ Usage           │ Rare            │ Very common                         │
└─────────────────┴─────────────────┴─────────────────────────────────────┘

⚡ SLICE OPERATIONS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Create                                                               │
│ var s []int              // nil slice                                   │
│ s := []int{}             // empty slice                                 │
│ s := make([]int, 5)      // length 5, capacity 5                       │
│ s := make([]int, 0, 10)  // length 0, capacity 10                      │
│                                                                         │
│ // Add elements                                                         │
│ s = append(s, 1)         // add single element                          │
│ s = append(s, 2, 3, 4)   // add multiple elements                      │
│ s = append(s, other...)  // add another slice                          │
│                                                                         │
│ // Access                                                               │
│ fmt.Println(s[0])        // first element                              │
│ fmt.Println(s[len(s)-1]) // last element                               │
│                                                                         │
│ // Slice operations                                                     │
│ sub := s[1:3]            // elements 1 and 2                           │
│ sub := s[:3]             // first 3 elements                           │
│ sub := s[2:]             // from index 2 to end                        │
└─────────────────────────────────────────────────────────────────────────┘

🎯 SLICE INTERNALS:
┌─────────────────────────────────────────────────────────────────────────┐
│                    Slice Structure                                      │
│ ┌─────────┬─────────┬─────────┐                                         │
│ │ Pointer │ Length  │Capacity │                                         │
│ │   ptr   │   len   │   cap   │                                         │
│ └─────────┴─────────┴─────────┘                                         │
│      │                                                                  │
│      ▼                                                                  │
│ ┌─────┬─────┬─────┬─────┬─────┐  ← Underlying Array                     │
│ │  1  │  2  │  3  │  4  │  5  │                                         │
│ └─────┴─────┴─────┴─────┴─────┘                                         │
└─────────────────────────────────────────────────────────────────────────┘

🚨 COMMON GOTCHAS:
❌ Nil slice vs empty slice confusion
❌ Slice sharing underlying array (modifications affect both)
❌ Append can reallocate (pointer changes)
❌ Slicing doesn't free memory of original array

💡 MEMORY MANAGEMENT:
• append() may create new underlying array if capacity exceeded
• Multiple slices can share same underlying array
• Slicing large arrays can prevent garbage collection

🔧 BEST PRACTICES:
• Use slices instead of arrays in most cases
• Pre-allocate capacity if you know the size: make([]int, 0, expectedSize)
• Be careful with slice sharing and modifications
• Use copy() when you need independent slices

🎯 WHEN TO USE:
✅ Dynamic collections
✅ Function parameters (instead of arrays)
✅ Growing/shrinking data
✅ Most general-purpose collections

❌ When you need fixed size
❌ When every byte of memory matters
❌ When you need value semantics

=============================================================================
*/