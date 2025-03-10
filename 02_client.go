package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		fmt.Println("Error getting response from server")
		return
	}
	defer resp.Body.Close()
	fmt.Printf("Response status: %s\n\n", string(resp.Status))
	for field, content := range resp.Header {
		fmt.Printf("%v: %v\n", field, content)
	}
}
