package main

import "fmt"

type Exmpl struct {
	X, Y int
	Z string
}

func main() {
	var p Exmpl
	p.X, p.Y, p.Z = 30, -10, "Viidakko"
	fmt.Printf("\nX kordinaatti: %v\nY kordinaatti: %v\nOlemme kohteessa: %v\n\n", p.X, p.Y, p.Z)
}
