package main

//	IOTA
// A counter which starts with zero
// Increases by 1 after each line
// Is only used with constant

import "fmt"

const (
	a = iota
	b
	c
)

//const (
//	a = iota
//	b = iota
//	c = iota
//)

// Iota value will reset and again start with zero if the const keyword is used again
//const (
//	a = iota
//	b
//)
//const (
//	c = iota
//)

// iota increment can be skipped using a blank identifier
//const (
//	a = iota
//	_
//	b
//	c
//)

//iota allows expressions which can be used to set any value for the constant
//const (
//	a = iota
//	b = iota + 4
//	c = iota * 4
//)

type Size int

// Enum
const (
	small = Size(iota)
	medium
	large
	extraLarge
)

func main() {
	fmt.Println(a, b, c)

	fmt.Println(small)
	fmt.Println(medium)
	fmt.Println(large)
	fmt.Println(extraLarge)
}

func (s Size) toString() {
	switch s {
	case small:
		fmt.Println("Small")
	case medium:
		fmt.Println("Medium")
	case large:
		fmt.Println("Large")
	case extraLarge:
		fmt.Println("Extra Large")
	default:
		fmt.Println("Invalid Size entry")

	}
}
