package main

import (
	"fmt"
	"strings"
)

func main() {
	var baseString = "BASE STRING:"

	var abcs = "ABCDEFGUIJKLMNOPQRSTUZWXYZ"
	var oneTwoThrees = "12345678910"
	var gabe = "gabe"
	var josif = "josif"
	var joon = "joon"

	// Concat those strings!
	concat(
		&baseString,
		abcs,
		oneTwoThrees,
		gabe,
		josif,
		joon,
	)

	fmt.Println("New String:", baseString)
}

func concat(concatString *string, strs ...string) {
	for _, runeChar := range strs {
		fmt.Println("Concating:", string(runeChar))
	}

	fmt.Println()
	*concatString = strings.Join(strs, "")
}
