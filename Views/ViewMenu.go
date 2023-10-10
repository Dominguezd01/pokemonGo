package views

import (
	"fmt"
	"os"
)

func MainMenu() {
	var options [5]string
	options[0] = "Exit"
	options[1] = "Create new Pokemon"
	options[2] = "List your Pokemons"
	options[3] = "Import your Pokemons"
	options[4] = "Export your Pokemons"
	println("\n")
	for i := 0; i < len(options); i++ {
		fmt.Printf("%d. %s \n", i, options[i])
	}

	var option int = GetUserInt()
	switch option {
	case 0:
		println("Exiting...")
		os.Exit(0)
		break
	case 1:
		CreatePokemonView()
		MainMenu()
		break
	case 2:
		ListPokemons()
		MainMenu()
		break
	case 4:
		ExportPokemonMenu()
		MainMenu()
		break
	default:
		println("Wrong option")
		MainMenu()
	}
}
