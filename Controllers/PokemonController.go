package controllers

import (
	"pokemonGo/structs"
)

var pokemons []structs.Pokemon

func CreatePokemon(id int, name string, type1 string, type2 string, hpMax int) structs.Pokemon {
	pokemonUser := structs.Pokemon{
		ID:    id,
		Name:  name,
		Type1: type1,
		Type2: type2,
		HpMax: hpMax,
	}

	return pokemonUser
}

func AddToList(newPokemon structs.Pokemon) {
	pokemons = append(pokemons, newPokemon)
}

func GetPokemons() []structs.Pokemon {
	return pokemons
}

/*
Gets the list and export it to a file called "PokemonExport.txt"
*/
func ExportPokemonToFile() {
	ExportPokemons(GetPokemons())
}

func ImportPokemonsToList() {

}
