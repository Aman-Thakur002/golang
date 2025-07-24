/*
=============================================================================
                           🏷️ GO ENUMS TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Go doesn't have built-in enums, but we can create enum-like behavior using
custom types and constants. Think of them as "named choices" from a fixed set.

🔑 KEY FEATURES:
• Custom types with predefined constant values
• Two main approaches: iota (integer) and string-based
• Type safety prevents invalid values
• Self-documenting code with meaningful names

💡 REAL-WORLD ANALOGY:
Enum = Restaurant Menu Categories
- Type = "Menu Section" (appetizers, mains, desserts)
- Constants = Specific items in each section
- Type safety = Can't order "dessert" from "appetizer" section
- iota = Auto-numbering menu items (1, 2, 3...)

🎯 WHY USE ENUMS?
• Limit values to predefined set
• Make code more readable and maintainable
• Prevent invalid state values
• Enable exhaustive switch statements

=============================================================================
*/

package main

import "fmt"

// enums -> enumarated variables

// 🔢 INTEGER-BASED ENUM: Using iota for auto-incrementing values
// type orderStatus int
// const (
//  Received orderStatus = iota // iota is a predeclared identifier representing the untyped integer ordinal number of the current const specification in a (usually parenthesized) const declaration. It is zero-indexed.
//  Confirmed
//  Prepared
//  Delivered
// )

// 📝 STRING-BASED ENUM: More readable and debuggable
type orderStatus string
const (
	Received  orderStatus = "Received"   // 📦 Order just received
	Confirmed orderStatus = "Confirmed"  // ✅ Order confirmed by customer
	Prepared  orderStatus = "Prepared"   // 👨‍🍳 Order prepared by kitchen
	Delivered orderStatus = "Delivered"  // 🚚 Order delivered to customer
)

// 🎯 ANOTHER ENUM EXAMPLE: Priority levels using iota
type Priority int
const (
	Low Priority = iota    // 0 - Low priority
	Medium                 // 1 - Medium priority (iota auto-increments)
	High                   // 2 - High priority
	Critical               // 3 - Critical priority
)

// 🌈 COLOR ENUM: Using iota with custom values
type Color int
const (
	Red Color = iota + 1   // 1 - Start from 1 instead of 0
	Green                  // 2
	Blue                   // 3
	Yellow                 // 4
)

// 🔧 METHODS ON ENUM TYPES: Add behavior to enums
func (os orderStatus) String() string {
	return string(os)  // Convert to string for printing
}

func (os orderStatus) IsComplete() bool {
	return os == Delivered  // Check if order is complete
}

func (p Priority) String() string {
	switch p {
	case Low:
		return "Low Priority"
	case Medium:
		return "Medium Priority"
	case High:
		return "High Priority"
	case Critical:
		return "Critical Priority"
	default:
		return "Unknown Priority"
	}
}

// 🎯 FUNCTION USING ENUM: Type-safe parameter
func changeOrderStatus(status orderStatus) {
	fmt.Println("Changed Order Status:", status)
	
	// 🔄 EXHAUSTIVE SWITCH: Handle all enum values
	switch status {
	case Received:
		fmt.Println("📦 Processing new order...")
	case Confirmed:
		fmt.Println("✅ Order confirmed, preparing...")
	case Prepared:
		fmt.Println("👨‍🍳 Order ready for delivery...")
	case Delivered:
		fmt.Println("🚚 Order completed!")
	}
}

func processPriority(p Priority) {
	fmt.Printf("Processing %s task\n", p)
	
	if p >= High {  // 💡 Can compare enum values
		fmt.Println("⚡ High priority - handle immediately!")
	}
}

