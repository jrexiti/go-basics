package main

import "fmt"

var myInt int
var myInt32 int32
var myInt64 int64
var myFloa32 float32
var myFloat64 float64

type Owner struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string
}

type Car struct {
	Owner
	Make      string
	Model     string
	Year      int
	Available bool
}

func types() {

	fmt.Printf("type of myInt: %T\n", myInt)
	fmt.Printf("type of myInt32: %T\n", myInt32)
	fmt.Printf("type of myInt64: %T\n", myInt64)
	fmt.Printf("type of myFloat32: %T\n", myFloa32)
	fmt.Printf("type of myFloat64: %T\n", myFloat64)

	var JaysCar Car

	JaysCar.FirstName = "Jay"
	JaysCar.LastName = "Rexiti"
	JaysCar.PhoneNumber = "2026603881"
	JaysCar.Email = "jay.rexiti@gmail.com"
	JaysCar.Make = "Porsche"
	JaysCar.Model = "Cayenne"
	JaysCar.Year = 2021
	JaysCar.Available = false

	fmt.Println(JaysCar)

}
