package main

import (
	"bufio"
	"fmt"
	"io"
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

	// Close it now so the later copy step can recreate/write the same file safely.
	if err := f1.Close(); err != nil {
		fmt.Println("Error closing file:", err)
		return
	}

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

	// copy the content of example.txt to example1.txt
	srcFile, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}

	defer srcFile.Close()

	destFile, err := os.Create("example1.txt")
	if err != nil {
		fmt.Println("Error creating destination file:", err)
		return
	}

	reader := bufio.NewReader(srcFile)
	writer := bufio.NewWriter(destFile)

	for {
		line, err := reader.ReadBytes('\n')
		if len(line) > 0 {
			_, werr := writer.Write(line)
			if werr != nil {
				fmt.Println("Error writing line:", werr)
				return
			}
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading line:", err)
			return
		}
	}
	if err := writer.Flush(); err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}
	fmt.Println("Written to new file successfully")

	if err := destFile.Close(); err != nil {
		fmt.Println("Error closing destination file:", err)
		return
	}

	// delete a file
	err = os.Remove("example1.txt")
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}
	fmt.Println("File deleted successfully.")

}
