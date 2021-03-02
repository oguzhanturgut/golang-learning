package main

import (
	"fmt"
	"time"
)

func main() {
	//var a chan int
	//a = make(chan int)
	//fmt.Println(a)

	// Send
	//ch <- val

	// Receive
	//val := <- ch

	ch := make(chan int, 1024)

	fmt.Println("Sending value to channel")
	go send(ch, 42)

	fmt.Println("Receiving from channel")
	go receive(ch)

	time.Sleep(time.Second * 1)
}

func send(ch chan int, val int) {
	ch <- 10
	ch <- 20
	ch <- val
}

func receive(ch chan int) {
	//val := <-ch
	for val := range ch {
		fmt.Printf("Value Received=%d in receive function\n", val)
	}
}
