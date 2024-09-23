package main

import (
	"log"
	"net"
	"os"
)

const NETWORK = "tcp"
const ADDRESS = "127.0.0.1:8000"

func main() {
	msg := "default msg"

	if len(os.Args) > 1 {
		msg = os.Args[1]
	}

	conn, dErr := net.Dial(NETWORK, ADDRESS)
	if dErr != nil {
		log.Fatalln("cant connect to server ", conn.RemoteAddr())
	}

	_, wErr := conn.Write([]byte(msg))
	if wErr != nil {
		log.Fatalln("cant write data to connection ", wErr)
	}
}
