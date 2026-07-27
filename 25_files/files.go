package main

import (
	"fmt"
	"os"
)

// main demonstrates basic file and directory operations using the os package.
// It expects example.txt to exist in the current working directory.
func main() {
	// 1) Create a new file.
	f1, err := os.Create("example1.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	// 2) Open an existing file for reading.
	f2, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// 3) Read metadata such as name, size, permissions, and modification time.
	fileInfo, err := f2.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}

	fmt.Println("File Name:", fileInfo.Name())
	fmt.Println("Size:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last Modified:", fileInfo.ModTime())

	// Close the created file when the function exits.
	defer f1.Close()

	// 4) Read content manually using a byte buffer and file.Read.
	buf := make([]byte, fileInfo.Size())
	data, err := f2.Read(buf)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("File Content:", string(buf[:data]))

	// 5) Read the entire file at once with os.ReadFile.
	f3, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File Content using ReadFile:", string(f3))

	// 6) Open the current directory and list all entries.
	dir, err := os.Open(".")
	if err != nil {
		fmt.Println("Error opening directory:", err)
		return
	}
	defer dir.Close()
	infos, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}
	for _, info := range infos {
		fmt.Println("Name:", info.Name(), "IsDir:", info.IsDir())
	}

}
