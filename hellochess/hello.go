package main

import "fmt"

func main(){
	white := "B"
	black := "W"
	var line string
	var board string

	for i := 0; i < 8; i++{
		for c := 0; c < 8; c++{
			if i % 2 == 0 {
				line += white + black
			} else {
				line += black + white
			}
		}
		board += "\n" + line
		line = ""
	}
	fmt.Println(board)
}
