package commands

import "fmt"

func Help(_ *Config, _ string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	commands := GetCommands()
	for _, command := range commands {
		fmt.Println(command.Name, ":", command.Description)
	}

	return nil
}
