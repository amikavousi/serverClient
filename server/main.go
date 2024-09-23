package main

import (
	"fmt"
	"log"
	"net"
)

const NETWORK = "tcp"
const ADDRESS = "127.0.0.1:8000"

func main() {
	// set listener
	listener, err := net.Listen(NETWORK, ADDRESS)
	defer func() {
		e := listener.Close()
		if e != nil {
			log.Fatalln("cant close listener")
		}
	}()

	if err != nil {
		log.Fatalln("cant listen to address")
	}

	fmt.Println("Server listening at:", listener.Addr())

	//	process requests
	for {
		connection, aErr := listener.Accept()

		if aErr != nil {
			_ = fmt.Errorf("we have error in connection, %e", aErr)
		}

		data := make([]byte, 1200)
		_, rErr := connection.Read(data)
		if rErr != nil {
			fmt.Println("cant read data from server", rErr)
		}

		fmt.Printf("new message => addres: %s , data: %s\n",
			connection.RemoteAddr(), string(data))
	}
}
