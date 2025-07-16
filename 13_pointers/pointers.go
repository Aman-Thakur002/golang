package main

import "fmt"


// pass by value , so num passed in this function is just a copy 
func changeNum(num int) {
	num = 5
	fmt.Println("In chnage num", num)
}

//pass by refernece 
func fun1(num *int){
   *num = 5  // dereference
}

func main() {
  num := 1
  changeNum(num)
  fmt.Println("After change num ", num)

  fmt.Println("Memory address", &num) 
  fun1(&num)
  fmt.Println("After change num (pass by reference)", num)

}