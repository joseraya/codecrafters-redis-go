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
	for {
		ok := scanner.Scan()
		if !ok {
			break
		}
		request, err := ReadRequest(scanner)
		if err != nil {
			return err
		}
		fmt.Printf("read request [%s]\n", request)
		response, err := HandleCommand(strings.ToUpper(request))
		if err != nil {
			return err
		}
		fmt.Printf("sending response [%s]\n", response)
		_, writeErr := WriteResponse(conn, response)
		if writeErr != nil {
			return writeErr
		}
	}
	return nil

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
