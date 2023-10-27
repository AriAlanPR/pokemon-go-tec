package controllers

import (
	"gofiber/models"

	"github.com/gofiber/fiber"
)

func GetPokemons(c *fiber.Ctx) {
	c.JSON(models.GetAllPokemons())
}

func GetPokemon(c *fiber.Ctx) {
	name := c.Params("name")

	c.JSON(models.GetPokemon(name))
}

func CreatePokemon(c *fiber.Ctx) {
	var pokemon models.Pokemon

	if err := c.BodyParser(&pokemon); err != nil {
		c.Status(503).Send(err)
	}

	models.CreatePokemon(pokemon)
}
