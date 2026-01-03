package commands

import (
	"fmt"
	"time"

	"github.com/fc1g/pokedexcli/internal/pokecache"
	"github.com/fc1g/pokedexcli/pkg"
)

func Explore(_ *Config, explore string) error {
	var encounters PokemonEncounters
	cache := pokecache.NewCache(5 * time.Second)

	url := GetApiUrl() + "location-area/" + explore

	if bytes, ok := cache.Get(url); ok {
		err := pkg.Unmarshal[PokemonEncounters](bytes, &encounters)
		if err != nil {
			return fmt.Errorf("no Pokemon's found in this area")
		}

		PrintPokemon(encounters)

		return nil
	}

	bytes, err := pkg.Get(url)
	cache.Add(url, bytes)
	err = pkg.Unmarshal[PokemonEncounters](bytes, &encounters)
	if err != nil {
		return fmt.Errorf("no Pokemon's found in this area")
	}

	PrintPokemon(encounters)

	return nil
}
