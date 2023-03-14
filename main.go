package main

import (
	"fmt"
	"math/rand"
	"time"

	"www.github.com/jrexiti/go-basics/packageone"
)

func main() {
	packageone.Exported()
	rand.Seed(time.Now().UnixNano())
	newstring := packageone.PublicVar
	fmt.Println("From packageone:", newstring)
	aString := packageone.PublicVar
	fmt.Println(aString)
	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	var answer = firstNumber*secondNumber - subtraction
	theGame(firstNumber, secondNumber, subtraction, answer)

}
