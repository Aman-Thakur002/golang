
/* 
An interface is a contract that defines a set of methods (and sometimes properties) that a type must implement. It does not contain implementation, only the method signatures.

Interfaces define behavior, not data.

Any type that implements all methods of an interface implicitly satisfies that interface (in Go).

Interfaces promote loose coupling and make code more flexible.

General Definition
An interface is an abstract type that specifies what methods a type must have, without specifying how they are implemented.

*/


package main

import (
	"fmt"
	"math"
)

type shape interface { 
	area() float32
}

type measureable interface {
	perimeter() float32
}

type geometery interface {    // embedding multiple interfaces -> interface composition
	shape
	measureable
}

// ---------rectabgle-------
type rectangle struct {
	width, height float32
}

func (r rectangle) area() float32 {
	return r.width * r.height
}

func (r rectangle) perimeter() float32 {
	return 2 * (r.width + r.height)
}

// ---------circle-----------
type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return math.Pi* c.radius* c.radius
}

func (c circle) perimeter() float32 {
	return 2 * math.Pi * c.radius
}


//--------- calculate area function which implements interface --------------
func calculateArea(s shape) float32 {
	return s.area()
}

//------------ decirbe shape function which is implementing interface composition ---------------------
func decribeShape(g geometery){
	fmt.Println("Area : ", g.area())
	fmt.Println("Perimerter : ", g.perimeter())
}


func main() {

	rect := rectangle{
		width:  20,
		height: 30,
	}

	cir := circle{
		radius: 3,
	}
    
	// fmt.Println("Area of rectangle : ", calculateArea(rect))
	// fmt.Println("Area of rectangle : ", calculateArea(cir))

	 decribeShape(rect)
	 decribeShape(cir)

}
