package main

// func (receiver receiver_type) some_func_name(arguments) return_values

import "fmt"

type employeeNew struct {
	name   string
	age    int
	salary int
}

func (e employeeNew) details() {
	fmt.Printf("Name: %s\n", e.name)
	fmt.Printf("Age: %d\n", e.age)
}

func (e *employeeNew) setNewName(newName string) {
	e.name = newName
}

func (e employeeNew) getSalary() int {
	return e.salary
}

func main() {
	emp := employeeNew{name: "Sam", age: 31, salary: 2000}
	emp.details()
	emp.setNewName("Dave")
	emp.details()
	fmt.Printf("Name: %s\n", emp.name)
	(&emp).setNewName("Mike")
	emp.details()
	//fmt.Printf("Salary %d\n", emp.getSalary())
}
