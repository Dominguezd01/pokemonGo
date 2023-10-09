package views

import "pokemonGo/controllers"

func ExportPokemonMenu() {
	println("Exporting...")
	defer controllers.ExportPokemonToFile()
	println("Exported correctly")
}
