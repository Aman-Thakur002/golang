package main

import "fmt"
 // func add(a, b int) this means all parameters are f int type

func add(a int, b int) int { // (paramters), return value
	return a + b
}

// to return multiple values
func getLanguages()(string, string, bool){

	return "golang","javascript", true
}

//function passed in a  function
func processIt(f1 func(a int) int){
   
}

//function returning a function
func func2() func(a int) int {
	return func(a int) int {
		return 4
	}
}

func main() {
	result := add(2, 4)
	fmt.Println(result)
	l1, l2, l3 := getLanguages()
	fmt.Println(l1, l2, l3)

}