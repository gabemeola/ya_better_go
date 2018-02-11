package main

import (
	"fmt"
	"strconv"
	"time"
)

var pizzaNum = 0

func makeDough(stringChan chan string) {
	pizzaNum++

	var pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)

	fmt.Println("Made Dough. Sending for Sauce!")

	stringChan <- pizzaName

	time.Sleep(time.Millisecond * 10)
}

func makeSauce(stringChan chan string) {
	var pizza = <-stringChan

	fmt.Println("Sauce Added! Sending", pizza, "for toppings.")

	stringChan <- pizza

	time.Sleep(time.Millisecond * 10)
}

func makeToppings(stringchan chan string) {
	var pizza = <-stringchan

	fmt.Println("Added Toppings to Pizza.", pizza, "Ready to be Delivered.")
}

func makePizza() {
	// Creating a new channel
	var pizzaChan = make(chan string)

	go makeDough(pizzaChan)
	go makeSauce(pizzaChan)
	go makeToppings(pizzaChan)
}

func main() {
	for i := 0; i < 3; i++ {
		go makePizza()

		time.Sleep(time.Millisecond * 5000)
	}
}
