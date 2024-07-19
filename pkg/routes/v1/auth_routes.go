package v1

import (
	"phunsupgolang/app/controller"

	"github.com/gofiber/fiber/v2"
)

func InitAuthRoutes(app fiber.Router){
	auth := app.Group("/auth")
	ctl:= controller.NewAuthController()
	auth.Get("/login",clt.LoginHandler)
}