func main() {
	fmt.Println("🏷️ ENUMS LEARNING JOURNEY")
	fmt.Println("=========================")

	fmt.Println("\n🎯 STRING-BASED ENUM USAGE")
	fmt.Println("===========================")

	// ✅ TYPE-SAFE: Can only use predefined values
	changeOrderStatus(Confirmed)
	changeOrderStatus(Delivered)

	fmt.Println("\n🎯 ENUM METHODS")
	fmt.Println("===============")

	currentStatus := Prepared
	fmt.Println("Current status:", currentStatus.String())
	fmt.Println("Is complete?", currentStatus.IsComplete())

	deliveredStatus := Delivered
	fmt.Println("Delivered status:", deliveredStatus.String())
	fmt.Println("Is complete?", deliveredStatus.IsComplete())

	fmt.Println("\n🎯 INTEGER-BASED ENUM (iota)")
	fmt.Println("=============================")

	// 🔢 IOTA ENUM: Auto-incrementing integer values
	fmt.Println("Priority values:")
	fmt.Printf("Low: %d (%s)\n", Low, Low)
	fmt.Printf("Medium: %d (%s)\n", Medium, Medium)
	fmt.Printf("High: %d (%s)\n", High, High)
	fmt.Printf("Critical: %d (%s)\n", Critical, Critical)

	processPriority(High)
	processPriority(Low)

	fmt.Println("\n🎯 ENUM COMPARISON")
	fmt.Println("==================")

	// 🔍 COMPARING ENUM VALUES
	task1 := High
	task2 := Critical

	if task1 < task2 {
		fmt.Printf("%s has lower priority than %s\n", task1, task2)
	}

	fmt.Println("\n🎯 COLOR ENUM WITH CUSTOM START")
	fmt.Println("===============================")

	fmt.Printf("Red: %d\n", Red)     // 1 (started from 1)
	fmt.Printf("Green: %d\n", Green) // 2
	fmt.Printf("Blue: %d\n", Blue)   // 3
	fmt.Printf("Yellow: %d\n", Yellow) // 4

	// ❌ COMPILE ERROR PREVENTION: This would cause compile error
	// changeOrderStatus("InvalidStatus")  // Can't pass string directly
	// changeOrderStatus(123)              // Can't pass int directly
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

🏷️ ENUM PATTERNS IN GO:
┌─────────────────────────────────────────────────────────────────────────┐
│ // String-based enum (readable, debuggable)                            │
│ type Status string                                                      │
│ const (                                                                 │
│     Active   Status = "active"                                          │
│     Inactive Status = "inactive"                                        │
│ )                                                                       │
│                                                                         │
│ // Integer-based enum with iota (memory efficient)                     │
│ type State int                                                          │
│ const (                                                                 │
│     StateA State = iota  // 0                                           │
│     StateB               // 1                                           │
│     StateC               // 2                                           │
│ )                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

🔍 STRING vs INTEGER ENUMS:
┌─────────────────┬─────────────────┬─────────────────────────────────────┐
│    Aspect       │  String Enum    │         Integer Enum                │
├─────────────────┼─────────────────┼─────────────────────────────────────┤
│ Readability     │ Very readable   │ Need String() method                │
│ Memory usage    │ More memory     │ Less memory                         │
│ JSON/DB         │ Human readable  │ Compact storage                     │
│ Debugging       │ Self-explaining │ Need lookup                         │
│ Performance     │ String compare  │ Integer compare (faster)            │
└─────────────────┴─────────────────┴─────────────────────────────────────┘

⚡ IOTA FEATURES:
• Starts at 0 and increments by 1
• Resets to 0 in each const block
• Can be used in expressions: iota + 1, iota * 2
• Blank identifier _ skips values

🎯 ADVANCED IOTA PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Skip values                                                          │
│ const (                                                                 │
│     A = iota  // 0                                                      │
│     _         // 1 (skipped)                                            │
│     C         // 2                                                      │
│ )                                                                       │
│                                                                         │
│ // Custom expressions                                                   │
│ const (                                                                 │
│     KB = 1 << (10 * iota)  // 1024^0 = 1                               │
│     MB                     // 1024^1 = 1048576                         │
│     GB                     // 1024^2 = 1073741824                      │
│ )                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

🚨 GOTCHAS:
❌ No built-in enum validation (can cast any value)
❌ iota resets in each const block
❌ String enums use more memory than integers
❌ No automatic exhaustiveness checking in switch

💡 ENUM VALIDATION:
// Add validation method to prevent invalid values
func (os orderStatus) IsValid() bool {
    switch os {
    case Received, Confirmed, Prepared, Delivered:
        return true
    default:
        return false
    }
}

🔧 BEST PRACTICES:
• Use string enums for external APIs (JSON, databases)
• Use integer enums for internal performance-critical code
• Add String() method for integer enums
• Add validation methods for type safety
• Use exhaustive switch statements
• Group related constants in same const block

🎯 WHEN TO USE EACH TYPE:
✅ String enums: APIs, configuration, debugging
✅ Integer enums: Performance-critical, flags, internal state
✅ Custom iota: File permissions, bit flags, powers of 2

❌ Don't use enums for: Open-ended values, user input, dynamic sets

=============================================================================
*/