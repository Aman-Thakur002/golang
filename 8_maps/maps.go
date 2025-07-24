/*
=============================================================================
                           🗺️ GO MAPS TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Maps are Go's built-in hash table - key-value pairs for fast lookups.
Think of them as a "dictionary" where you look up values using keys.

🔑 KEY FEATURES:
• Fast O(1) average lookup time
• Keys must be comparable types (strings, numbers, booleans, arrays)
• Values can be any type
• Reference type (like slices)

💡 REAL-WORLD ANALOGY:
Map = Phone Book
- Key = Person's name
- Value = Phone number
- Lookup = Find number by name (very fast!)

🎯 WHY USE MAPS?
• Fast data retrieval by key
• Count occurrences of items
• Cache/memoization
• Configuration settings

=============================================================================
*/

package main

import (
	"fmt"
	"maps"
)

// maps -> hash, object, dictionary

func main(){

	fmt.Println("🗺️ MAPS LEARNING JOURNEY")
	fmt.Println("========================")

	// 🏗️ CREATING MAP: Using make()
	m := make(map[string]string)  // make(map[KeyType]ValueType)

	// 📝 SETTING ELEMENTS: map[key] = value
	m["name"] = "aman"
	m["age"] = "23"
  
	fmt.Println("📋 Full map:", m)
	fmt.Println("👤 Name:", m["name"])  // 🔍 ACCESSING: map[key]
	
	// ⚠️ IMPORTANT: If key doesn't exist, returns zero value
	fmt.Println("🏠 Address:", m["address"])  // Returns "" (empty string)

	// 🗑️ DELETE ELEMENT: delete(map, key)
	delete(m,"name")
	fmt.Println("📋 After deleting name:", m)

	// 🧹 CLEAR MAP: Removes all elements
	clear(m)
	fmt.Println("📋 After clearing:", m)

	// 🎯 MAP LITERAL: Create and initialize in one step
	 m1 := map[string]int{"price" : 30, "phone" : 3}  // Direct initialization
	 fmt.Println("💰 Price map:", m1)

	 // ✅ CHECK IF KEY EXISTS: The "comma ok" idiom
	 value, ok := m1["price"]  // Returns (value, exists_boolean)
	 fmt.Printf("💰 Price value: %d, exists: %t\n", value, ok)
	 
	 if ok {
		fmt.Println("✅ Price found!")
	 } else {
		fmt.Println("❌ Price not found!")
	 }

	 // 🔍 CHECKING NON-EXISTENT KEY
	 value2, ok2 := m1["discount"]
	 fmt.Printf("🏷️ Discount value: %d, exists: %t\n", value2, ok2)

    // 📦 MAPS PACKAGE: Utility functions (Go 1.21+)
	m2 := map[string]int{"price" : 20, "phones" : 2}
	m3 := map[string]int{"price" : 20, "phones" : 3}
	m4 := map[string]int{"price" : 20, "phones" : 2}
	
	fmt.Println("🔄 Comparing maps:")
	fmt.Printf("m2 == m3: %t\n", maps.Equal(m2,m3))  // false - different values
	fmt.Printf("m2 == m4: %t\n", maps.Equal(m2,m4))  // true - same key-value pairs

	// 🔄 ITERATING OVER MAPS
	fmt.Println("\n🔄 Iterating over map:")
	for key, value := range m1 {
		fmt.Printf("🔑 %s: %d\n", key, value)
	}

	// 🎯 NESTED MAPS: Maps containing maps
	userProfiles := map[string]map[string]string{
		"user1": {"name": "Alice", "city": "NYC"},
		"user2": {"name": "Bob", "city": "LA"},
	}
	fmt.Println("👥 User profiles:", userProfiles)
	fmt.Println("👤 User1 name:", userProfiles["user1"]["name"])
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

🗺️ MAP OPERATIONS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Create                                                               │
│ var m map[string]int        // nil map (can't write to it)             │
│ m = make(map[string]int)    // empty map (can write to it)             │
│ m := map[string]int{}       // empty map literal                       │
│ m := map[string]int{"a":1}  // initialized map literal                 │
│                                                                         │
│ // Operations                                                           │
│ m[key] = value              // set                                      │
│ value := m[key]             // get (returns zero value if not found)   │
│ value, ok := m[key]         // get with existence check                │
│ delete(m, key)              // delete                                   │
│ clear(m)                    // remove all elements                      │
└─────────────────────────────────────────────────────────────────────────┘

🔑 VALID KEY TYPES:
✅ Strings, numbers, booleans
✅ Arrays (but not slices)
✅ Structs with comparable fields
✅ Pointers
❌ Slices, maps, functions

📊 MAP vs SLICE COMPARISON:
┌─────────────────┬─────────────────┬─────────────────────────────────────┐
│    Aspect       │      Map        │              Slice                  │
├─────────────────┼─────────────────┼─────────────────────────────────────┤
│ Access          │ By key          │ By index                            │
│ Lookup Time     │ O(1) average    │ O(1) by index, O(n) by value       │
│ Ordering        │ No order        │ Maintains order                     │
│ Zero Value      │ nil             │ nil                                 │
│ Iteration       │ Random order    │ Sequential order                    │
└─────────────────┴─────────────────┴─────────────────────────────────────┘

🎯 COMMON PATTERNS:
• Counting: map[string]int for frequency counting
• Lookup tables: map[string]func() for dispatch tables
• Caching: map[string]Result for memoization
• Sets: map[string]bool (value doesn't matter)

🚨 GOTCHAS:
❌ Maps are not thread-safe (use sync.Map for concurrent access)
❌ Iteration order is random (not guaranteed)
❌ Zero value is nil (can't write to nil map)
❌ Comparing maps with == only works with nil

💡 MEMORY MANAGEMENT:
• Maps grow automatically as needed
• Deleting doesn't shrink the map
• Use clear() to reset but keep allocated memory
• Create new map to free memory completely

🔧 BEST PRACTICES:
• Use "comma ok" idiom to check key existence
• Initialize maps with make() or map literal
• Use meaningful key types (avoid interface{})
• Consider sync.Map for concurrent access
• Pre-size maps if you know approximate size

🎯 WHEN TO USE MAPS:
✅ Fast lookups by key
✅ Counting/frequency analysis
✅ Caching results
✅ Configuration data
✅ Implementing sets

❌ When you need ordered data
❌ When you need thread safety (without extra synchronization)
❌ When keys are not comparable

=============================================================================
*/