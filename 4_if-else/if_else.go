package main

import "fmt"

func main() {
	age := 17

	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}

	//else if same as other language like cpp , js, java


	// we can delcare variable directly inside if construct and that varibale will be block scoped
	if age := 11; age>= 18{
		fmt.Println("Adult")
	} else if age >=12 {
		fmt.Println("Age : ", age)
	} else {
		fmt.Println("Minor alert")
	}

	// go doesn't have ternary operator

}