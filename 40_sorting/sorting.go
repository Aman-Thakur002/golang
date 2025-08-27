/*
=============================================================================
                           📊 GO SORTING TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Go's sort package provides functions for sorting slices and user-defined
collections. It includes built-in sorting for common types and interfaces
for custom sorting logic.

🔑 KEY FEATURES:
• Built-in sorting for basic types
• Custom sorting with sort.Interface
• Stable and unstable sorting
• Binary search functions
• Reverse sorting

💡 REAL-WORLD ANALOGY:
Sorting = Organizing Library Books
- sort.Ints = Arranging by page numbers
- sort.Strings = Alphabetical arrangement
- Custom sorting = Arranging by genre, author, or date
- Stable sort = Maintaining original order for equal items

🎯 WHY LEARN SORTING?
• Data organization and presentation
• Efficient searching algorithms
• Database-like operations
• Performance optimization

=============================================================================
*/

package main

import (
	"fmt"
	"sort"
	"strings"
)

// 📊 CUSTOM TYPES FOR SORTING
type Person struct {
	Name string
	Age  int
	City string
}

type People []Person

// Implement sort.Interface for People
func (p People) Len() int           { return len(p) }
func (p People) Less(i, j int) bool { return p[i].Age < p[j].Age }
func (p People) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// ByName implements sort.Interface for []Person based on Name field
type ByName []Person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	fmt.Println("📊 SORTING TUTORIAL")
	fmt.Println("===================")

	// 🎯 DEMO 1: Basic Type Sorting
	fmt.Println("\n🎯 DEMO 1: Basic Type Sorting")
	fmt.Println("=============================")

	// Sort integers
	numbers := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("Original integers: %v\n", numbers)
	sort.Ints(numbers)
	fmt.Printf("Sorted integers:   %v\n", numbers)

	// Sort strings
	words := []string{"banana", "apple", "cherry", "date", "elderberry"}
	fmt.Printf("\nOriginal strings: %v\n", words)
	sort.Strings(words)
	fmt.Printf("Sorted strings:   %v\n", words)

	// Sort floats
	prices := []float64{19.99, 5.50, 12.75, 3.25, 25.00}
	fmt.Printf("\nOriginal floats: %v\n", prices)
	sort.Float64s(prices)
	fmt.Printf("Sorted floats:   %v\n", prices)

	// 🎯 DEMO 2: Reverse Sorting
	fmt.Println("\n🎯 DEMO 2: Reverse Sorting")
	fmt.Println("==========================")

	numbers2 := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("Original: %v\n", numbers2)
	
	// Sort in reverse order
	sort.Sort(sort.Reverse(sort.IntSlice(numbers2)))
	fmt.Printf("Reverse:  %v\n", numbers2)

	// Alternative: sort then reverse
	words2 := []string{"banana", "apple", "cherry", "date"}
	sort.Strings(words2)
	fmt.Printf("Sorted strings: %v\n", words2)
	sort.Sort(sort.Reverse(sort.StringSlice(words2)))
	fmt.Printf("Reverse strings: %v\n", words2)

	// 🎯 DEMO 3: Custom Sorting with sort.Slice
	fmt.Println("\n🎯 DEMO 3: Custom Sorting with sort.Slice")
	fmt.Println("=========================================")

	people := []Person{
		{"Alice", 30, "New York"},
		{"Bob", 25, "Los Angeles"},
		{"Charlie", 35, "Chicago"},
		{"Diana", 28, "Houston"},
	}

	fmt.Println("Original people:")
	for _, p := range people {
		fmt.Printf("  %s (%d) from %s\n", p.Name, p.Age, p.City)
	}

	// Sort by age
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println("\nSorted by age:")
	for _, p := range people {
		fmt.Printf("  %s (%d) from %s\n", p.Name, p.Age, p.City)
	}

	// Sort by name
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})

	fmt.Println("\nSorted by name:")
	for _, p := range people {
		fmt.Printf("  %s (%d) from %s\n", p.Name, p.Age, p.City)
	}

	// Sort by city length
	sort.Slice(people, func(i, j int) bool {
		return len(people[i].City) < len(people[j].City)
	})

	fmt.Println("\nSorted by city name length:")
	for _, p := range people {
		fmt.Printf("  %s (%d) from %s (%d chars)\n", p.Name, p.Age, p.City, len(p.City))
	}

	// 🎯 DEMO 4: Implementing sort.Interface
	fmt.Println("\n🎯 DEMO 4: sort.Interface Implementation")
	fmt.Println("=======================================")

	people2 := People{
		{"Eve", 32, "Boston"},
		{"Frank", 27, "Seattle"},
		{"Grace", 29, "Denver"},
		{"Henry", 31, "Miami"},
	}

	fmt.Println("Original (People type):")
	for _, p := range people2 {
		fmt.Printf("  %s (%d) from %s\n", p.Name, p.Age, p.City)
	}

	// Sort using our People type (sorts by age)
	sort.Sort(people2)
	fmt.Println("\nSorted by age (using sort.Interface):")
	for _, p := range people2 {
		fmt.Printf("  %s (%d) from %s\n", p.Name, p.Age, p.City)
	}

	// Sort by name using ByName type
	sort.Sort(ByName(people2))
	fmt.Println("\nSorted by name (using ByName type):")
	for _, p := range people2 {
		fmt.Printf("  %s (%d) from %s\n", p.Name, p.Age, p.City)
	}

	// 🎯 DEMO 5: Stable Sorting
	fmt.Println("\n🎯 DEMO 5: Stable vs Unstable Sorting")
	fmt.Println("=====================================")

	type Student struct {
		Name  string
		Grade int
		Class string
	}

	students := []Student{
		{"Alice", 85, "A"},
		{"Bob", 90, "B"},
		{"Charlie", 85, "A"},
		{"Diana", 90, "B"},
		{"Eve", 85, "C"},
	}

	fmt.Println("Original students:")
	for i, s := range students {
		fmt.Printf("  %d: %s (Grade: %d, Class: %s)\n", i, s.Name, s.Grade, s.Class)
	}

	// Stable sort by grade
	sort.SliceStable(students, func(i, j int) bool {
		return students[i].Grade < students[j].Grade
	})

	fmt.Println("\nStable sort by grade (maintains relative order):")
	for i, s := range students {
		fmt.Printf("  %d: %s (Grade: %d, Class: %s)\n", i, s.Name, s.Grade, s.Class)
	}

	// 🎯 DEMO 6: Multi-level Sorting
	fmt.Println("\n🎯 DEMO 6: Multi-level Sorting")
	fmt.Println("==============================")

	products := []struct {
		Name     string
		Category string
		Price    float64
		Rating   float64
	}{
		{"Laptop", "Electronics", 999.99, 4.5},
		{"Phone", "Electronics", 699.99, 4.2},
		{"Book", "Education", 29.99, 4.8},
		{"Tablet", "Electronics", 399.99, 4.0},
		{"Notebook", "Education", 5.99, 4.3},
	}

	fmt.Println("Original products:")
	for _, p := range products {
		fmt.Printf("  %s (%s) - $%.2f (%.1f★)\n", p.Name, p.Category, p.Price, p.Rating)
	}

	// Sort by category first, then by price within category
	sort.Slice(products, func(i, j int) bool {
		if products[i].Category != products[j].Category {
			return products[i].Category < products[j].Category
		}
		return products[i].Price < products[j].Price
	})

	fmt.Println("\nSorted by category, then by price:")
	for _, p := range products {
		fmt.Printf("  %s (%s) - $%.2f (%.1f★)\n", p.Name, p.Category, p.Price, p.Rating)
	}

	// 🎯 DEMO 7: Checking if Sorted
	fmt.Println("\n🎯 DEMO 7: Checking if Sorted")
	fmt.Println("=============================")

	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{5, 2, 8, 1, 9}

	fmt.Printf("nums1 %v is sorted: %t\n", nums1, sort.IntsAreSorted(nums1))
	fmt.Printf("nums2 %v is sorted: %t\n", nums2, sort.IntsAreSorted(nums2))

	words3 := []string{"apple", "banana", "cherry"}
	words4 := []string{"zebra", "apple", "banana"}

	fmt.Printf("words3 %v is sorted: %t\n", words3, sort.StringsAreSorted(words3))
	fmt.Printf("words4 %v is sorted: %t\n", words4, sort.StringsAreSorted(words4))

	// Check if custom slice is sorted
	ages := []int{25, 30, 35, 40}
	isSorted := sort.SliceIsSorted(ages, func(i, j int) bool {
		return ages[i] < ages[j]
	})
	fmt.Printf("ages %v is sorted: %t\n", ages, isSorted)

	// 🎯 DEMO 8: Binary Search
	fmt.Println("\n🎯 DEMO 8: Binary Search")
	fmt.Println("========================")

	sortedNumbers := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	fmt.Printf("Sorted array: %v\n", sortedNumbers)

	// Search for existing elements
	targets := []int{5, 11, 20, 1}
	for _, target := range targets {
		index := sort.SearchInts(sortedNumbers, target)
		if index < len(sortedNumbers) && sortedNumbers[index] == target {
			fmt.Printf("Found %d at index %d\n", target, index)
		} else {
			fmt.Printf("%d not found (would be inserted at index %d)\n", target, index)
		}
	}

	// Binary search in strings
	sortedWords := []string{"apple", "banana", "cherry", "date", "elderberry"}
	fmt.Printf("\nSorted words: %v\n", sortedWords)
	
	searchWords := []string{"banana", "fig", "apple"}
	for _, word := range searchWords {
		index := sort.SearchStrings(sortedWords, word)
		if index < len(sortedWords) && sortedWords[index] == word {
			fmt.Printf("Found '%s' at index %d\n", word, index)
		} else {
			fmt.Printf("'%s' not found (would be inserted at index %d)\n", word, index)
		}
	}

	// 🎯 DEMO 9: Custom Comparisons
	fmt.Println("\n🎯 DEMO 9: Custom Comparisons")
	fmt.Println("=============================")

	// Case-insensitive string sorting
	mixedCase := []string{"Banana", "apple", "Cherry", "date", "Elderberry"}
	fmt.Printf("Original mixed case: %v\n", mixedCase)

	sort.Slice(mixedCase, func(i, j int) bool {
		return strings.ToLower(mixedCase[i]) < strings.ToLower(mixedCase[j])
	})
	fmt.Printf("Case-insensitive sort: %v\n", mixedCase)

	// Sort by string length
	lengths := []string{"a", "hello", "go", "programming", "sort"}
	fmt.Printf("\nOriginal by length: %v\n", lengths)

	sort.Slice(lengths, func(i, j int) bool {
		if len(lengths[i]) != len(lengths[j]) {
			return len(lengths[i]) < len(lengths[j])
		}
		return lengths[i] < lengths[j] // Secondary sort alphabetically
	})
	fmt.Printf("Sorted by length: %v\n", lengths)

	// 🎯 DEMO 10: Performance Comparison
	fmt.Println("\n🎯 DEMO 10: Performance Notes")
	fmt.Println("=============================")

	fmt.Println("Sorting algorithm characteristics:")
	fmt.Println("• Go uses introsort (hybrid of quicksort, heapsort, and insertion sort)")
	fmt.Println("• Average time complexity: O(n log n)")
	fmt.Println("• Worst case time complexity: O(n log n)")
	fmt.Println("• Space complexity: O(log n)")
	fmt.Println("• Not stable by default (use sort.SliceStable for stable sorting)")

	fmt.Println("\nWhen to use each sorting method:")
	fmt.Println("• sort.Ints/Strings/Float64s: For basic types, fastest")
	fmt.Println("• sort.Slice: For custom sorting logic, most flexible")
	fmt.Println("• sort.SliceStable: When you need stable sorting")
	fmt.Println("• sort.Interface: For reusable sorting types")

	fmt.Println("\n✨ All sorting demos completed!")
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

