package views

import (
	"pokemonGo/controllers"
	"fmt"
)

func ListPokemons() {
	pokemons := controllers.GetPokemons()
	for _, pokemon := range pokemons {
		fmt.Printf("\nId: %d Name: %s Types: %s %s Hp max: %d \n\n ", pokemon.ID, pokemon.Name, pokemon.Type1, pokemon.Type2, pokemon.HpMax)
	}
}
