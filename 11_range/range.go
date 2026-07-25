package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}

	fmt.Println("Display using for loop:")
	for i := 0; i < len(nums); i++ {
		fmt.Printf("%d\n", nums[i])
	}

	fmt.Println("Display using range1.")
	for _, num := range nums {
		fmt.Printf("%d\n", num)
	}

	fmt.Println("Display using range2.")
	for i, num := range nums {
		fmt.Printf("Index: %d, Value: %d\n", i, num)
	}

	// Iterate over a map using range
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println("Display map using range.")
	for key := range m {
		fmt.Printf("Key: %s\n", key)
	}

	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	for i, c := range "hello" {
		fmt.Printf("Index: %d, Character: %c\n", i, c)
	}

	fmt.Println("Dispaly whole mal")
}