📊 BASIC SORTING FUNCTIONS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Built-in type sorting                                                │
│ sort.Ints([]int)           // Sort integers                             │
│ sort.Float64s([]float64)   // Sort floats                               │
│ sort.Strings([]string)     // Sort strings                              │
│                                                                         │
│ // Check if sorted                                                      │
│ sort.IntsAreSorted([]int)     // Check if ints are sorted               │
│ sort.Float64sAreSorted([]float64) // Check if floats are sorted         │
│ sort.StringsAreSorted([]string)   // Check if strings are sorted        │
│                                                                         │
│ // Binary search                                                        │
│ sort.SearchInts([]int, target)       // Binary search in ints           │
│ sort.SearchFloat64s([]float64, target) // Binary search in floats       │
│ sort.SearchStrings([]string, target)   // Binary search in strings      │
└─────────────────────────────────────────────────────────────────────────┘

🔧 CUSTOM SORTING:
┌─────────────────────────────────────────────────────────────────────────┐
│ // sort.Slice - most common for custom sorting                          │
│ sort.Slice(slice, func(i, j int) bool {                                 │
│     return slice[i] < slice[j]  // Define comparison                    │
│ })                                                                      │
│                                                                         │
│ // sort.SliceStable - stable sorting                                    │
│ sort.SliceStable(slice, func(i, j int) bool {                           │
│     return slice[i] < slice[j]                                          │
│ })                                                                      │
│                                                                         │
│ // Check if custom slice is sorted                                      │
│ sort.SliceIsSorted(slice, func(i, j int) bool {                         │
│     return slice[i] < slice[j]                                          │
│ })                                                                      │
│                                                                         │
│ // Binary search in custom slice                                        │
│ sort.Search(len(slice), func(i int) bool {                              │
│     return slice[i] >= target                                           │
│ })                                                                      │
└─────────────────────────────────────────────────────────────────────────┘

