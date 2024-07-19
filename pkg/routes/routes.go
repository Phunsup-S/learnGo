package routes

import (
	v1 "phunsupgolang/pkg/routes/V1"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App){
	root := app.Group("/api")
	root.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Art")
	})
	v1.InitAuthRoutes(root)
	
}