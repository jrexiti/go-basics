package main

import "fmt"

func main() {
	fmt.Println("hello world!")

	var a = 1
	var b = &a
	*b = 4
	a = 9

	fmt.Println(b, *b, a)

	myChannel := make(chan string)

	go func() {
		myChannel <- "data"
	}()

	msg := <-myChannel

	fmt.Println(msg)
}
