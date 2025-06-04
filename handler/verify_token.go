package handler

import (
	"context"
	"log"

	"github.com/F1nn-ach/web-expense/internal"
	"github.com/gofiber/fiber/v2"
)

func SecureAction(c *fiber.Ctx) error {
	idToken := c.Get("Authorization")
	if idToken == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Missing ID Token")
	}

	authClient, err := internal.GetAuthClient()
	if err != nil {
		log.Println("Auth client error:", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	token, err := authClient.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		log.Println("Invalid token:", err)
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	uid := token.UID
	log.Println("🎉 Authenticated user:", uid)

	c.Cookie(&fiber.Cookie{
		Name: ,
	})
	return nil
}