🎯 SORT.INTERFACE:
┌─────────────────────────────────────────────────────────────────────────┐
│ type Interface interface {                                              │
│     Len() int           // Number of elements                           │
│     Less(i, j int) bool // Compare elements at i and j                  │
│     Swap(i, j int)      // Swap elements at i and j                     │
│ }                                                                       │
│                                                                         │
│ // Example implementation                                               │
│ type People []Person                                                    │
│                                                                         │
│ func (p People) Len() int           { return len(p) }                   │
│ func (p People) Less(i, j int) bool { return p[i].Age < p[j].Age }      │
│ func (p People) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }         │
│                                                                         │
│ // Usage                                                                │
│ sort.Sort(people)                                                       │
└─────────────────────────────────────────────────────────────────────────┘

🔄 REVERSE SORTING:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Reverse built-in types                                               │
│ sort.Sort(sort.Reverse(sort.IntSlice(ints)))                            │
│ sort.Sort(sort.Reverse(sort.StringSlice(strings)))                      │
│ sort.Sort(sort.Reverse(sort.Float64Slice(floats)))                      │
│                                                                         │
│ // Reverse custom sorting                                               │
│ sort.Slice(slice, func(i, j int) bool {                                 │
│     return slice[i] > slice[j]  // Note: > instead of <                 │
│ })                                                                      │
│                                                                         │
│ // Reverse sort.Interface                                               │
│ sort.Sort(sort.Reverse(customType))                                     │
└─────────────────────────────────────────────────────────────────────────┘

