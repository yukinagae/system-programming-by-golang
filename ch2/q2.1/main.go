package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(file, "%d\n", 1)
	fmt.Fprintf(file, "%s\n", "hello")
	fmt.Fprintf(file, "%f\n", 2.3)
	file.Close()
}
