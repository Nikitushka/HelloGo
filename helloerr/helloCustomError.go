package main

import (
	"fmt"
	"errors"
)

func dividefour(x int) (int, error) {
	if x == 0{
		return 0, errors.New("You can't do that foo'!")
	} else {
		return 4 / x, nil
}
}

func main() {
	f, err := dividefour(0)
	if err != nil {
		fmt.Println("Error: ",err)
	}else {
		fmt.Println(f)
	}
}