🔍 BINARY SEARCH:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Generic binary search                                                │
│ index := sort.Search(len(slice), func(i int) bool {                     │
│     return slice[i] >= target                                           │
│ })                                                                      │
│                                                                         │
│ // Check if found                                                       │
│ if index < len(slice) && slice[index] == target {                       │
│     // Found at index                                                   │
│ } else {                                                                │
│     // Not found, would be inserted at index                           │
│ }                                                                       │
│                                                                         │
│ // Find insertion point                                                 │
│ insertIndex := sort.Search(len(slice), func(i int) bool {               │
│     return slice[i] > target                                            │
│ })                                                                      │
└─────────────────────────────────────────────────────────────────────────┘

📋 MULTI-LEVEL SORTING:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Sort by multiple criteria                                            │
│ sort.Slice(people, func(i, j int) bool {                                │
│     // Primary sort by age                                              │
│     if people[i].Age != people[j].Age {                                 │
│         return people[i].Age < people[j].Age                            │
│     }                                                                   │
│     // Secondary sort by name                                           │
│     return people[i].Name < people[j].Name                              │
│ })                                                                      │
│                                                                         │
│ // Three-level sorting                                                  │
│ sort.Slice(items, func(i, j int) bool {                                 │
│     if items[i].Category != items[j].Category {                         │
│         return items[i].Category < items[j].Category                    │
│     }                                                                   │
│     if items[i].Priority != items[j].Priority {                         │
│         return items[i].Priority > items[j].Priority // Desc           │
│     }                                                                   │
│     return items[i].Name < items[j].Name                                │
│ })                                                                      │
└─────────────────────────────────────────────────────────────────────────┘

💡 BEST PRACTICES:
• Use built-in functions (sort.Ints, etc.) for basic types
• Use sort.Slice for most custom sorting needs
• Use sort.SliceStable when order of equal elements matters
• Implement sort.Interface for reusable sorting logic
• Always sort before binary search
• Consider performance implications of complex comparisons

🚨 COMMON MISTAKES:
❌ Forgetting to sort before binary search
❌ Using unstable sort when stability is needed
❌ Complex comparison functions that hurt performance
❌ Not handling edge cases in custom comparisons
❌ Assuming sort is stable by default

⚡ PERFORMANCE TIPS:
• Go's sort is very efficient (introsort algorithm)
• Avoid complex operations in comparison functions
• Use sort.SliceStable only when needed (slightly slower)
• Consider pre-computing expensive comparison values
• Profile sorting-heavy code

🎯 ALGORITHM COMPLEXITY:
┌─────────────────┬─────────────────┬─────────────────┬─────────────────┐
│   Operation     │   Time (Avg)    │   Time (Worst)  │     Space       │
├─────────────────┼─────────────────┼─────────────────┼─────────────────┤
│ Sort            │ O(n log n)      │ O(n log n)      │ O(log n)        │
│ Binary Search   │ O(log n)        │ O(log n)        │ O(1)            │
│ Is Sorted       │ O(n)            │ O(n)            │ O(1)            │
└─────────────────┴─────────────────┴─────────────────┴─────────────────┘

🎯 WHEN TO USE EACH:
• sort.Ints/Strings/Float64s: Basic types, maximum performance
• sort.Slice: Custom sorting logic, most flexible
• sort.SliceStable: When relative order of equal elements matters
• sort.Interface: Reusable sorting for custom types
• Binary search: Fast lookups in sorted data

=============================================================================
*/