package main

import "fmt"

func main() {
	//letters := []string{"a", "b", "c"}
	//
	////With index and value
	//fmt.Println("Both Index and Value")
	//for i, letter := range letters {
	//	fmt.Printf("Index: %d Value:%s\n", i, letter)
	//}
	//
	////Only value
	//fmt.Println("\nOnly value")
	//for _, letter := range letters {
	//	fmt.Printf("Value: %s\n", letter)
	//}
	//
	////Only index
	//fmt.Println("\nOnly Index")
	//for i := range letters {
	//	fmt.Printf("Index: %d\n", i)
	//}
	//
	////Without index and value. Just print array values
	//fmt.Println("\nWithout Index and Value")
	//i := 0
	//for range letters {
	//	fmt.Printf("Index: %d Value: %s\n", i, letters[i])
	//	i++
	//}

	//sample := "1£B"
	//fmt.Printf("Length is %d\n", len(sample))

	//// Iterating over map
	//sample := map[string]string{
	//	"a": "x",
	//	"b": "y",
	//}
	//
	////Iterating over all keys and values
	//fmt.Println("Both Key and Value")
	//for k, v := range sample {
	//	fmt.Printf("key :%s value: %s\n", k, v)
	//}
	//
	////Iterating over only keys
	//fmt.Println("\nOnly keys")
	//for k := range sample {
	//	fmt.Printf("key :%s\n", k)
	//}
	//
	////Iterating over only values
	//fmt.Println("\nOnly values")
	//for _, v := range sample {
	//	fmt.Printf("value :%s\n", v)
	//}

	// Iterate over channel
	ch := make(chan string)
	go pushToChannel(ch)
	for val := range ch {
		fmt.Println(val)
	}

}

func pushToChannel(ch chan<- string) {
	ch <- "a"
	ch <- "b"
	ch <- "c"
	close(ch)
}
