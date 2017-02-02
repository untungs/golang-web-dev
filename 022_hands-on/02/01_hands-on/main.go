package main

import (
	"net"
	"log"
	"io"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		io.WriteString(conn, "I see you connected.\n")

		conn.Close()
	}

}
