package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpecies struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func main() {
	res, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		panic(err)
	}

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var response Response

	json.Unmarshal(result, &response)

	// fmt.Println(response.Pokemon)

	for i := 0; i < len(response.Pokemon); i++ {
		fmt.Println(response.Pokemon[i].Species.Name)
	}

}
