package main

import "fmt"

func main() {
	//a := new(int)
	//*a = 10
	//// Use * to dereference a pointer and get its value
	//fmt.Println(*a) //Print the value stored at address a

	//b := 2
	//c := &b
	//fmt.Println(*c) //Output will be 2

	//Declare
	//var y *int
	//x := 2
	//y = &x

	//Will print a address. Output will be different everytime.
	//fmt.Println(y)
	//fmt.Println(*y)
	//y = new(int)
	//*y = 10
	//fmt.Println(*y)

	// Dereference a pointer which means getting the value at address stored in the pointer.

	//a := 2
	//b := &a
	//fmt.Println(a)  // 2
	//fmt.Println(*b) // 2
	//
	//*b = 3
	//fmt.Println(a)  // 3
	//fmt.Println(*b) // 3
	//
	//a = 4
	//fmt.Println(a)  // 4
	//fmt.Println(*b) // 4

	// Pointer to a pointer
	//a := 2
	//b := &a
	//c := &b
	//
	//fmt.Printf("a: %d\n", a)
	//fmt.Printf("b: %x\n", b)
	//fmt.Printf("c: %x\n", c)
	//
	//fmt.Println()
	//fmt.Printf("a: %d\n", a)
	//fmt.Printf("*&a: %d\n", *&a)
	//fmt.Printf("*b: %d\n", *b)
	//fmt.Printf("**c: %d\n", **c)
	//
	//fmt.Println()
	//fmt.Printf("&a: %d\n", &a)
	//fmt.Printf("b: %d\n", b)
	//fmt.Printf("&*b: %d\n", &*b)
	//fmt.Printf("*&b: %d\n", *b)
	//fmt.Printf("*c: %d\n", *c)
	//
	//fmt.Println()
	//fmt.Printf("b: %d\n", &b)
	//fmt.Printf("*c: %d\n", c)

	//* Default zero value of a pointer is nil
	var a *int
	fmt.Print("Default Zero Value of a pointer: ")
	fmt.Println(a)
}
