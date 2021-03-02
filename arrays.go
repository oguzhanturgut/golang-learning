package main

import "fmt"

// sample := [size_of_array]{type}{a1, a2... an}
//	size_of_array – number of elements in the array
//	<type> is type of each element in the array
//	a1, a2 … an are the actual elements.

func main() {
	////Both number of elements and actual elements
	//sample1 := [2]int{1, 2}
	//fmt.Printf("Sample1: Len: %d, %v\n", len(sample1), sample1)
	//
	////Only actual elements
	//sample2 := [...]int{2, 3}
	//fmt.Printf("Sample2: Len: %d, %v\n", len(sample2), sample2)
	//
	////Only number of elements
	//sample3 := [2]int{}
	//fmt.Printf("Sample3: Len: %d, %v\n", len(sample3), sample3)
	//
	////Without both number of elements and actual elements
	//sample4 := [...]int{1, 2, 3, 4, 5, 6}
	//fmt.Printf("Sample4: Len: %d, %v\n", len(sample4), sample4)

	// Accessing and manipulating array values
	//sample1 := [2]string{"a", "b"}
	//fmt.Printf("Sample1 Before: %v\n", sample1)
	//sample2 := sample1
	//sample2[0] = "c"
	//fmt.Printf("Sample1 After assignment: %v\n", sample1)
	//fmt.Printf("Sample2: %v\n", sample2)
	//test(sample1)
	//fmt.Printf("Sample1 After Test Function Call: %v\n", sample1)

	// Multi dimensional arrays
	sample := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("First Run")
	for _, row := range sample {
		for _, val := range row {
			fmt.Println(val)
		}
	}

	sample[0][0] = 6
	sample[1][2] = 1
	fmt.Println("\nSecond Run")
	for _, row := range sample {
		for _, val := range row {
			fmt.Println(val)
		}
	}
}

//func test(sample [2]string) {
//	sample[0] = "d"
//	fmt.Printf("Sample in Test function: %v\n", sample)
//}
