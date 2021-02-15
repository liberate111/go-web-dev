package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
	Id     string `json:"id"`
}

func StoreUsers(m map[string]User) {
	f, err := os.Create("dataJSON")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	err = json.NewEncoder(f).Encode(m)
	if err != nil {
		fmt.Println(err)
	}
}

func LoadUsers() map[string]User {
	m := make(map[string]User)

	f, err := os.Open("dataJSON")
	if err != nil {
		fmt.Println(err)
		return m
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&m)
	if err != nil {
		fmt.Println(err)
	}
	return m
}
