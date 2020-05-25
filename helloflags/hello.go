package main

import (
	"fmt"
	"flag"
)

func main() {
	// declare variables
	var word string
	var count int

	// tell "flag" to put the flags into variables
	flag.StringVar(&word, "s", "", "enter the string you want to loop")
	flag.IntVar(&count, "n", 0, "enter the amount of times to loop")

	// parse the command line arguments
	flag.Parse()

	// check if the flags are set
	if word == "" || count < 0 {
		fmt.Println("\nPlease set all of the flags \n\nUse -h or --help flag for help\n")
	} else {
		fmt.Println(loopity(word, count))
	}
}

// print the word as many times as the count is set to
func loopity(s string, n int) string {
	output := ""
	for i := 0; i < n; i++{
		if i + 1 == n{
		output += "\n" + s + "\n"
	}else{
		output += "\n" + s
	}
	}
	return output
}
