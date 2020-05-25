package main

import "fmt"

func main() {
	for i := 0; i < 11; i++ {
		if i == 5 {
			fmt.Println("==Hello world!==")
		} else {
			fmt.Println("================")
		}
	}
}

