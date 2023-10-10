package controllers

import (
	"fmt"
	"os"
	"pokemonGo/structs"
	"strconv"
)

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return true
}

/*
Exports the pokemons array to a file so it could be saved
@param filePath path of the file to export
*/
func ExportPokemons(listPokemons []structs.Pokemon) {
	var data string
	for i := 0; i < len(listPokemons); i++ {
		data += strconv.Itoa(listPokemons[i].ID) + "#" + listPokemons[i].Name + "#" + listPokemons[i].Type1 + "#" + listPokemons[i].Type2 + "#" + strconv.Itoa(listPokemons[i].HpMax) + "\n"
	}
	if !fileExists("./PokemonExport.txt") {
		file, err := os.Create("./PokemonExport.txt")
		if err != nil {
			fmt.Println("Error creating file:", err)
		}

		defer file.Close()
		file.WriteString(data)

	} else {
		file, err := os.OpenFile("./PokemonExport.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 6666)

		if err != nil {
			fmt.Println("Error opening the file:", err)
		}
		defer file.Close()
		n, err := file.WriteString(data)
		if n == 0 {
			println("No bytes written")
		}
		if err == nil {
			println("And error occured")
		}
	}
}
