package main

import (
	"github.com/BotAldaris/pokedexcli/internal/pokeapi"
)

type config struct {
	next       *string
	previous   *string
	pokeClient pokeapi.Client
	name       *string
	pokemons   map[string]pokeapi.Pokemon
}
