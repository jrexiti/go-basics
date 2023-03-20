package main

import "fmt"

func main() {
	fmt.Println("hello world!")

	a := 1
	b := &a
	*b = 4
	a = 9

	fmt.Println(b, *b, a)

	channelOne := make(chan string)
	channelTwo := make(chan string)
	channelThree := make(chan string)

	go func() {
		channelOne <- "first channel"
	}()

	go func() {
		channelTwo <- "second channel"
	}()
	go func() {
		channelThree <- "third channel"
	}()

	select {
	case msgFromChannel := <-channelOne:
		fmt.Println("message from:", msgFromChannel)
	case msgFromChannelTwo := <-channelTwo:
		fmt.Println("message from:", msgFromChannelTwo)
	case msgFromThirdChannel := <-channelThree:
		fmt.Println("message from:", msgFromThirdChannel)
	}

}
