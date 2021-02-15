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
	sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}

}
