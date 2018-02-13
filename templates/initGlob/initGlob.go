package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

// Package scoped tpl
var tpl *template.Template

// Init runs once when you go program is run before main
func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	err := tpl.Execute(os.Stdout, nil)
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
