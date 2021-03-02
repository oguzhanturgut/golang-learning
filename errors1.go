package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	if file, err := os.Open("non-existing.txt"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(file.Name() + "opened successfully")
	}

	sampleErr := errors.New("error occurred")
	fmt.Println(sampleErr)

	sampleErr = fmt.Errorf("Err is: %s", "database connection issue")
	fmt.Println(sampleErr)
}
