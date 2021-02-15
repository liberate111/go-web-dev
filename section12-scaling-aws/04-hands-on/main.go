package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	var err error

	// user:password@tcp(localhost:5555)/dbname?charset=utf8
	db, err = sql.Open("mysql", "admin:chayanon@tcp(aws-rds-database-1.cdlgv1w7rv7w.ap-southeast-1.rds.amazonaws.com:3306)/test01?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/instance", instance)
	http.HandleFunc("/cVac", cVac)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello from AWS.")
}

func ping(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "OK")
}

func instance(w http.ResponseWriter, req *http.Request) {
	s := getInstance()
	io.WriteString(w, s)
}

func getInstance() string {
	resp, err := http.Get("http://169.254.169.254/latest/meta-data/instance-id")
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	bs := make([]byte, resp.ContentLength)
	resp.Body.Read(bs)
	resp.Body.Close()
	return string(bs)
}

func cVac(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT country FROM country_vaccinations;`)
	check(err)
	defer rows.Close()

	// data to be used in query
	s := getInstance()
	s += "RETRIVED RECORD:\n"

	var c string
	// query
	for rows.Next() {
		err = rows.Scan(&c)
		check(err)
		s += c + "\n"
	}
	fmt.Fprintln(w, s)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
