package commands

import "fmt"

func Pokedex(config *Config, _ string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range config.pokedex {
		fmt.Println(" -", pokemon.Name)
	}
	return nil
}
