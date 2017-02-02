package main

import (
	"net"
	"log"
	"io"
	"bufio"
	"fmt"
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

		go serve(conn)
	}

}

func serve(c net.Conn) {
	defer c.Close()

	scanner := bufio.NewScanner(c)

	for scanner.Scan() {
		txt := scanner.Text()
		fmt.Println(txt)

		if txt == "" {
			break
		}
	}
}
