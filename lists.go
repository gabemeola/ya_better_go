package main

import (
	"fmt"
)

func sep() {
	fmt.Println("==========================")
}

func main() {
	fmt.Println("Arrays, Slices, n Maps!\n")

	// arrays()
	// sep()
	// slices()
	// sep()
	maps()

}

// Arrays are fixed length
func arrays() {
	// Init with data
	var months = [3]string{"April", "May", "June"}
	// Adding data later
	var salesByMonth [3]float64
	salesByMonth[0] = 1710.26
	salesByMonth[1] = 2245
	salesByMonth[2] = 3032.40

	// Brittle for loop because we will get an
	// fatal error if we loop over value that doesn't exists
	for i := 0; i < len(months); i++ {
		fmt.Println(months[i], salesByMonth[i])
	}

	fmt.Println()

	// Safer and less error prone
	for i, month := range months {
		fmt.Println(month, salesByMonth[i])
	}
}

func slices() {
	// To use a slice, don't define a length
	// Go will create an underlining Array for you
	var slice = []int{1, 2, 3}
	// When appending to a slice, Go will create a
	// new Bigger array and copy over the values
	slice = append(slice, 4, 5)
	slice = append(slice, 6, 7, 8)

	fmt.Println("cap", cap(slice))
	fmt.Println("slice", slice)

	for _, num := range slice {
		fmt.Print(num, " ")
	}
	fmt.Println()
}

// Maps are also called HashMaps (Java), Dictionaries (C#), Hash (Ruby)
func maps() {
	var students = map[string]float64{
		"Gabe":  20,
		"Josif": 21,
	}
	students["Alice"] = 33
	students["Jim"] = 19

	fmt.Println(students)
	fmt.Println(students["Alice"], students["Bob"])

	for name, age := range students {
		fmt.Println(name, "is", age, "years old.")
	}
}
