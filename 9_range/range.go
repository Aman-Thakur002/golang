package main

import "fmt"

// use for iteration over data structures

func main() {
	nums := []int{1, 3, 5}

	// nor iteration with for loop
	//    for i := 0; i<len(nums) ; i++ {
	// 	fmt.Println(nums[i])
	//    }

	// using range
	for i, num := range nums { // index, value
		fmt.Println("Index : ", i, " ", "Value : ", num)
	}

	// iteration over maps
	m := map[string]string{"name": "aman", "occupation": "backend engineer"}
	for k, v := range m { // key, value
		fmt.Println(k, v)
	}

	// iteration over string
	// c is the unicode of every character , e.g for A unicode is 65
	// unicode point rune
	// if unicode <=255 -> 1 byte, if unicode is bigger then 255 then it takes more than 1 byte so it changes the index of other character, e.g if string "AM", unicode 0f A is 300 , i =0 then index of M i.e i would be 2 if A is taking 2 bytes
	for i, c := range "Aman Pratap" {  // i is statrting byte of index of rune
		fmt.Println(i, c,string(c))
	}

}
