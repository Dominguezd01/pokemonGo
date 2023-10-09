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

func ExportPokemons(filePath string, listPokemons []structs.Pokemon) {
	if !fileExists("./PokemonExport.txt") {
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Error creating file:", err)
		}
		defer file.Close()
		var data string
		for i := 0; i < len(listPokemons); i++ {
			data += strconv.Itoa(listPokemons[i].ID) + "#" + listPokemons[i].Name + "#" + listPokemons[i].Type1 + "#" + listPokemons[i].Type2 + "#" + strconv.Itoa(listPokemons[i].HpMax) + "\n"
		}

		file.WriteString(data)
	} else {

	}

}
