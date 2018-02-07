package main

import (
	"fmt"
	"log"
	"math"

	"./src/name"
	"github.com/golang/example/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("Hello World!"))
	fmt.Println(name.Mine)
	fmt.Println()
	fmt.Println("Multiply:", multiply(3, 9.4))
	root, rootErr := squareRoot(-3)
	if rootErr != nil {
		log.Fatal(rootErr)
	}
	fmt.Println("Square Root:", root)

}

func multiply(first float64, second float64) int {
	return int(first * second)
}

func squareRoot(num float64) (float64, error) {
	if num < 0 {
		return 0, fmt.Errorf("Can't take square root of a negative number")
	}
	return math.Sqrt(num), nil
}
