package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/BotAldaris/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations.",
			callback:    commandMapb,
		},
	}
}

func commandHelp(config *config) error {
	commands := getCommands()
	println("Welcome to the Pokedex!")
	println("Usage:")
	println()
	for _, v := range commands {
		println(v.name, ":", v.description)
	}
	println()
	return nil
}

func commandExit(config *config) error {
	os.Exit(0)
	return nil
}

func commandMap(config *config) error {
	if config.next == nil {
		return errors.New("you're on the last page")
	}
	body, ok := config.cache.Get(*config.next)
	locations := pokeapi.Locations{}
	if ok {
		err_json := json.Unmarshal(body, &locations)
		println("Getting from cache")
		if err_json != nil {
			return err_json
		}
	} else {
		l, err := config.pokeClient.GetLocation(config.next)
		locations = l
		if err != nil {
			return err
		}
		location_bytes, err := json.Marshal(locations)
		if err != nil {
			return err
		}
		config.cache.Add(*config.next, location_bytes)
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	config.next = &locations.Next
	config.previous = &locations.Previous
	return nil
}
func commandMapb(config *config) error {
	if config.previous == nil {
		return errors.New("you're on the first page")
	}
	body, ok := config.cache.Get(*config.previous)
	locations := pokeapi.Locations{}
	if ok {
		println("Getting from cache")
		err_json := json.Unmarshal(body, &locations)
		if err_json != nil {
			return err_json
		}
	} else {
		l, err := config.pokeClient.GetLocation(config.previous)
		locations = l
		if err != nil {
			return err
		}
		location_bytes, err := json.Marshal(locations)
		if err != nil {
			return err
		}
		config.cache.Add(*config.previous, location_bytes)
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	config.next = &locations.Next
	config.previous = &locations.Previous
	return nil
}
