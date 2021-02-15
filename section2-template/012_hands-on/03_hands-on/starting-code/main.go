package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	californiaHotels := []hotel{
		hotel{
			Name:    "hotel1",
			Address: "1 Street1",
			City:    "city1",
			Zip:     "1111",
			Region:  "Southern",
		},
		hotel{
			Name:    "hotel2",
			Address: "2 Street1",
			City:    "city2",
			Zip:     "2222",
			Region:  "Northern",
		},
		hotel{
			Name:    "hotel3",
			Address: "3 Street1",
			City:    "city3",
			Zip:     "3333",
			Region:  "Central",
		},
	}

	err := tpl.Execute(os.Stdout, californiaHotels)
	if err != nil {
		log.Fatalln(err)
	}
}
