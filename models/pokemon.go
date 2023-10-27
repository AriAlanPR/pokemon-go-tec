package models

import (
	"gofiber/database"
	"gofiber/pokemon"
)

type Pokemon pokemon.Pokemon

func GetAllPokemons() []Pokemon {
	dbconn := database.DBConn
	var pokemons []Pokemon
	dbconn.Find(&pokemons)
	return pokemons
}

func GetPokemon(id string) Pokemon {
	dbconn := database.DBConn
	var pokemon Pokemon
	dbconn.Find(&pokemon, id)

	return pokemon
}

func CreatePokemon(pokemon Pokemon) Pokemon {
	dbconn := database.DBConn

	dbconn.Create(&pokemon)
	return pokemon
}
