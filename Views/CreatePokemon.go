package views

import "apitest/controllers"

func CreatePokemonView() {
	var name string
	var id int
	var type1 string
	var type2 string
	var hpMax int
	println("Pokemon ID")
	id = GetUserInt()
	println("Name of the pokemon")
	name = GetUserString()
	println("First Type")
	type1 = GetUserString()
	println("Second Type if not leave empty")
	type2 = GetUserString()
	println("Hp max")
	hpMax = GetUserInt()
	pokemonCreated := controllers.CreatePokemon(id, name, type1, type2, hpMax)
	controllers.AddToList(pokemonCreated)
}
