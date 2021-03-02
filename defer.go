package main

//func main() {
//	defer test2()
//	defer func() { fmt.Println("In inline defer") }()
//	fmt.Println("Executed in main")
//}
//func test2() {
//	fmt.Println("In Defer")
//}

//import "fmt"
//
//func main() {
//	defer fmt.Println("Defer in main")
//	fmt.Println("Stat main")
//	f1()
//	fmt.Println("Finish main")
//}
//
//func f1() {
//	defer fmt.Println("Defer in f1")
//	fmt.Println("Start f1")
//	f2()
//	fmt.Println("Finish f1")
//}
//
//func f2() {
//	defer fmt.Println("Defer in f2")
//	fmt.Println("Start f2")
//	fmt.Println("Finish f2")
//}

import "fmt"

func main() {
	i := 0
	i = 1
	defer fmt.Println(i)
	i = 2
	defer fmt.Println(i)
	i = 3
	defer fmt.Println(i)
}
