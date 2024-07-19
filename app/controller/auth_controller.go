package controller

import "github.com/gofiber/fiber/v2"

type IAuthController interface {
 LoginHandler(c *fiber.Ctx) error
}

type authController struct {
 Name string
}

func NewAuthController() IAuthController {
 return &authController{
  Name: "auth",
 }
}

// LoginHandler implements IAuthController.
func (a *authController) LoginHandler(c *fiber.Ctx) error {
 return c.SendString(a.Name)
}

func (a *authController) LoginHandlerB(c *fiber.Ctx) error {
	return c.SendString(a.Name)
   }