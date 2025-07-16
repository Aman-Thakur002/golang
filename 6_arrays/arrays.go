package main

import "fmt"

func main() {
	var nums [4]int
	nums[0] = 1
	fmt.Println(nums[0])
	fmt.Println(nums)
	fmt.Println(len(nums))


	var boolArray [4]bool
	fmt.Println(boolArray)

	var names [3]string
	fmt.Println(names)
 
    // array declaration 
	numsArray := [3]int{1,2,4}
	fmt.Println(numsArray)

	//2d array
	num2dArray := [2][2]int{{1,2},{3,4}}
	fmt.Println(num2dArray)


}