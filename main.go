package main

import (
	"github.com/JulianKrmr/ratingsLambda/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON("Hello World")

	})
	app.Post("/ratings", func(ctx *fiber.Ctx) error {
		return api.HandleRatings(ctx)
	})

	err := app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}

}
