package commands

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fc1g/pokedexcli/internal/pokecache"
	"github.com/fc1g/pokedexcli/pkg"
)

func catchChance(baseExp int) float64 {
	chance := 0.60 - float64(baseExp)*0.002

	if chance < 0.10 {
		chance = 0.10
	}
	if chance > 0.60 {
		chance = 0.60
	}

	return chance
}

func tryCatch(baseExp int) bool {
	chance := catchChance(baseExp)
	roll := rand.Intn(100)
	return roll < int(chance*100.0)
}

func Catch(config *Config, pokemonName string) error {
	var pokemon Pokemon
	cache := pokecache.NewCache(5 * time.Second)

	url := GetApiUrl() + "pokemon/" + pokemonName

	if bytes, ok := cache.Get(url); ok {
		err := pkg.Unmarshal[Pokemon](bytes, &pokemon)
		if err != nil {
			return fmt.Errorf("pokemon not found")
		}

		rand.NewSource(time.Now().UnixNano())
		if tryCatch(pokemon.BaseExperience) {
			fmt.Printf("You caught %s!\n", pokemon.Name)
			config.pokedex[pokemonName] = pokemon
			fmt.Println("You may now inspect it with the inspect command.")
		} else {
			fmt.Printf("%s escaped!\n", pokemon.Name)
		}

		return nil
	}

	bytes, err := pkg.Get(url)
	cache.Add(url, bytes)
	err = pkg.Unmarshal[Pokemon](bytes, &pokemon)
	if err != nil {
		return fmt.Errorf("pokemon not found")
	}

	rand.NewSource(time.Now().UnixNano())
	if tryCatch(pokemon.BaseExperience) {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		config.pokedex[pokemonName] = pokemon
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
