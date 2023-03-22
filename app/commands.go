package main

import (
	"strconv"
	"strings"
	"time"
)

type entry struct {
	value   string
	expires int64
}

// This is not thread safe. Probably would break in a real-life
// scenario. We should probably be using https://pkg.go.dev/sync#Map instead
var values map[string]entry

func init() {
	values = map[string]entry{}
}

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
	case "SET":
		return set(request[1:])
	case "GET":
		return get(request[1:])
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

func set(args []string) []string {
	key := args[0]
	value := args[1]
	var expires int64 = 0
	if len(args) > 2 && strings.ToUpper(args[2]) == "PX" {
		expiryMs, err := strconv.Atoi(args[3])
		if err != nil {
			return nil
		}
		expires = time.Now().UnixMilli() + int64(expiryMs)
	}
	values[key] = entry{
		value:   value,
		expires: expires,
	}
	return []string{"OK"}
}

func get(args []string) []string {
	key := args[0]
	value := values[key]
	now := time.Now().UnixMilli()
	if value.expires != 0 && now >= value.expires {
		delete(values, key)
		return []string{}
	} else {
		return []string{value.value}
	}
}
