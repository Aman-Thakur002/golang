package main

import (
	"fmt"
	"slices"
)

// no need to tell the size
// -dynamic array
// - most used construct in go


func main(){
// uninitialised slice is nil	
 var nums []int
 fmt.Println("Slice",nums)
 fmt.Println(nums == nil)


 // if you don't want nil slice then 
  nums1 := make([]int,0,5)  // type, initial size, capacity
  nums1 = append(nums1, 1)
  nums1 = append(nums1, 2) //to add element in the last
 fmt.Println("Not null slice 1",nums1)
 fmt.Println("Length : ",len(nums1))
 fmt.Println("capacity : ",cap(nums1))



 //other way to initalise not nil slice
 nums2 := []int{}
 nums2 = append(nums2, 4)
 fmt.Println("Not null slice 2",nums2)

 //copy function
 copiedNums := make([]int, len(nums2))
 copy(copiedNums, nums2)
 fmt.Println("Copied Slice",copiedNums)


 //slice operator --> returns the subarray from the array (start, end]
 nums3 := [3]int{2,3,4}
 fmt.Println("Sliced array of nums3",nums3[0:2])
 fmt.Println("Sliced array of nums3",nums3[:3])
 fmt.Println("Sliced array of nums3",nums3[0:])


 //slices package
 s1 := []int{1,2}
 s2 := []int{1,2}
 s3 := []int{1,3}
 fmt.Println(slices.Equal(s1,s2))
 fmt.Println(slices.Equal(s1,s3))

 // - 2d slices
 array2D := [][]int{{1,2,4},{3,4}}
 fmt.Println(array2D)
 
}