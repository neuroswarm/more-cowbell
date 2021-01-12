package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	// Create new Fiber instance
	cowbell := fiber.New(
		fiber.Config {
			CaseSensitive: true,
			StrictRouting: false,
			ServerHeader: "This needs more Cowbell!",
		} )

	cowbell.Get("cowbell", func(c *fiber.Ctx) error { return c.SendString("You are wise to visit me") } )
	cowbell.Get("/", func(c *fiber.Ctx) error { return c.SendString("Hello, World 👋!") } )


	log.Fatal(cowbell.Listen(":80"))
}