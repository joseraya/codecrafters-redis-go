package main

import (
	"fmt"
	"strings"
)

func ReadRequest(input string) string {
	// * is for arrays and $ is for bulk strings
	// we can ignore both for now and respond only to commands
	firstCharacter := input[0:1]
	if firstCharacter == "*" || firstCharacter == "$" {
		return ""
	} else {
		return strings.ToUpper(input)
	}
}

func FormatResonse(message string) string {
	msg := fmt.Sprintf("+%s\r\n", message)
	return msg
}

func HandleCommand(command string) string {
	var response string
	if command == "PING" {
		response = "PONG"
	} else {
		response = ""
	}
	return response
}
