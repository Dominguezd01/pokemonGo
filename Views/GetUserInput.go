package views

import (
	"bufio"
	"os"
	"strconv"
)

func GetUserString() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userInput := scanner.Text()
	return userInput
}

func GetUserInt() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userInput := scanner.Text()
	parsedInt, err := strconv.Atoi(userInput)
	if err != nil {
		println(err)
	}
	return parsedInt

}
