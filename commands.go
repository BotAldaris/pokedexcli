package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
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
		"explore": {
			name:        "explore",
			description: "Displays the pokemons in this area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a pokemon",
			callback:    commandCatch,
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
		return errors.New("you're on the first page")
	}
	locations, err := config.pokeClient.GetLocation(config.next)
	if err != nil {
		return err
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
	locations, err := config.pokeClient.GetLocation(config.previous)
	if err != nil {
		return err
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	config.next = &locations.Next
	config.previous = &locations.Previous
	return nil
}

func commandExplore(config *config) error {
	if config.name == nil {
		return errors.New("you should pass a name")
	}
	locationArea, err := config.pokeClient.GetPokemonsInArea(*config.name)
	if err != nil {
		return err
	}
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}

func commandCatch(config *config) error {
	if config.name == nil {
		return errors.New("you should pass a name")
	}
	fmt.Printf("Throwing a Pokeball at %s...", *config.name)
	pokemon, err := config.pokeClient.GetPokemon(*config.name)
	if err != nil {
		return err
	}
	i := rand.Intn(pokemon.BaseExperience/50 + 1)
	fmt.Println(i)
	if i == 0 {
		config.pokemons[pokemon.Name] = pokemon
		fmt.Printf("%s was caught!", *config.name)
	} else {
		fmt.Printf("%s escaped!", *config.name)
	}
	return nil
}
