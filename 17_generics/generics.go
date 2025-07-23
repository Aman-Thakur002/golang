/*
 -> A generic function declares one or more type parameters in square brackets immediately after the function name. Each type parameter can have an optional constraint (an interface) that restricts which concrete types are allowed.

 Map is a generic function with two type parameters: T and U.

The constraint any means “no restrictions” (any type is allowed).

At call‑site you can use it for, say, []int → []string or []float64 → []float64.

 -> Benefits
DRY Code: write algorithms once instead of duplicating for each type.

Type Safety: catches mismatches at compile time, unlike interface{} + type assertions.

Expressiveness: constraints let you specify exactly what capabilities (“≤”, “+”, methods, etc.) you need.

-> When to Use
Collections–style utilities (Map, Filter, Reduce).

Data structures (linked lists, trees, caches) that work over many element types.

Algorithms (search, sort, set) that apply across types.

Anywhere you’d otherwise accept interface{} and cast internally.

Generics let you write clean, reusable, and type‑safe code without the boilerplate of duplicating logic for each type. Once you’ve internalized the syntax ([type parameters] + [constraints]), you can start refactoring existing code to share more logic generically!

*/

package main

import "fmt"

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
// func printSlice[T int | string](items []T) {  // -> now we are allowing integer and string 
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

//-> we use comparable too
func printSlice[T comparable](items []T) {  
	for _, item := range items {
		fmt.Println(item)
	}
}

//-----------generic struct --------------

type stack[T any] struct {
	elements []T
}

func main() {
    
	myStack := stack[string] {
		elements: []string{"Java"},
	}
     fmt.Println(myStack)

	nums := []int{1, 3, 9}
	stringsSlice := []string{"Golang", "Javascript"}
	printSlice(nums)
	printSlice(stringsSlice)
}
