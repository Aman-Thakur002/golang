package main

import "fmt"

func main() {

	//in while loop fashion
	k := 1
	for k <= 3 {
		fmt.Println(k)
		k = k + 1
	}

	// in for-loop fashion
	for i:= 0 ; i<3; i++ {
		fmt.Println(i)
	}

	//modern way for range based looping
	for y:= range 3 {
		fmt.Println(y)
	}
}