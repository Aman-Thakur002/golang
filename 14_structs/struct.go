/*
=============================================================================
                           🏗️ GO STRUCTS TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Structs are Go's way to group related data together - like a blueprint for objects.
Think of it as a "custom data type" that holds multiple fields.

🔑 KEY FEATURES:
• Group multiple fields of different types
• Create custom types for your domain
• Methods can be attached to structs
• Struct embedding = composition (like inheritance)

💡 REAL-WORLD ANALOGY:
Struct = House Blueprint
- Fields = Rooms (bedroom, kitchen, etc.)
- Methods = Actions you can do (clean, renovate)
- Embedding = Adding an extension to existing house

🎯 WHY USE STRUCTS?
• Organize related data logically
• Create domain-specific types
• Enable object-oriented programming patterns
• Better code organization and readability

=============================================================================
*/

package main

//user defined data-structure
// - use to group mutiple fields

import (
	"fmt"
	"time"
)

// 👤 SIMPLE STRUCT: Groups customer information
type customer struct {
	name string   // Field 1: customer name
	phone string  // Field 2: customer phone
}

// 📦 COMPLEX STRUCT: Groups order information + embedded struct
type order struct {
	id        int       // Field 1: order ID
	amount    float32   // Field 2: order amount
	status    string    // Field 3: order status
	createdAt time.Time // Field 4: timestamp (nano second precision)
	customer            // Field 5: STRUCT EMBEDDING - referencing customer struct
	                    // This gives order access to all customer fields!
}

// 🏭 CONSTRUCTOR PATTERN: Function that creates and returns a struct
// This is Go's way of having "constructors" (Go doesn't have built-in constructors)
func newOrder(id int, amount float32, status string) order{
	 myOrder := order {
		id : id,           // Initialize each field
		amount : amount,
		status : status,
	 }

	 return myOrder  // Return the created struct
}

// 🔧 STRUCT METHODS: Functions that belong to a struct type
// IMPORTANT: Use *order (pointer) when you need to MODIFY the struct
func (o *order) changeStatus(status string) { // (receiver) methodName(params)
	o.status = status   // Modifies the original struct (no need to dereference *)
	                    // Go automatically handles pointer dereferencing for structs
}

// 📖 STRUCT METHODS: Use order (value) when you only READ from struct
func (o order) getAmount() float32 {
	return o.amount  // Just reading, no modification needed
}

func main() {
    
	// 🏗️ METHOD 1: Create customer separately, then use in order
	// newCustomer := customer{
	// 	name : "Thakur",
	// 	phone : "98237429",
	// }

	// 📦 CREATING STRUCT WITH EMBEDDED STRUCT
	myOrder := order{
		id:     1,
		amount: 100.0,
		status: "pending",
		// customer : newCustomer,              // 1st approach: use pre-created customer
		customer : customer{                    // 2nd approach: create customer inline
			name:  "Thakur",
			phone : "8923649823",                 
		},
	}

	fmt.Println(myOrder)  // Print entire struct

	// 🔧 USING STRUCT METHODS
	myOrder.changeStatus("Confirmed") // Call method to modify struct
	myOrder.createdAt = time.Now()    // Direct field assignment

	fmt.Println(myOrder)
	fmt.Println(myOrder.getAmount())  // Call method to read from struct

	// 🏭 USING CONSTRUCTOR FUNCTION
	o1 := newOrder(1,200,"Approved")
	fmt.Println(o1)
	o2 := newOrder(2,500,"Cancelled")
	fmt.Println(o2)

	// 🚀 ANONYMOUS STRUCT: One-time use struct (no type definition needed)
	language := struct {
		name string    // Define fields inline
		isGood bool
	} {"Golang", true}  // Initialize values immediately

	fmt.Println(language)

	// 🔍 ACCESSING EMBEDDED STRUCT FIELDS
	// Because customer is embedded, we can access its fields directly:
	fmt.Println("Customer name:", myOrder.name)   // Direct access to embedded field
	fmt.Println("Customer phone:", myOrder.phone) // No need for myOrder.customer.phone
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

🏗️ STRUCT CREATION PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Method 1: Field by field                                             │
│ var p Person                                                            │
│ p.name = "John"                                                         │
│                                                                         │
│ // Method 2: Struct literal                                             │
│ p := Person{name: "John", age: 30}                                      │
│                                                                         │
│ // Method 3: Constructor function                                       │
│ p := NewPerson("John", 30)                                              │
│                                                                         │
│ // Method 4: Anonymous struct                                           │
│ p := struct{name string}{"John"}                                        │
└─────────────────────────────────────────────────────────────────────────┘

🔧 METHOD RECEIVERS:
┌─────────────────┬─────────────────┬─────────────────────────────────────┐
│   Receiver      │   When to Use   │           Example                   │
├─────────────────┼─────────────────┼─────────────────────────────────────┤
│ (s Struct)      │ Read-only       │ func (s Student) getName() string  │
│ (s *Struct)     │ Modify struct   │ func (s *Student) setName(n string)│
└─────────────────┴─────────────────┴─────────────────────────────────────┘

🎯 STRUCT EMBEDDING BENEFITS:
✅ Composition over inheritance
✅ Access embedded fields directly
✅ Embedded methods become available
✅ Multiple embedding possible

💡 COMMON PATTERNS:
• Configuration structs
• Request/Response structs (APIs)
• Database models
• Event structures

🚨 GOTCHAS:
❌ Structs are value types (copied when assigned)
❌ Zero value of struct has zero values for all fields
❌ Exported fields start with capital letter
❌ Method receiver choice affects behavior

🔧 BEST PRACTICES:
• Use pointer receivers for methods that modify
• Use value receivers for small structs or read-only methods
• Group related fields together
• Use embedding for "is-a" relationships
• Constructor functions for complex initialization

🏷️ STRUCT TAGS (Advanced):
type User struct {
    Name string `json:"name" db:"user_name"`  // Metadata for serialization
    Age  int    `json:"age" validate:"min=0"`
}

=============================================================================
*/