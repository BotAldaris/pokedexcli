package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/BotAldaris/pokedexcli/internal/pokeapi"
)

func repl() {
	scan := bufio.NewScanner(os.Stdin)
	commands_map := getCommands()
	next := "https://pokeapi.co/api/v2/location-area/"
	config := config{
		next:       &next,
		previous:   nil,
		pokeClient: pokeapi.NewClient(5*time.Second, time.Minute*5),
		pokemons:   make(map[string]pokeapi.Pokemon),
	}
	for {
		print("Pokedex > ")
		scan.Scan()
		text := scan.Text()
		parsedText := parseInput(text)
		if len(parsedText) == 0 {
			continue
		}
		if len(parsedText) > 1 {
			config.name = &parsedText[1]
		}
		command, ok := commands_map[parsedText[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(&config)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func parseInput(text string) []string {
	lower := strings.ToLower(text)
	return strings.Fields(lower)
}
