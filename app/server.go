package main

import (
	"fmt"
	"net"
	"os"
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
func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	buf := make([]byte, 0, 4096)
	_, err = conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading from the connection", err.Error())
	}
	_, writeErr := writeString(conn, "PONG")

	if writeErr != nil {
		fmt.Println("Write error", writeErr)
	}
}
