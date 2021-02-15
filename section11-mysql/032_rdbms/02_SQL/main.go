package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	// user:password@tcp(localhost:5555)/dbname?charset=utf8
	db, err = sql.Open("mysql", "admin:chayanon@tcp(aws-rds-database-1.cdlgv1w7rv7w.ap-southeast-1.rds.amazonaws.com:3306)/test01?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/cVac", cVac)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/drop", drop)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	check(err)
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err = io.WriteString(w, "Welcome to report progress country vaccine COVID19")
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func cVac(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT country FROM country_vaccinations;`)
	check(err)
	defer rows.Close()

	// data to be used in query
	var s, c string
	s = "RETRIVED RECORD:\n"

	// query
	for rows.Next() {
		err = rows.Scan(&c)
		check(err)
		s += c + "\n"
	}
	fmt.Fprintln(w, s)
}

func create(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`CREATE TABLE recover (country VARCHAR(20));`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "CREATED TABLE recover", n)
}

func insert(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`INSERT INTO recover VALUES ("Thailand");`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "INSERTED RECORD", n)
}

func read(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT * FROM recover;`)
	check(err)
	defer rows.Close()

	var c string
	for rows.Next() {
		err = rows.Scan(&c)
		check(err)
		fmt.Fprintln(w, "RETRIEVED RECORD:", c)
	}
}

func update(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`UPDATE recover SET country="Singapore" WHERE country="Thailand";`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "UPDATED RECORD", n)
}

func delete(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DELETE FROM recover WHERE country="Singapore";`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "DELETED RECORD", n)
}

func drop(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DROP TABLE recover;`)
	check(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	check(err)

	fmt.Fprintln(w, "DROPPED TABLE recover")
}
