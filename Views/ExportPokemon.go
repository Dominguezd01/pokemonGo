package views

import "pokemonGo/controllers"

func ExportPokemonMenu() {
	println("Exporting...")
	controllers.ExportPokemonToFile()
	println("Exported correctly")
}
