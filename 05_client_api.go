package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type

func main() {
	resp, err_get := http.Get("https://pokeapi.co/api/v2/pokedex/kanto/")
	if err_get != nil {
		panic(err_get)
	}
	defer resp.Body.Close()
	fmt.Printf("Response status: %s\n\n", string(resp.Status))
	for field, content := range resp.Header {
		fmt.Printf("%v: %v\n", field, content)
	}

	fmt.Println()

	byteValue, err_readall := ioutil.ReadAll(resp.Body)
	if err_readall != nil {
		panic(err_readall)
	}
	

}
