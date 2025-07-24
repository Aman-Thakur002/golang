/*
=============================================================================
                           🎯 GO INTERFACES TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Interfaces define WHAT a type can do, not HOW it does it.
Think of it as a "contract" - if you can do these things, you belong to this group!

🔑 KEY PRINCIPLES:
• Interfaces define behavior (methods), not data (fields)
• Types implement interfaces IMPLICITLY (no "implements" keyword)
• Empty interface{} can hold ANY type
• Interface composition = combining multiple interfaces

💡 REAL-WORLD ANALOGY:
Interface = Job Description
- "Must be able to: calculate area, find perimeter"
- Rectangle says: "I can do both!" ✅
- Circle says: "I can do both too!" ✅
- Both get hired for the "Shape" job!

🎯 WHY USE INTERFACES?
• Write functions that work with ANY type that has certain methods
• Polymorphism: same function, different behaviors
• Loose coupling: depend on behavior, not concrete types

=============================================================================
*/

package main

import (
	"fmt"
	"math"
)

// 🎯 INTERFACE DEFINITION: Defines what methods a "shape" must have
// Any type with an area() method automatically satisfies this interface
type shape interface { 
	area() float32  // Contract: "You must be able to calculate your area"
}

// 🎯 ANOTHER INTERFACE: Defines measurement capability
type measureable interface {
	perimeter() float32  // Contract: "You must be able to calculate perimeter"
}

// 🔗 INTERFACE COMPOSITION: Combining multiple interfaces
// Any type that satisfies BOTH shape AND measureable satisfies geometry
type geometery interface {    
	shape        // Must have area() method
	measureable  // Must have perimeter() method
	// Now requires BOTH methods to satisfy this interface
}

// 🟦 RECTANGLE TYPE: Concrete implementation
type rectangle struct {
	width, height float32  // Data fields
}

// 🎯 IMPLEMENTING INTERFACE METHODS:
// By having area() method, rectangle automatically satisfies 'shape' interface
func (r rectangle) area() float32 {
	return r.width * r.height  // Rectangle's way of calculating area
}

// By having perimeter() method, rectangle also satisfies 'measureable' interface
// Since it has BOTH methods, it satisfies 'geometry' interface too!
func (r rectangle) perimeter() float32 {
	return 2 * (r.width + r.height)  // Rectangle's way of calculating perimeter
}

// ⭕ CIRCLE TYPE: Different implementation, same interface
type circle struct {
	radius float32  // Different data structure
}

// 🎯 SAME INTERFACE, DIFFERENT IMPLEMENTATION:
// Circle also has area() method, so it ALSO satisfies 'shape' interface
func (c circle) area() float32 {
	return math.Pi* c.radius* c.radius  // Circle's way of calculating area
}

// Circle also has perimeter() method, so it satisfies 'measureable' too
// Both rectangle and circle can be used wherever 'geometry' interface is expected!
func (c circle) perimeter() float32 {
	return 2 * math.Pi * c.radius  // Circle's way of calculating perimeter
}

// 🚀 POLYMORPHIC FUNCTION: Works with ANY type that satisfies 'shape' interface
// This function doesn't care if it's rectangle, circle, or triangle - 
// as long as it has area() method!
func calculateArea(s shape) float32 {
	return s.area()  // Calls the appropriate area() method based on actual type
}

// 🎯 INTERFACE COMPOSITION IN ACTION: Function accepts any type with BOTH methods
// Only types that satisfy BOTH shape AND measureable can be passed here
func decribeShape(g geometery){
	fmt.Println("Area : ", g.area())        // Uses shape interface method
	fmt.Println("Perimerter : ", g.perimeter()) // Uses measureable interface method
}

func main() {
	// 🏗️ CREATE CONCRETE TYPES
	rect := rectangle{
		width:  20,
		height: 30,
	}

	cir := circle{
		radius: 3,
	}
    
	// 🎯 POLYMORPHISM IN ACTION:
	// Same function works with different types!
	// fmt.Println("Area of rectangle : ", calculateArea(rect))  // rect satisfies shape
	// fmt.Println("Area of circle : ", calculateArea(cir))     // cir also satisfies shape

	// 🔗 INTERFACE COMPOSITION:
	// Both rect and cir satisfy geometry interface (have both area() and perimeter())
	 decribeShape(rect)  // Works because rectangle has both methods
	 decribeShape(cir)   // Works because circle has both methods

	// 🎉 MAGIC: Same function call, different behaviors based on actual type!
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

🔍 INTERFACE SATISFACTION (IMPLICIT):
• No "implements" keyword needed
• If type has all required methods → automatically satisfies interface
• Duck typing: "If it walks like a duck and quacks like a duck, it's a duck"

📊 INTERFACE TYPES:
┌─────────────────┬─────────────────┬─────────────────┐
│   Interface     │    Methods      │   Satisfied By  │
├─────────────────┼─────────────────┼─────────────────┤
│ shape           │ area()          │ rectangle, circle│
│ measureable     │ perimeter()     │ rectangle, circle│
│ geometry        │ area()+perimeter│ rectangle, circle│
│ interface{}     │ (none)          │ ANY type        │
└─────────────────┴─────────────────┴─────────────────┘

🎯 INTERFACE BENEFITS:
✅ Polymorphism: One function, many types
✅ Testability: Easy to mock interfaces
✅ Flexibility: Add new types without changing existing code
✅ Loose coupling: Depend on behavior, not implementation

💡 COMMON PATTERNS:
• io.Reader, io.Writer (standard library)
• error interface (built-in)
• Stringer interface (fmt package)
• Handler interface (http package)

🚨 GOTCHAS:
❌ Interface{} loses type safety
❌ Nil interface vs nil pointer confusion
❌ Interface values hold both type and value
❌ Method sets: pointer vs value receivers

🔧 BEST PRACTICES:
• Keep interfaces small (1-3 methods)
• Define interfaces where you USE them, not where you implement them
• Accept interfaces, return concrete types
• Use composition over large interfaces

=============================================================================
*/