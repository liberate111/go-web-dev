// ListenAndServe on port ":8080" using the default ServeMux.

// Use HandleFunc to add the following routes to the default ServeMux:

// "/" "/dog/" "/me/

// Add a func for each of the routes.

// Have the "/me/" route print out your name.

// Take the previous program in the previous folder and change it so that:

// a template is parsed and served
// you pass data into the template

package main

import (
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	//io.WriteString(w, "Hello from HandleFunc #1 /\n")
	tpl.ExecuteTemplate(w, "index.gohtml", "Hello from HandleFunc #1 /\n")
}

func dog(w http.ResponseWriter, req *http.Request) {
	//io.WriteString(w, "Hello from HandleFunc #2 /dog/\n")
	tpl.ExecuteTemplate(w, "index.gohtml", "Hello from HandleFunc #2 /dog/\n")
}

func me(w http.ResponseWriter, req *http.Request) {
	//io.WriteString(w, "Hello Fufuu from HandleFunc #3 /me/\n")
	tpl.ExecuteTemplate(w, "index.gohtml", "Hello Fufuu from HandleFunc #3 /me/\n")
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
