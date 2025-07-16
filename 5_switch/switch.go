package main

import (
	"fmt"
	"time"
)

func main() {

	// simple switch, no need to provide break keyword
	i := 8
	switch i {
	case 1:
		fmt.Println("one")
	case 2 :
		fmt.Println("Two")
	case 3 :
		fmt.Println("Three")
	case 4 :
		fmt.Println("four")
	case 5 :
		fmt.Println("five")
		default : 
		fmt.Println("Number is greater than five")
	}

	// multiple condition switch
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday :
		fmt.Println("It's weekend")

		default : 
		fmt.Println("It's word day")
	}


	//type switch
	whoAmI := func(i interface{}){
		switch i.(type){
		case int : 
		fmt.Println("Integer")
		case bool :
			fmt.Println("Boolean")
		case string :
			fmt.Println("String")
		default :
		fmt.Println("Other")
		}
	}

	whoAmI(1)

}