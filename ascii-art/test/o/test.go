package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//const programName = "main.go"
const fileName = "test-cases.txt"

type data struct {
	user_arg string
	expected string
}

var testCases = []data{}

func main() {

	getTestCases()
	for _, testCase := range testCases {
		fmt.Println(escapedString(testCase.expected))
	}
}

func getTestCases() {
	user_arg := ""
	parser := ""

	Tests, err := readlines(fileName)
	if err != nil {
		log.Fatalf("Error in readLines function: %s", err)
	}

	for i, line := range Tests {
		if line[0] == '£' {
			if parser != "" {
				testCases = append(testCases, data{user_arg, parser})
				parser = ""
			}
			user_arg = line[1 : len(line)-1]
		} else {
			parser += line + "/n"
		}

		if i == len(Tests)-1 {
			testCases = append(testCases, data{user_arg, parser})
		}
	}
}

func readlines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()

}

func escapedString(s string) string {
	s = fmt.Sprint(s)
	return ("\"" + s + "\"")
}
