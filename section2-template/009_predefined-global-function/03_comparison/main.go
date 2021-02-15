package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	g1 := struct {
		Score1 int
		Score2 int
	}{
		7,
		9,
	}

	err := tpl.Execute(os.Stdout, g1)
	if err != nil {
		log.Fatalln(err)
	}
}

// eq
// Returns the boolean truth of arg1 == arg2

// ne
// Returns the boolean truth of arg1 != arg2

// lt
// Returns the boolean truth of arg1 < arg2

// le
// Returns the boolean truth of arg1 <= arg2

// gt
// Returns the boolean truth of arg1 > arg2

// ge
// Returns the boolean truth of arg1 >= arg2

// source:
// https://godoc.org/text/template#Functions
