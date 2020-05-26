package main

import "fmt"
import "io/ioutil"

func read() {
	data, err := ioutil.ReadFile("/etc/shadow")
	if err != nil {
		fmt.Println("Failed to read file:", err)
		return
	}

	fmt.Println(string(data))
}

func main() {
	read()
}
