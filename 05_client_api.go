package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PokemonSpecies struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Pokemon struct {
	EntryNumber    int            `json:"entry_number"`
	PokemonSpecies PokemonSpecies `json:"pokemon_species"`
}

type Descriptions struct {
	Description string `json:"description"`
	Language    struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"language"`
}

type Names struct {
	Name     string `json:"name"`
	Language struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"language"`
}

type Response struct {
	Descriptions []Descriptions `json:"descriptions"`
	Id           int            `json:"id"`
	IsMainSeries bool           `json:"is_main_series"`
	Name         string         `json:"name"`
	Names        []Names        `json:"names"`
	Pokemon      []Pokemon      `json:"pokemon_entries"`
	Region       struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"region"`
	VersionGroups []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"version_groups"`
}

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

	httpResponse, err_readall := ioutil.ReadAll(resp.Body)
	if err_readall != nil {
		panic(err_readall)
	}
	var response Response
	var pokemons []Pokemon

	json.Unmarshal(httpResponse, &response)
	pokemons = response.Pokemon
	for i := 0; i < len(pokemons); i++ {
		fmt.Println("Pokemon Entry Number: ", pokemons[i].EntryNumber)
		fmt.Println("Pokemon Name: ", pokemons[i].PokemonSpecies.Name)
		fmt.Println("Pokemon URL: ", pokemons[i].PokemonSpecies.Url)
		fmt.Println()
	}

}
