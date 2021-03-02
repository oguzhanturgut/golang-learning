package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//GO struct is named collection of data fields which can be of different types.
//Struct acts as a container that has different heterogeneous data types which
//together represents an entity.

type employee struct {
	name   string
	age    int
	salary int
}

type point struct {
	x float64
	y float64
}

func main() {
	emp1 := employee{}
	fmt.Printf("Emp1: %+v\n", emp1)

	emp2 := employee{name: "Sam", age: 31, salary: 2000}
	fmt.Printf("Emp2: %+v\n", emp2)

	emp3 := employee{
		name:   "Sam",
		age:    31,
		salary: 2000,
	}
	fmt.Printf("Emp3: %+v\n", emp3)

	emp4 := employee{
		name: "Sam",
		age:  31,
	}
	fmt.Printf("Emp4: %+v\n", emp4)

	emp := employee{name: "Sam", age: 31, salary: 2000}

	//Accessing a struct field
	n := emp.name
	fmt.Printf("Current name is: %s\n", n)

	//Assigning a new value
	emp.name = "John"
	fmt.Printf("New name is: %s\n", emp.name)

	// Pointer to a struct
	empP := &emp
	fmt.Printf("Emp: %+v\n", empP)
	empP = &employee{name: "John", age: 30, salary: 3000}
	fmt.Printf("Emp: %+v\n", empP)

	//Using the  new() keyword will:
	//	Create the struct
	//	Initialize all the field to the zero default value of their type
	//	Return the pointer to the newly created struct
	empP2 := new(employee)
	fmt.Printf("Emp Pointer Address: %p\n", empP2)
	fmt.Printf("Emp Pointer: %p\n", empP2)
	fmt.Printf("Emp Value: %+v\n", *empP2)

	empJSON, err := json.MarshalIndent(emp, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(string(empJSON))
}
