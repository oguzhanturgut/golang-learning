package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee3 struct {
	Name   string `json:"n"`
	Age    int    `json:"a"`
	salary int    `json:"s"`
}

type employee4 struct {
	Name   string
	Age    int
	salary int
}

func main() {
	e1Json := `{"n":"John","a":21}`

	var e1Converted employee3
	err := json.Unmarshal([]byte(e1Json), &e1Converted)
	if err != nil {
		log.Fatalf("Error occured during unmarshaling. Error: %s", err.Error())
	}
	fmt.Printf("employee3 Struct: %#v\n", e1Converted)

	e2Json := `{"Name":"John","Age":21}`
	var e2Converted employee4
	err = json.Unmarshal([]byte(e2Json), &e2Converted)
	if err != nil {
		log.Fatalf("Error occured during unmarshaling. Error: %s", err.Error())
	}
	fmt.Printf("\nemployee2 Struct: %#v\n", e2Converted)
}
