package main

import "fmt"

func main() {
	// empty slice declaration
	var sample []int
	fmt.Println(len(sample))
	fmt.Println(cap(sample))
	fmt.Println(sample)

	letters := []string{"a", "b", "c"}
	fmt.Println(len(letters))
	fmt.Println(cap(letters))
	fmt.Println(letters)

	// re-slicing an array
	numbers := [5]int{1, 2, 3, 4, 5}

	//Both start and end
	num1 := numbers[2:4]
	fmt.Println("Both start and end")
	fmt.Printf("num1=%v\n", num1)
	fmt.Printf("length=%d\n", len(num1))
	fmt.Printf("capacity=%d\n", cap(num1))

	//Only start
	num2 := numbers[2:]
	fmt.Println("\nOnly start")
	fmt.Printf("num1=%v\n", num2)
	fmt.Printf("length=%d\n", len(num2))
	fmt.Printf("capacity=%d\n", cap(num2))

	//Only end
	num3 := numbers[:3]
	fmt.Println("\nOnly end")
	fmt.Printf("num1=%v\n", num3)
	fmt.Printf("length=%d\n", len(num3))
	fmt.Printf("capacity=%d\n", cap(num3))

	//None
	num4 := numbers[:]
	fmt.Println("\nOnly end")
	fmt.Printf("num1=%v\n", num4)
	fmt.Printf("length=%d\n", len(num4))
	fmt.Printf("capacity=%d\n", cap(num4))

	numbers[3] = 8
	// each of the new slice is still referring to the original array
	fmt.Printf("num1=%v\n", num2)
	fmt.Printf("num3=%v\n", num3)
	fmt.Printf("num4=%v\n", num4)

	// create slice with make()
	numbers2 := make([]int, 3, 5)
	fmt.Println()
	fmt.Printf("numbers=%v\n", numbers2)
	fmt.Printf("length=%d\n", len(numbers2))
	fmt.Printf("capacity=%d\n", cap(numbers2))

	//With capacity ommited
	numbers2 = make([]int, 3)
	fmt.Println("\nCapacity Ommited")
	fmt.Printf("numbers=%v\n", numbers2)
	fmt.Printf("length=%d\n", len(numbers2))
	fmt.Printf("capacity=%d\n", cap(numbers2))

	// Increasing and decreasing the length of slice
	//numbers := make([]int, 3, 5)
	//fmt.Printf("numbers=%v\n", numbers)
	//fmt.Printf("length=%d\n", len(numbers))
	//fmt.Printf("capacity=%d\n", cap(numbers))
	//
	////This line will cause a runtime error index out of range [4] with length 3
	////numbers[4] = 5
	//
	////Increasing the length from 3 to 5
	//numbers = numbers[0:5]
	//fmt.Println("\nIncreasing length from 3 to 5")
	//fmt.Printf("numbers=%v\n", numbers)
	//fmt.Printf("length=%d\n", len(numbers))
	//fmt.Printf("capacity=%d\n", cap(numbers))
	//
	////Decresing the length from 3 to 2
	//numbers = numbers[0:2]
	//fmt.Println("\nDecreasing length from 3 to 2")
	//fmt.Printf("numbers=%v\n", numbers)
	//fmt.Printf("length=%d\n", len(numbers))
	//fmt.Printf("capacity=%d\n", cap(numbers))

	array := [5]int{1, 2, 3, 4, 5}
	// creates a new array
	array2 := array
	// pointer to original array
	slice := array[:]
	slice[0] = 99
	fmt.Printf("Array: %v \n", array2)
	fmt.Printf("Slice: %v \n", slice)

	// Appending values to slice
	numbers3 := []int{1, 2}
	//This function in the background increases the length and capacity of the slice if needed
	numbers3 = append(numbers3, 3, 4, 5) //Slice will become [1, 2, 3, 4, 5]
	fmt.Printf("New Array: %v\n", numbers3)

	// Spreading a slice to another slice
	numbers4 := []int{1, 2}
	numbers5 := []int{3, 4}
	numbers4 = append(numbers4, numbers5...)
	fmt.Printf("numbers=%v\n", numbers4)
	fmt.Printf("length=%d\n", len(numbers4))
	fmt.Printf("capacity=%d\n", cap(numbers4))

	// Copying slices
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, 5)

	numberOfElementsCopied := copy(dst, src)
	fmt.Printf("Number Of Elements Copied: %d\n", numberOfElementsCopied)
	fmt.Printf("dst: %v\n", dst)
	fmt.Printf("src: %v\n", src)

	//After changing numbers2
	dst[0] = 10
	fmt.Println("\nAfter changing dst")
	fmt.Printf("dst: %v\n", dst)
	fmt.Printf("src: %v\n", src)

	var numbers6 []int
	fmt.Println()
	fmt.Printf("numbers6=%v\n", numbers6)
	fmt.Printf("length=%d\n", len(numbers6))
	fmt.Printf("capacity=%d\n", cap(numbers6))
	numbers6 = append(numbers6, 1)
	fmt.Printf("numbers6=%v\n", numbers6)
	fmt.Printf("length=%d\n", len(numbers6))
	fmt.Printf("capacity=%d\n", cap(numbers6))

	// multi-dimensional slices
	twoDSlice1 := make([][]int, 3)
	for i := range twoDSlice1 {
		twoDSlice1[i] = make([]int, 3)
	}
	fmt.Printf("Number of rows in slice: %d\n", len(twoDSlice1))
	fmt.Printf("Number of columns in arsliceray: %d\n", len(twoDSlice1[0]))
	fmt.Printf("Total number of elements in slice: %d\n", len(twoDSlice1)*len(twoDSlice1[0]))
	fmt.Println("First Slice")
	for _, row := range twoDSlice1 {
		for _, val := range row {
			fmt.Println(val)
		}
	}
	twoDSlice2 := make([][]int, 2)
	twoDSlice2[0] = []int{1, 2, 3}
	twoDSlice2[1] = []int{4, 5, 6}
	fmt.Println()
	fmt.Printf("Number of rows in slice: %d\n", len(twoDSlice2))
	fmt.Printf("Number of columns in arsliceray: %d\n", len(twoDSlice2[0]))
	fmt.Printf("Total number of elements in slice: %d\n", len(twoDSlice2)*len(twoDSlice2[0]))
	fmt.Println("Second Slice")
	for _, row := range twoDSlice2 {
		for _, val := range row {
			fmt.Println(val)
		}
	}

}
