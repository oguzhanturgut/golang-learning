package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		if i%3 == 0 {
			continue
		}
		fmt.Println(i)
	}

	i := 1
	//Function call in the init part in for loop
	for test2(); i < 3; i++ {
		fmt.Println(i)
	}

	//Assignment in the init part in for loop
	for i = 2; i < 3; i++ {
		fmt.Println(i)
	}

	// While loop implementation
	j := 1
	for j <= 5 {
		fmt.Println(j)
		j++
	}

}

func test2() {
	fmt.Println("In test function")
}
