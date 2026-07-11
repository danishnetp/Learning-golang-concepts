package main

import (
	"fmt"
	"slices"
)

func main() {
	var nums []int
	fmt.Println(nums)      // prints an empty slice
	fmt.Println(len(nums)) // prints 0

	var nums2 = make([]int, 2)
	fmt.Println(nums2)      // prints a slice with 2 zeroed elements
	fmt.Println(len(nums2)) // prints 2
	fmt.Println(cap(nums2)) // prints 2

	var nums3 = make([]int, 2, 5)
	fmt.Println("nums3 : ", nums3) // prints a slice with 2 zeroed elements
	fmt.Println(len(nums3))        // prints 2
	fmt.Println(cap(nums3))
	nums3 = append(nums3, 1, 2, 3) // nums3 now has 5 elements
	nums3 = append(nums3, 4)       // nums3 now has 6 elements, capacity will increase
	fmt.Println("nums3 after append: ", nums3)
	fmt.Println(len(nums3)) // prints 6
	fmt.Println(cap(nums3)) // prints 10

	nums4 := []int{}
	fmt.Println("nums4 : ", nums4) // prints an empty slice

	// copying slices
	nums5 := []int{1, 2, 3}
	nums6 := make([]int, len(nums5))
	copy(nums6, nums5)
	fmt.Println("nums5 : ", nums5) // prints [1 2 3]
	fmt.Println("nums6 : ", nums6) // prints [1 2 3]

	// slice package
	var nums7 = []int{1, 2, 3, 4, 5}
	var nums8 = []int{1, 2, 3, 4, 5}
	fmt.Println(slices.Equal(nums7, nums8)) // prints true

	// slice operators
	nums9 := []int{1, 2, 3, 4, 5}
	slice1 := nums9[1:4]
	fmt.Println("slice1 : ", slice1) // prints [2 3 4]
	fmt.Println(nums9[1:])           // prints [2 3 4 5]
	fmt.Println(nums9[:1])

}
