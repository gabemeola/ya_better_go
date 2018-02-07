package main

import (
	"fmt"
)

type Car struct {
	Doors        int64
	Transmission string
	Odometer     int64
}

func main() {
	var value float64 = 5
	var pointer = &value
	fmt.Println("pointer", pointer)
	fmt.Println("*pointer", *pointer)

	// Passing the memory address of value
	halve(&value)
	fmt.Println("Number is:", value)

	// Car Struct
	car := Car{
		Doors:        4,
		Transmission: "automatic",
		Odometer:     36000,
	}
	// Passing as a pointer prevents a copy from being made for arguments
	// This can help with having a smaller memory footprint
	describe(&car)
}

// Modifying the original number
func halve(number *float64) {
	*number = *number / 2
}

func describe(c *Car) {
	fmt.Println()
	fmt.Println("A", c.Doors, "door")
	fmt.Println(c.Transmission, "car")
	fmt.Println("with", c.Odometer, "miles")
}
