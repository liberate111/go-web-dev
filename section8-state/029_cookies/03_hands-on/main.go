package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "0",
		Path:  "/",
	})
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "dev tools / application / cookies")
}

func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
	}
	track, err := strconv.Atoi(c.Value)
	if err != nil {
		fmt.Println(err)
	}
	track++
	valCookie := strconv.Itoa(track)
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: valCookie,
		Path:  "/",
	})
	fmt.Fprintln(w, "YOUR COOKIE: ", c)
}
