package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundance", abundance)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "fufuu",
		Path:  "/",
	})
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "dev tools / application / cookies")
}

func read(w http.ResponseWriter, req *http.Request) {
	c1, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #1: ", c1)
	}

	c2, err := req.Cookie("my-cookie2")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #2: ", c2)
	}

	c3, err := req.Cookie("my-cookie3")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #3: ", c3)
	}
}

func abundance(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie2",
		Value: "fufuu22222222222222",
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie3",
		Value: "fufuu333333333333",
	})

	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "dev tools / application / cookies")
}
