package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func writeString(conn net.Conn, message string) (int, error) {
	msg := fmt.Sprintf("+%s\r\n", message)
	n, err := conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("Error writing the response", err.Error())
		return 0, err
	}
	return n, nil
}

func handleCommand(command string) (string, error) {
	var response string
	if command == "PING" {
		response = "PONG"
	} else {
		response = ""
	}
	return response, nil
}
func handleConnection(l net.Listener) error {
	var err error

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		return err
	}
	scanner := bufio.NewScanner(conn)
	for {
		ok := scanner.Scan()
		if !ok {
			break
		}
		command := scanner.Text()
		fmt.Printf("read command [%s]\n", command)
		firstCharacter := command[0:1]
		// * is for arrays and $ is for bulk strings
		// we can ignore both for now and respond only to commands
		if firstCharacter == "*" || firstCharacter == "$" {
			continue
		}
		response, err := handleCommand(strings.ToUpper(command))
		if err != nil {
			return err
		}
		fmt.Printf("sending response [%s]\n", response)
		_, writeErr := writeString(conn, response)
		if writeErr != nil {
			return writeErr
		}
	}
	return nil

}
func main() {

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	err = handleConnection(l)
	if err != nil {
		fmt.Println("Error handling connection", err.Error())
		os.Exit(1)
	}
}
