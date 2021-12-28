package main

import "art"
import "os"
import "fmt"

func main() {
	if len(os.Args) == 2 {
		art.Art()
	} else {
		fmt.Print("Please enter a word to print in ascii art\n")
	}
}
