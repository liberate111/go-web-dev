package models

import "time"

// capitalize to export from package
type User struct {
	Username string
	Password []byte
	First    string
	Last     string
	Role     string
}

// capitalize to export from package
type Session struct {
	Username     string
	LastActivity time.Time
}
