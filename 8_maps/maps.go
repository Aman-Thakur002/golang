package main

import (
	"fmt"
	"maps"
)

// maps -> hash, object, dictionary

func main(){

	//creating map
	m := make(map[string]string)

	// setting elements
	m["name"] = "aman"
	m["age"] = "23"
  
	fmt.Println(m)
	fmt.Println(m["name"])  // IMP : if key doesn't exist in map , it will return zero, empty string, false

	// to delete an element
	delete(m,"name")
	fmt.Println(m)

	//to clear a map
	clear(m)
	fmt.Println(m)

	// map without make function
	 m1 := map[string]int{"price" : 30, "phone" : 3}
	 fmt.Println(m1)

	 // to check an element an map
	value, ok := m1["price"]  // map return multiple values and "ok" is an idiom to check for an idiom
	fmt.Println(value)
	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}


    // maps package just like slices package
	m2 := map[string]int{"price" : 20, "phones" : 2}
	m3 := map[string]int{"price" : 20, "phones" : 3}
	fmt.Println(maps.Equal(m2,m3))
}