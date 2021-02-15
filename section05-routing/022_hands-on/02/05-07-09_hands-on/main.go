// Building upon the code from the previous problem:

// Extract the code you wrote to READ from the connection using bufio.NewScanner into its own function called "serve".

// Pass the connection of type net.Conn as an argument into this function.

// Add "go" in front of the call to "serve" to enable concurrency and multiple connections.

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go serve(conn)
	}
}

func serve(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if ln == "" {
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
			break
		}
	}
	defer conn.Close()

	io.WriteString(conn, "connect TCP server")
	fmt.Println("Code get here")
}
