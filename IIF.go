package main

import "fmt"

func main() {
	// Immediately invoked function
	squareOf2 := func() int {
		return 2 * 2
	}()
	fmt.Println(squareOf2)
}
