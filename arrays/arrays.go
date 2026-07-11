package main

import "fmt"

func main() {
	var nums [4]int
	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	nums[3] = 4

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	var vals [4]bool
	vals[0] = true
	fmt.Println(vals)

	var names [3]string
	names[0] = "Alice"
	fmt.Println(names[0])

	// arrays can also be initialized with values
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)

	// 2d arrays can be created by specifying the size of each dimension
	var matrix [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	nums2 := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(matrix)
	fmt.Println(nums2)
}
