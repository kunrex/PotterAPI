package api

import (
	"github.com/gofiber/fiber/v2"
)

type Route struct {
	Url string
	Callback func(c *fiber.Ctx) error
}

var routes = []Route{
	{
		Url: "/books",
		Callback: func(c *fiber.Ctx) error {
			book, err := ReadFile("./data/books.json")

			if err != nil { return c.SendStatus(500) }

			return c.SendString(string(book))
		},
	},
	{
		Url: "/houses",
		Callback: func(c *fiber.Ctx) error {
			houses, err := ReadFile("./data/houses.json")

			if err != nil {
				return c.SendStatus(500)
			}

			return c.SendString(string(houses))
		},
	},
	{
		Url: "/horcruxes",
		Callback: func(c *fiber.Ctx) error {
			horcruxes, err := ReadFile("./data/horcruxes.json")

			if err != nil { return c.SendStatus(500) }

			return c.SendString(string(horcruxes))
		},
	},
	{
		Url: "/schools",
		Callback: func(c *fiber.Ctx) error {
			schools, err := ReadFile("./data/schools.json")

			if err != nil { return c.SendStatus(500) }

			return c.SendString(string(schools))
		},
	},
	{
		Url: "/subjects",
		Callback: func(c *fiber.Ctx) error {
			return c.SendString("Working on it!")
		},
	},
}