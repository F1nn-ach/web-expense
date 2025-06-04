package routes

import "github.com/gofiber/fiber/v2"

func GetLoginRoute(c *fiber.Ctx) error {
	return c.SendFile("./public/view/index.html")
}
