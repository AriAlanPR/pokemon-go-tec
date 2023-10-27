package models

import (
	"gofiber/pokemon"
)

type Pokemon pokemon.Pokemon

var pokemons []Pokemon = []Pokemon{
	{
		Name: "Charmander", Type: "Fire", Attack: 5, HP: 100, Defense: 10,
	},
	{
		Name: "Bulbasaur", Type: "Grass", Attack: 10, HP: 100, Defense: 10,
	},
}

func GetAllPokemons() []Pokemon {
	return pokemons
}

func GetPokemon(name string) Pokemon {
	for _, p := range pokemons {
		if p.Name == name {
			return p
		}
	}

	return Pokemon{}
}

func CreatePokemon(pokemon Pokemon) {
	pokemons = append(pokemons, pokemon)
}
