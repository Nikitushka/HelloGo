package main

import (
	"fmt"
	"strings"
	"flag"
)

func main() {
	var str string

	flag.StringVar(&str, "s", "", "String to rot13")
	flag.Parse()

	if str != "" {
		rotified := strings.Map(rot13, str)
		fmt.Println("\nString to encode:", str)
		fmt.Println("\nEncoded string:" , rotified, "\n")
	}else {
		fmt.Println("\nPlease enter a valid string to rot13\n\nUse -h or --help for more info\n")
	}
}

func rot13(r rune) rune{
	if r >= 'a' && r <= 'z' {
		return 'a' + (r -'a' + 13) % 26
	}else if r >= 'A' && r <= 'Z'{
		return 'A' + (r - 'A' + 13) % 26
	}
	return r
}
