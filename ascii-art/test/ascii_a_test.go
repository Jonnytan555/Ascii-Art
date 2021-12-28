package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"
)

const programName = "main.go"
const fileName = "test-cases.txt"

type data struct {
	user_arg string
	expected string
}

var testCases = []data{}

func TestAsciiArt(t *testing.T) {
	// func to get the test cases from the file
	getTestCases()

	for _, testCase := range testCases {

		cmd := "go run " + programName + " " + escapedString(testCase.user_arg) + "| cat -e"

		output, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			t.Fatal("Check your", fileName, "and", testCase.user_arg, " \n", err)
		}

		if string(output) != (testCase.expected) {
			t.Errorf("For argument '%v' \n\033[1;34mexpected \n%s\033[0m\n\033[1;31mbut got \n%s\033[0m", testCase.user_arg, testCase.expected, string(output))
			panic("err")
		} else {
			t.Logf("output for \033[38;5;105m\"%s\"\033[39;49m is \n\033[1;32m%s\033[0m", testCase.user_arg, output)
			t.Log("\n")
			t.Logf("expected output for \033[38;5;105m\"%s\"\033[39;49m is \n\033[1;32m%s\033[0m\n\033[1;36m------------------------------\033[0m\n", testCase.user_arg, testCase.expected)
		}
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
