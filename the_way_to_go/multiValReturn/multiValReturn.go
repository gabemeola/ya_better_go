package main

import (
	"fmt"
	"math"
)

func main() {
	// Unnamed Return
	fmt.Println(multiValReturn(30, 120))
	// Named Return
	fmt.Println(multiValReturnNamed(5, 17))
}

func multiValReturn(value1 int, value2 int) (int, int, int) {
	sum := value1 + value2
	product := value1 * value2
	difference := int(math.Abs(float64(value1 - value2)))

	return sum, product, difference
}

func multiValReturnNamed(value1 int, value2 int) (sum int, product int, difference int) {
	sum = value1 + value2
	product = value1 * value2
	difference = int(math.Abs(float64(value1 - value2)))
	return
}
