package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	// Write
	fmt.Fprintln(conn, "I dialed you.")

	// Read
	// bs, err := ioutil.ReadAll(conn)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Println(string(bs))

}
