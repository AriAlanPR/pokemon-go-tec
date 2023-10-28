package controllers

import (
	"fmt"
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

	if err := c.BodyParser(&pokemon); err != nil {
		return
	}

	fmt.Printf("Pokemon: %+v\n", pokemon)
	models.CreatePokemon(pokemon)
}

func UpdatePokemon(c *fiber.Ctx) {
	id := c.Params("id")
	var pokemon models.Pokemon

	if err := c.BodyParser(&pokemon); err == nil {
		if models.UpdatePokemon(id, pokemon) {
			c.Send("Pokemon actualizado satisfactoriamente")
			return
		}
	}

	c.Send("Error al actualizar Pokemon")
}

func DeletePokemon(c *fiber.Ctx) {
	id := c.Params("id")

	if models.DeletePokemon(id) {
		c.Send("Pokemon borrado satisfactoriamente")
		return
	}

	c.Send("Error al borrar Pokemon")
}
