package main

import (
	"fmt"
	"maps"
)

func main() {
	// Creting map

	m := make(map[string]int)
	// Adding key-value pairs to the map
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3

	fmt.Println(m)

	// Getting value from the map
	fmt.Println(m["a"])

	// When key is not present in the map, it returns the zero value of the value type
	fmt.Println(m["d"]) // prints 0

	// Length of the map
	fmt.Println(len(m)) // prints 3

	// Deleting a key-value pair from the map
	delete(m, "b")
	fmt.Println(m) // prints map[a:1 c:3]

	// Checking if a key exists in the map
	value, exists := m["b"]
	if exists {
		fmt.Println("Key 'b' exists with value:", value)
	} else {
		fmt.Println("Key 'b' does not exist.")
	}

	// Creating a map with initial values
	m2 := map[string]int{
		"x": 10,
		"y": 20,
		"z": 30,
	}
	fmt.Println(m2)

	// Iterating over a map
	for key, value := range m2 {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	// Comparing maps
	m3 := map[string]int{
		"x": 10,
		"y": 20,
		"z": 30,
	}
	//fmt.Printf("%v\n", m2 == m3) //Compile error: invalid operation: m2 == m3 (map can only be compared to nil)

	fmt.Printf("%v\n", maps.Equal(m2, m3)) // prints true

}
