package main

// made for course by Tero Karvinen, http://terokarvinen.com/2020/go-programming-course-2020-w22/
// made with https://gobyexample.com/http-clients

import (
	"fmt"
	"net/http"
	"log"
	"bufio"
)

func main() {
	response, err := http.Get("http://terokarvinen.com")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	fmt.Println("\nResponse code:", response.Status)
	fmt.Println("\n Headers: ")

	for key, v := range response.Header{
		fmt.Printf("\n%v: %v\n", key, v)
	}

	scanner := bufio.NewScanner(response.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err!= nil {
		log.Fatal(err)
	}
}
