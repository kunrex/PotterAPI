package api

import (
	"github.com/gofiber/fiber/v2"
	"os"
)

type Route struct {
	Url string
	Callback func(c *fiber.Ctx) error
}

var routes = []Route{
	{
		Url: "/",
		Callback: func(c *fiber.Ctx) error {
			return c.Render("index", fiber.Map{
				"Title": "Harry Potter API",
			})
		},
	},
	{
		Url: "/books",
		Callback: func(c *fiber.Ctx) error {
			book, err := os.ReadFile("./data/books.json")

			if err != nil { return c.SendStatus(500) }

			return c.SendString(string(book))
		},
	},
}