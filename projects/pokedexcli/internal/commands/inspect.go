package commands

import "fmt"

func Inspect(config *Config, pokemonName string) error {
	if pokemon, ok := config.pokedex[pokemonName]; !ok {
		fmt.Println("you have not caught that pokemon")
	} else {
		fmt.Println("Name:", pokemon.Name)
		fmt.Println("Height:", pokemon.Height)
		fmt.Println("Weight:", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Println(" -"+stat.Stat.Name+":", stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, type_ := range pokemon.Types {
			fmt.Println(" -", type_.Type.Name)
		}
	}
	return nil
}
