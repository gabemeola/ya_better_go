package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

func firstThree(str string) string {
	return strings.TrimSpace(str)[:3]
}

var funcMap = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(funcMap).ParseFiles("tpl.html"))
}

type sage struct {
	Name  string
	Motto string
}

func main() {
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	gabe := sage{
		Name:  "Gabe",
		Motto: "CODE CODE CODE!",
	}

	sages := []sage{buddha, gandhi, gabe}

	data := struct {
		Wisdom []sage
	}{
		Wisdom: sages,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
