package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	listener := initialize()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(-2)
		}
		go handleConnection(conn)
	}
}

func initialize() net.Listener {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	return l
}

func handleConnection(conn net.Conn) error {
	scanner := bufio.NewScanner(conn)

	for {
		request, err := ReadStrings(scanner)
		if err != nil {
			return err
		}
		fmt.Printf("read request %s\n", request)

		response := HandleCommand(request)

		fmt.Printf("write response %s\n", response)
		if len(response) > 0 {
			for _, line := range response {
				_, err := conn.Write([]byte(EncodeValue(line)))
				if err != nil {
					return err
				}
			}
		}
	}
}
