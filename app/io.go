package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func ReadRequest(scanner *bufio.Scanner) (string, error) {
	command := ""
	for command == "" {
		ok := scanner.Scan()
		if !ok {
			return "", fmt.Errorf("error reading from connection")
		}
		command = scanner.Text()
		//fmt.Printf("read [%s]\n", command)
		firstCharacter := command[0:1]
		// * is for arrays and $ is for bulk strings
		// we can ignore both for now and respond only to commands
		if firstCharacter == "*" || firstCharacter == "$" {
			// ignore this line and go to the next
			command = ""
		}
	}
	return strings.ToUpper(command), nil
}

func WriteResponse(conn net.Conn, message string) (int, error) {
	msg := fmt.Sprintf("+%s\r\n", message)
	n, err := conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("Error writing the response", err.Error())
		return 0, err
	}
	return n, nil
}
