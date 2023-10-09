package views

import "apitest/controllers"

func ExportPokemonMenu() {
	println("Exporting...")
	controllers.ExportPokemonToFile()
	println("Exported correctly")
}
