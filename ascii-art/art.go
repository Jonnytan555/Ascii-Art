package art

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func Art() {
	output, err := os.Open("test.txt")
	if err != nil {
		log.Fatalf("failed to open")
	}

	NewData := bufio.NewScanner(io.Reader(output))

	// Scan the lines from the text file

	var lines []string

	NewData.Split(bufio.ScanLines)
	for NewData.Scan() {
		lines = append(lines, NewData.Text())
	}

	// Make a map for ascii characters --> lines

	aMap := make(map[int][]string)

	// Fill up the map with the lines
	// Each ascii character chnages when there is a line break

	i := 31

	for _, line := range lines {
		if line == "" {
			i++
		} else {
			aMap[i] = append(aMap[i], line)
		}
	}

	arguments := os.Args[1]

	// Check if user input contains a newline character
	// if so, split the arguments by this newline character

	argsChecked := strings.Split(arguments, "\\n")

	// For each charcter of the arguements, loop and print out each line

	for _, word := range argsChecked {
		for j := 0; j < 8; j++ {
			for i := range word {
				fmt.Print(string(aMap[int(word[i])][j]))
			}
			fmt.Println()
		}
	}

}
