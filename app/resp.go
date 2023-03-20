package main

import (
	"bufio"
	"fmt"
	"strconv"
)

func ReadStrings(scanner *bufio.Scanner) ([]string, error) {
	// Read the first line to know how many lines we need to read
	ok := scanner.Scan()
	if !ok {
		return nil, fmt.Errorf("error reading request")
	}
	line := scanner.Text()
	if line[0:1] != "*" {
		return nil, fmt.Errorf("error reading request, first line: %s", line)
	}

	numValues, err := strconv.Atoi(line[1:])
	if err != nil {
		return nil, err
	}

	valuesRead := 0
	command := make([]string, numValues)
	for valuesRead < numValues {
		line, err = readString(scanner)
		if err != nil {
			return nil, err
		}
		command[valuesRead] = line
		valuesRead++
	}
	return command, nil
}

func readString(scanner *bufio.Scanner) (string, error) {
	ok := scanner.Scan()
	if !ok {
		return "", fmt.Errorf("error scanning")
	}
	line := scanner.Text()
	if line[0:1] == "$" {
		//This is the length of the string, read the next one
		return readString(scanner)
	} else {
		return line, nil
	}
}

func EncodeValue(message string) string {
	msg := fmt.Sprintf("+%s\r\n", message)
	return msg
}
