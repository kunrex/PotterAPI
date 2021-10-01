package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func Start() error {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
		CaseSensitive: true,
	})

	app.Static("/", "./static")


	for _, r := range routes {
		app.Get(r.Url, r.Callback)
	}

	app.Use(func(c* fiber.Ctx) error {
		return c.SendStatus(404)
	})

	err := app.Listen(":3000")

	if err != nil { return err }

	return nil
}