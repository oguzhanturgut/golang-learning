package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Printf("%U\n", []rune("0b£"))

	rPound := '£'
	fmt.Printf("Type: %s\n", reflect.TypeOf(rPound))
	fmt.Printf("Unicode CodePoint: %U\n", rPound)
	fmt.Printf("Character: %c\n", rPound)

	s := "ab£"
	fmt.Println([]byte(s))

	for _, c := range s {
		fmt.Println(string(c))
	}
}
