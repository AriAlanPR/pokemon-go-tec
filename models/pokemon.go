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

func UpdatePokemon(id string, newPokemonData Pokemon) bool {
	dbconn := database.DBConn
	//seek pokemon
	var pokemon Pokemon
	dbconn.First(&pokemon, id)

	//update pokemon
	if pokemon.ID > 0 {
		dbconn.Model(&pokemon).Updates(newPokemonData)
		return true
	}

	return false
}

func DeletePokemon(id string) bool {
	dbconn := database.DBConn

	var pokemon Pokemon
	dbconn.First(&pokemon, id)

	if pokemon.ID > 0 {
		dbconn.Delete(&pokemon)
		return true
	}

	return false
}
