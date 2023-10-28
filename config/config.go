package config

import (
	"gofiber/controllers"

	"github.com/gofiber/fiber"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	app.Get("/pokemons", controllers.GetPokemons)
	app.Get("/pokemons/:id", controllers.GetPokemon)
	app.Post("/pokemons", controllers.CreatePokemon)
	app.Put("/pokemons/:id", controllers.UpdatePokemon)
	app.Delete("/pokemons/:id", controllers.DeletePokemon)
}
