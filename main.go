package main

import (
	"fmt"
	"phunsupgolang/pkg/routes"

	"github.com/gofiber/fiber/v2"
)

// init ทำก่อน main
func init() {
	//run พวก config
	fmt.Println("init")
}

func main(){
	fmt.Println("main")
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	routes.InitRoutes(app)
	
	app.Listen(":3000")
}