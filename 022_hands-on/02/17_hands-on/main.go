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

	switch {
	case m == "GET" && u == "/":
		index(c)
	case m == "GET" && u == "/apply":
		apply(c)
	case m == "POST" && u == "/apply":
		applyPost(c)
	default:
		notFound(c)
	}
}

func index(c net.Conn) {
	body := `
		<h1>INDEX</h1>
		<a href="/apply">apply</a>
	`

	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}

func apply(c net.Conn)  {
	body := `
		<h1>APPLY</h1>
		<a href="/">index</a>
		<form method="POST" action="/apply">
			<input type="submit" value="submit">
		</form>
	`

	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}

func applyPost(c net.Conn)  {
	body := `
		<h1>APPLY POST</h1>
		<a href="/">index</a><br>
		<a href="/apply">apply</a>
	`

	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}

func notFound(c net.Conn) {
	body := "404 gopher not found"

	io.WriteString(c, "HTTP/1.1 404 Not Found\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}