package main

import (
	"github.com/BotAldaris/pokedexcli/internal/pokeapi"
	"github.com/BotAldaris/pokedexcli/internal/pokecache"
)

type config struct {
	next       *string
	previous   *string
	pokeClient pokeapi.Client
	cache      pokecache.Cache
}
