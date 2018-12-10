package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"bytes"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	//listen for incoming connections
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	//close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

	for {
		//listen for an incoming connection
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		//handle connections in a new goroutine
		go handleRequest(conn)
	}
}

//handles incoming request
func handleRequest(conn net.Conn) {
	//Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	//Read the incoming connection into the buffer.
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	message := "Hi, I received your message! It was "
	message += strconv.Itoa(reqLen)
	message += " bytes long and that's what it said: \""
	n := bytes.Index(buf, []byte{0})
	message += string(buf[:n-1])
	message += "\" I have no clue what to do with these messages, so Bye!\n"
	//Send a response back to person contacting us.
	conn.Write([]byte(message))
	//Close the connection when you're done with it.
	conn.Close()
}
