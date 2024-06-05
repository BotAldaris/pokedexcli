package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/BotAldaris/pokedexcli/internal/pokecache"
)

func repl() {
	scan := bufio.NewScanner(os.Stdin)
	commands_map := getCommands()
	next := "https://pokeapi.co/api/v2/location-area/"
	config := config{
		next:     &next,
		previous: nil,
		cache:    pokecache.NewCache(time.Duration(5) * time.Second),
	}
	for {
		print("Pokedex > ")
		scan.Scan()
		text := scan.Text()
		parsedText := parseInput(text)
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
