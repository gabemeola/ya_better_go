package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\n\n===========[Gabe Template]============")
	err = tpl.ExecuteTemplate(os.Stdout, "gabe.html", nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\n\n============[Joon Template]============")
	err = tpl.ExecuteTemplate(os.Stdout, "joon.html", nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\n\n============[Josif Template]============")
	err = tpl.ExecuteTemplate(os.Stdout, "josif.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
