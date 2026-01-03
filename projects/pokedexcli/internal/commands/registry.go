package commands

import "fmt"

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    Help,
		},
		"map": {
			Name:        "map",
			Description: "Displays the names of 20 locations areas in Pokemon world",
			Callback:    Map,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the names of 20 locations areas in Pokemon world backwards",
			Callback:    Mapb,
		},
		"explore": {
			Name:        "explore",
			Description: "Displays a list of all the Pokémon located there",
			Callback:    Explore,
		},
		"catch": {
			Name:        "catch",
			Description: "Catch a Pokemon",
			Callback:    Catch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect a Pokemon",
			Callback:    Inspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Display a list of all the names of the Pokemon the user has caught",
			Callback:    Pokedex,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    Exit,
		},
	}
}

func GetConfig() *Config {
	return &Config{
		pokedex: make(map[string]Pokemon),
	}
}

func GetApiUrl() string {
	return "https://pokeapi.co/api/v2/"
}

func PrintLocationNames(locations *Locations) {
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
}

func PrintPokemon(encounters PokemonEncounters) {
	fmt.Println("Found Pokemon:")
	for _, encounter := range encounters.PokemonEncounters {
		fmt.Println(" -", encounter.Pokemon.Name)
	}
}
