package main

import "fmt"
  // functions which can accept any number of parameters

  func sum(nums ...int) int {
	total := 0
	for _, num := range nums{
		total = total + num
	}

	return total
  }

// to receive anytype of parameters
func fun1(para ...interface{}) interface{}{  // or use "any" keyword
	fmt.Print("Mixed values: ")
	for i, item := range para {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%v", item)
	}
	fmt.Println()
	return len(para) // return count of parameters
}


func main() {
	nums := []int{2,4,23,45,13,2,23}
	result := sum(nums ...)
	fmt.Println(result)

	// Test fun1 with mixed types
	count := fun1(1, "hello", 3.14, true, []int{1,2,3})
	fmt.Printf("Total parameters: %v\n", count)
}