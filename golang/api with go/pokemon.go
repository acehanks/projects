package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
)
// A struct to map our Pokemon's Species which includes it's name
type PokemonSpecies struct {
	Name string `json:"name"`
	Url  string `json: "url"`
}

// A Pokemon Struct to map every pokemon to.
type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

// A Response struct to map the Entire Response
type PokeMain struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
	Filepath string
}

func main() {
	//fmt.Println("hello world\n")

	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto")
	if err != nil{
		log.Println(err)
	}
	
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(red))

	var responseObject PokeMain
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Name)

	fmt.Println(len(responseObject.Pokemon))

	for i:=0; i<len(responseObject.Pokemon); i++{
		//fmt.Println(responseObject.Pokemon)
		fmt.Println(responseObject.Pokemon[i].Species.Url)
	}

}