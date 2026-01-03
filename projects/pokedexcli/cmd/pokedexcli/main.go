package main

import (
	"bufio"
	"fmt"
	"os"

	_commands "github.com/fc1g/pokedexcli/internal/commands"
	"github.com/fc1g/pokedexcli/pkg"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := _commands.GetCommands()
	config := _commands.GetConfig()

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := pkg.CleanInput(scanner.Text())
		commandName := text[0]
		var extra string

		command, exists := commands[commandName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		switch commandName {
		case "explore":
			extra = text[1]
			fmt.Println("Exploring", extra+"...")
		case "catch":
			extra = text[1]
			fmt.Println("Throwing a Pokeball at", extra+"...")
		case "inspect":
			extra = text[1]
		default:
			extra = ""
		}

		err := command.Callback(config, extra)
		if err != nil {
			fmt.Println(err)
		}
	}
}
