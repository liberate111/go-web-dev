package session

import (
	"fmt"
	"net/http"
	"time"

	"mvc/models"

	uuid "github.com/satori/go.uuid"
)

const Length int = 30

var Users = map[string]models.User{}       // userID, user
var Sessions = map[string]models.Session{} // sessionID, session
var LastCleaned time.Time = time.Now()

func GetUser(w http.ResponseWriter, req *http.Request) models.User {

	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	c.MaxAge = Length
	http.SetCookie(w, c)

	// if the user exists already, get user
	var u models.User
	if s, ok := Sessions[c.Value]; ok {
		s.LastActivity = time.Now()
		Sessions[c.Value] = s
		u = Users[s.Username]
	}
	return u
}

func AlreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := Sessions[c.Value]
	if ok {
		s.LastActivity = time.Now()
	}
	_, ok = Users[s.Username]
	// refresh session
	c.MaxAge = Length
	http.SetCookie(w, c)
	return ok
}

func Clean() {
	fmt.Println("BEFORE CLEAN")
	Show()

	for k, v := range Sessions {
		if time.Now().Sub(v.LastActivity) > (time.Second * 30) {
			delete(Sessions, k)
		}
	}
	LastCleaned = time.Now()
	fmt.Println("AFTER CLEAN")
	Show()
}

func Show() {
	fmt.Println("=========")
	for k, v := range Sessions {
		fmt.Println(k, v.Username)
	}
	fmt.Println("_________")
}
