package main

import "fmt"

func greetTwo(x, y string) (string, string){
	h := "Hello "
	return h + x, h + y
}

func main() {
	var first, second = greetTwo("World", "Underworld")

	fmt.Println(first, second)
}
