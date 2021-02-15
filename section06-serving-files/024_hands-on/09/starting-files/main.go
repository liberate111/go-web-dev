// Serve the files in the "starting-files" folder
// To get your images to serve, use:

// 	func StripPrefix(prefix string, h Handler) Handler
// 	func FileServer(root FileSystem) Handler

// Constraint: you are not allowed to change the route being used for images in the template file

package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	http.HandleFunc("/", surf)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func surf(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
