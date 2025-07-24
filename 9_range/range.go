/*
=============================================================================
                           🔄 GO RANGE TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Range is Go's way to iterate over data structures in a clean, readable way.
Think of it as a "for each" loop that works with different collection types.

🔑 KEY FEATURES:
• Works with slices, arrays, maps, strings, channels
• Returns different values based on data type
• Can ignore values using underscore (_)
• More readable than traditional for loops

💡 REAL-WORLD ANALOGY:
Range = Walking through different types of collections
- Slice/Array: Walking through a line of people (index, person)
- Map: Walking through a phone book (name, number)
- String: Walking through letters in a word (position, character)

🎯 WHY USE RANGE?
• Cleaner syntax than traditional loops
• Automatic handling of collection bounds
• Works consistently across different data types
• Less error-prone (no off-by-one errors)

=============================================================================
*/

package main

import "fmt"

// use for iteration over data structures

func main() {
	fmt.Println("🔄 RANGE LEARNING JOURNEY")
	fmt.Println("=========================")

	nums := []int{1, 3, 5}

	// 🔢 TRADITIONAL FOR LOOP (verbose way)
	fmt.Println("📝 Traditional for loop:")
	for i := 0; i < len(nums); i++ {
		fmt.Printf("  Index: %d, Value: %d\n", i, nums[i])
	}

	fmt.Println("\n🎯 RANGE WITH SLICES/ARRAYS")
	fmt.Println("============================")
	
	// 🎯 USING RANGE: Much cleaner!
	for i, num := range nums { // index, value
		fmt.Printf("  Index: %d, Value: %d\n", i, num)
	}

	// 💡 IGNORE INDEX: Use underscore when you don't need it
	fmt.Println("\n🎯 Range - Values only:")
	for _, num := range nums {  // _ ignores the index
		fmt.Printf("  Value: %d\n", num)
	}

	// 💡 IGNORE VALUE: Get only indices
	fmt.Println("\n🎯 Range - Indices only:")
	for i := range nums {  // Only index, no second variable
		fmt.Printf("  Index: %d\n", i)
	}

	fmt.Println("\n🗺️ RANGE WITH MAPS")
	fmt.Println("===================")
	
	// 🗺️ ITERATION OVER MAPS
	m := map[string]string{"name": "aman", "occupation": "backend engineer"}
	for k, v := range m { // key, value
		fmt.Printf("  %s: %s\n", k, v)
	}

	// 💡 KEYS ONLY from map
	fmt.Println("\n🔑 Map keys only:")
	for key := range m {
		fmt.Printf("  Key: %s\n", key)
	}

	fmt.Println("\n📝 RANGE WITH STRINGS")
	fmt.Println("======================")
	
	// 📝 ITERATION OVER STRING
	// c is the unicode of every character, e.g for A unicode is 65
	// unicode point rune
	// if unicode <=255 -> 1 byte, if unicode is bigger then 255 then it takes more than 1 byte so it changes the index of other character, e.g if string "AM", unicode of A is 300, i=0 then index of M i.e i would be 2 if A is taking 2 bytes
	for i, c := range "Aman Pratap" {  // i is starting byte index of rune
		fmt.Printf("  Byte index: %d, Unicode: %d, Character: %c\n", i, c, c)
	}

	fmt.Println("\n🌍 RANGE WITH UNICODE")
	fmt.Println("======================")
	
	// 🌍 UNICODE EXAMPLE: Shows byte vs character difference
	for i, c := range "Hello 世界" {  // Mixed ASCII and Unicode
		fmt.Printf("  Byte index: %d, Unicode: %d, Character: %c\n", i, c, c)
	}

	fmt.Println("\n🔢 RANGE WITH NUMBERS (Go 1.22+)")
	fmt.Println("==================================")
	
	// 🔢 RANGE OVER INTEGERS: New in Go 1.22
	fmt.Println("Counting 0 to 4:")
	for i := range 5 {  // Iterates from 0 to 4
		fmt.Printf("  %d ", i)
	}
	fmt.Println()

	fmt.Println("\n📡 RANGE WITH CHANNELS")
	fmt.Println("======================")
	
	// 📡 RANGE WITH CHANNELS: Receives until channel is closed
	ch := make(chan int, 3)
	ch <- 10
	ch <- 20
	ch <- 30
	close(ch)  // Must close to end the range loop
	
	fmt.Println("Channel values:")
	for value := range ch {  // Receives until channel closed
		fmt.Printf("  Received: %d\n", value)
	}
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

🔄 RANGE RETURN VALUES BY TYPE:
┌─────────────────┬─────────────────┬─────────────────────────────────────┐
│   Data Type     │   First Value   │           Second Value              │
├─────────────────┼─────────────────┼─────────────────────────────────────┤
│ Slice/Array     │ Index (int)     │ Element value                       │
│ Map             │ Key             │ Value                               │
│ String          │ Byte index      │ Rune (Unicode code point)          │
│ Channel         │ Value           │ (none - only one value)             │
│ Integer (1.22+) │ Value (0 to n-1)│ (none - only one value)             │
└─────────────────┴─────────────────┴─────────────────────────────────────┘

🎯 RANGE PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Both index and value                                                 │
│ for i, v := range slice { }                                             │
│                                                                         │
│ // Value only (ignore index)                                            │
│ for _, v := range slice { }                                             │
│                                                                         │
│ // Index only (ignore value)                                            │
│ for i := range slice { }                                                │
│                                                                         │
│ // Key and value (maps)                                                 │
│ for k, v := range myMap { }                                             │
│                                                                         │
│ // Keys only (maps)                                                     │
│ for k := range myMap { }                                                │
└─────────────────────────────────────────────────────────────────────────┘

📝 STRING ITERATION DETAILS:
• Range over string iterates by runes (Unicode code points), not bytes
• Index is byte position, not character position
• ASCII characters: 1 byte each
• Unicode characters: 1-4 bytes each
• Use []byte(string) to iterate by bytes instead

🚨 GOTCHAS:
❌ Map iteration order is random (not guaranteed)
❌ String range gives runes, not bytes
❌ Modifying slice during range can cause issues
❌ Range variable is reused (be careful with goroutines)

💡 PERFORMANCE NOTES:
• Range is generally as fast as traditional for loops
• For strings, range is more efficient than manual rune handling
• For large slices, consider if you need both index and value

🔧 BEST PRACTICES:
• Use range instead of traditional for loops when possible
• Use _ to ignore unused values (cleaner code)
• Be aware of Unicode vs byte differences in strings
• Don't modify collection while ranging over it
• Use descriptive variable names (not just i, v)

🎯 WHEN TO USE RANGE:
✅ Iterating over any collection
✅ When you need clean, readable loops
✅ Processing all elements in a collection
✅ When bounds checking is important

❌ When you need complex loop control (break to specific labels)
❌ When you need to modify the collection during iteration
❌ When performance is critical and you need manual optimization

=============================================================================
*/