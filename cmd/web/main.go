package main

import (
	"log"
	"os"

	"github.com/F1nn-ach/web-expense/handler"
	"github.com/F1nn-ach/web-expense/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	app.Static("/", "./public/view")
	app.Get("/login", routes.GetLoginRoute)
	app.Post("/secure-action", handler.SecureAction)
	log.Println("🚨 ROUTE /login UPDATED")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Println("Have no port")
	}
	app.Listen(":" + port)
}
