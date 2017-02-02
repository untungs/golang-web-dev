package main

import (
	"net"
	"log"
	"io"
	"bufio"
	"fmt"
	"strings"
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

	var m string
	var u string
	scanner := bufio.NewScanner(c)

	for i := 0; scanner.Scan(); i++ {
		txt := scanner.Text()
		fmt.Println(txt)

		if i == 0 {
			m = strings.Fields(txt)[0]
			u = strings.Fields(txt)[1]
			fmt.Println("***Method:", m)
			fmt.Println("***URI:", u)
		}
		if txt == "" {
			break
		}
	}

	body := fmt.Sprintln("I see you connected.\nMethod:", m, "\nURI:", u)
	io.WriteString(c, "HTTP/1.1 302 Found\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/plain\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}
