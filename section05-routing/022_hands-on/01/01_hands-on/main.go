// ListenAndServe on port ":8080" using the default ServeMux.

// Use HandleFunc to add the following routes to the default ServeMux:

// "/" "/dog/" "/me/

// Add a func for each of the routes.

// Have the "/me/" route print out your name.

package main

import (
	"io"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello from HandleFunc #1 /\n")
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello from HandleFunc #2 /dog/\n")
}

func me(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello Fufuu from HandleFunc #3 /me/\n")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
