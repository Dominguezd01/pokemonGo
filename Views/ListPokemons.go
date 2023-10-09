package views

import (
	"fmt"
	"pokemonGo/controllers"
)

func ListPokemons() {
	pokemons := controllers.GetPokemons()
	println("--------------------")
	for _, pokemon := range pokemons {
		fmt.Printf("\nId: %d \nName: %s \nTypes: %s %s \nHp max: %d \n\n", pokemon.ID, pokemon.Name, pokemon.Type1, pokemon.Type2, pokemon.HpMax)
		println("--------------------")
	}
}
