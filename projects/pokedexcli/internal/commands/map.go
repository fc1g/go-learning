package commands

import (
	"fmt"
	"time"

	"github.com/fc1g/pokedexcli/internal/pokecache"
	"github.com/fc1g/pokedexcli/pkg"
)

func handleLocationsResponse(bytes []byte, locations *Locations, config *Config) error {
	err := pkg.Unmarshal[*Locations](bytes, &locations)
	if err != nil {
		return err
	}

	PrintLocationNames(locations)

	config.nextLocationURL = locations.Next
	config.previousLocationURL = locations.Previous

	return nil
}

func Map(config *Config, _ string) error {
	var locations Locations
	cache := pokecache.NewCache(5 * time.Second)

	url := GetApiUrl() + "location-area/"
	if config.nextLocationURL != "" {
		url = config.nextLocationURL
	}

	if bytes, ok := cache.Get(url); ok {
		err := handleLocationsResponse(bytes, &locations, config)
		if err != nil {
			return err
		}
	}

	bytes, err := pkg.Get(url)
	cache.Add(url, bytes)
	err = handleLocationsResponse(bytes, &locations, config)
	if err != nil {
		return err
	}

	return nil
}

func Mapb(config *Config, _ string) error {
	if config.previousLocationURL == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	var locations Locations
	cache := pokecache.NewCache(5 * time.Second)

	url := GetApiUrl() + "location-area/"
	if config.previousLocationURL != "" {
		url = config.previousLocationURL
	}

	if bytes, ok := cache.Get(url); ok {
		err := handleLocationsResponse(bytes, &locations, config)
		if err != nil {
			return err
		}

		return nil
	}

	bytes, err := pkg.Get(url)
	cache.Add(url, bytes)
	err = handleLocationsResponse(bytes, &locations, config)
	if err != nil {
		return err
	}

	return nil
}
