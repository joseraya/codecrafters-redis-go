package main

import (
	"strings"
)

func HandleCommand(request []string) []string {
	if len(request) == 0 {
		return nil
	}

	switch cmd := strings.ToUpper(request[0]); cmd {
	case "PING":
		return ping()
	case "ECHO":
		return echo(request[1:])
	case "COMMAND":
		return command(request[1:])
	default:
		return nil
	}
}

func ping() []string {
	return []string{"PONG"}
}

func echo(args []string) []string {
	return args
}

func command(args []string) []string {
	return []string{"DOCS"}
}
