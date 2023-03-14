package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {

	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	fmt.Println("Menu")
	fmt.Println("____")

	fmt.Println("1 - Coffee")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Cappucino")
	fmt.Println("4 - Chai tea")
	fmt.Println("5 - Tea")
	fmt.Println("5 - Americano")
	fmt.Println("Q - Quit the program")

	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		i, _ := strconv.Atoi(string(char))
		t := fmt.Sprintf("You chose %d", i)
		fmt.Println(t)

		if char == 'q' || char == 'Q' {
			break
		}
	}
	fmt.Println("Program exiting...")
}
