package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// infinite look
	/*for {
		fmt.Println("infinite loop")
	}*/

	// classic for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// range loop from 1.22
	for i := range 3 {
		fmt.Println(i)
	}

}
