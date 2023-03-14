package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func menu() {

	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	beverages := make(map[int]string)
	beverages[1] = "Coffee"
	beverages[2] = "Latte"
	beverages[3] = "Cappucino"
	beverages[4] = "Chai tea"
	beverages[5] = "Tea"
	beverages[6] = "Americano"

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
		if char == 'q' || char == 'Q' {
			break
		}

		i, _ := strconv.Atoi(string(char))
		fmt.Println(fmt.Sprintf("You chose %s", beverages[i]))

	}
	fmt.Println("Program exiting...")
}
