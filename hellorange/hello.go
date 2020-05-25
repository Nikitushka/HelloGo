package main

import "fmt"

func main(){
	var seinfeld = []string{"Jerry", "Kramer", "George", "Ellaine"}

	for _, cast := range seinfeld{
		fmt.Println(cast)
	}
}
