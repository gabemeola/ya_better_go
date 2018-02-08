package main

import (
	"log"
	"os"
)

func main() {
	var file, err = os.Create("TestFile.txt")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("Oooh we got some cool test code in here")

}
