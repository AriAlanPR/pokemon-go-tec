package controllers

import (
	"gofiber/models"

	"github.com/gofiber/fiber"
)

func GetPokemons(c *fiber.Ctx) {
	c.JSON(models.GetAllPokemons())
}

func GetPokemon(c *fiber.Ctx) {
	id := c.Params("id")

	if pokemon := models.GetPokemon(id); pokemon.ID > 0 {
		c.JSON(pokemon)
		return
	}

	c.Send("Pokemon no encontrado")
}

func CreatePokemon(c *fiber.Ctx) {
	var pokemon models.Pokemon
	pokemon.Attack = 10
	pokemon.Defense = 10
	pokemon.HP = 100
	pokemon.Name = "Bulbasaur"
	pokemon.Type = "Grass"

	models.CreatePokemon(pokemon)
}
