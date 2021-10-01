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
		Url: "/books",
		Callback: func(c *fiber.Ctx) error {
			book, err := os.ReadFile("./data/books.json")

			if err != nil { return c.SendStatus(500) }

			return c.SendString(string(book))
		},
	},
	{
		Url: "/houses",
		Callback: func(c *fiber.Ctx) error {
			houses, err := os.ReadFile("./data/houses.json")

			if err != nil {
				return c.SendStatus(500)
			}

			return c.SendString(string(houses))
		},
	},
	{
		Url: "/horcruxes",
		Callback: func(c *fiber.Ctx) error {
			horcruxes, err := os.ReadFile("./data/horcruxes.json")

			if err != nil { return c.SendStatus(500) }

			return c.SendString(string(horcruxes))
		},
	},
	{
		Url: "/schools",
		Callback: func(c *fiber.Ctx) error {
			return c.SendString("Workin on it")
		},
	},
}