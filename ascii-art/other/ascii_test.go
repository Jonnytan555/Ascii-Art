package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"
)

// **************************
// NAME_OF_PROGRAM
const programName = "main.go"

// **************************

// All additional test cases will be read from this file
var fileName string = "test-cases.txt"

type testData struct {
	arg      string
	expected string
}

var testCases = []testData{}

// TestAsciiArt will test the output of the main.go program to match the desired representation in Ascii Art characters
func TestAsciiArt(t *testing.T) {
	// add additional testcases from the file into the testCases
	addTestCases(fileName)

	for _, test := range testCases {

		cmd := "go run " + programName + " " + escapedString(test.arg) + " | cat -e"

		output, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			t.Fatal(">>> TRY CHECKING the input argument in the", fileName, " file. For >>>> ", test.arg, " \n", err)
		}

		if string(output) != (test.expected) {
			t.Errorf("For argument '%v' \n\033[1;34mexpected \n%s\033[0m\n\033[1;31mbut got \n%s\033[0m", test.arg, test.expected, string(output))
			panic("err")
		} else {
			t.Logf("output for \033[38;5;105m\"%s\"\033[39;49m is \n\033[1;32m%s\033[0m", test.arg, output)
			t.Log("\n")
			t.Logf("expected output for \033[38;5;105m\"%s\"\033[39;49m is \n\033[1;32m%s\033[0m\n\033[1;36m------------------------------\033[0m\n", test.arg, test.expected)
		}
	}
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
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

// addTestCases will read the contents of the specified file, extract and add the desired output into the testCases
func addTestCases(fileName string) {
	arg, tempStr := "", ""

	lines, err := readLines(fileName)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	for i, line := range lines {
		if line[0] == '#' {
			if tempStr != "" {
				testCases = append(testCases, testData{arg, tempStr})
				tempStr = ""
			}
			arg = string(line[1 : len(line)-1])

		} else {
			tempStr += line + "\n"
		}
		// Last test case(No more testcases in the file)
		if i == len(lines)-1 {
			testCases = append(testCases, testData{arg, tempStr})
		}
	}
}

// escapedString is used to pass a string that is already escaped as an argument into the command line
// example:	"1a\"#F"
// becomes:	"1a"#F"
func escapedString(s string) string {
	s = fmt.Sprint(s)
	return ("\"" + s + "\"")
}
