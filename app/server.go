package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func handleConnection(conn net.Conn) error {
	scanner := bufio.NewScanner(conn)
	var err error = nil
	for scanner.Scan() && err == nil {
		request := ReadRequest(scanner.Text())
		if request == "" {
			continue
		}
		fmt.Printf("read request [%s]\n", request)

		response := HandleCommand(strings.ToUpper(request))

		if response != "" {
			fmt.Printf("sending response [%s]\n", response)
			_, err := conn.Write([]byte(FormatResonse(response)))
			if err != nil {
				return err
			}
		}
	}
	return scanner.Err()
}

func initialize() net.Listener {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	return l
}

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
