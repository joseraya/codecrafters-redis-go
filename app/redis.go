package main

func HandleCommand(command string) (string, error) {
	var response string
	if command == "PING" {
		response = "PONG"
	} else {
		response = ""
	}
	return response, nil
}
