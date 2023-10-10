package views

import "pokemonGo/controllers"

func ImportPokemonMenu() {
	println("Importing...")
	controllers.ImportPokemonsToList()
	println("Imported!!")
}
