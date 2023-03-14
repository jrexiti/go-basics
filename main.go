package main

import (
	"math/rand"
	"time"
)

const prompt = "and press ENTER when ready."

func main() {
	rand.Seed(time.Now().UnixNano())

	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	var answer = firstNumber*secondNumber - subtraction
	theGame(firstNumber, secondNumber, subtraction, answer)

}
