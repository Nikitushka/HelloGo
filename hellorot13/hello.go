package main

import (
	"fmt"
	"strings"
	"flag"
)

func main() {
	// declared variable
	var str string

	// set flag to point to variable
	flag.StringVar(&str, "s", "", "String to rot13")
	flag.Parse()

	// check that variable is not empty, if it is print a message to check help and fill flag with string
	if str != "" {
		rotified := strings.Map(rot13, str)
		fmt.Println("\nString to encode:", str)
		fmt.Println("\nEncoded string:" , rotified, "\n")
	}else {
		fmt.Println("\nPlease enter a valid string to rot13\n\nUse -h or --help for more info\n")
	}
}

// function that takes a rune and adds 13 to it using modulo
// made with the help of golang docs strings.Map example:
// https://golang.org/pkg/strings/#Map
func rot13(r rune) rune{
	if r >= 'a' && r <= 'z' {
		return 'a' + (r -'a' + 13) % 26
	}else if r >= 'A' && r <= 'Z'{
		return 'A' + (r - 'A' + 13) % 26
	}
	return r
}
