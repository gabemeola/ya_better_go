package main

import (
	"fmt"
	"strings"
)

// Title type
type Title string

// Method
func (t Title) FixCase() string {
	return strings.Title(string(t))
}

// Struct
type Monitor struct {
	Resolution string
	Connector  string
	Value      float64
	onSale     bool
}

func main() {
	var name = Title("the matrix")
	var fixed = name.FixCase()
	fmt.Println("Fixed", fixed)
	fmt.Println("=================")

	// Create a new Monitor Struct
	var monitor = Monitor{
		Resolution: "4k",
	}
	monitor.Connector = "USB-C"
	monitor.Value = 789.99
	fmt.Println(monitor)
}
