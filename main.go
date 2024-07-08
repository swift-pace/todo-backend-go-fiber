package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

var books []Book

func main() {
	books = append(books, Book{Id: 1, Name: "New Game", Author: "NTK"})
	books = append(books, Book{Id: 2, Name: "Journal of Me", Author: "Ratri"})

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	app.Use(func(c *fiber.Ctx) error {
		start := time.Now()

		fmt.Printf(
			"URL = %s, Method = %s, Time = %s\n",
			c.OriginalURL(), c.Method(), start,
		)

		return c.Next()
	})

	app.Get("/books", handleGetBooks)
	app.Get("/books/:id", handleGetBook)
	app.Post("/books", handleCreateBook)
	app.Put("/books/:id", handleUpdateBook)
	app.Delete("/books/:id", handleDeleteBook)

	app.Post("/uploads", handleUploadImage)

	app.Get("/test", handleGetTestView)

	app.Get("/env", handleGetEnvironments)

	app.Post("/login", handleLogin)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}

func handleGetTestView(c *fiber.Ctx) error {
	return c.Render("test", fiber.Map{
		"Title": "Hello, world!",
	})
}

func handleGetEnvironments(c *fiber.Ctx) error {

	if secret, exists := os.LookupEnv("SECRET_KEY"); exists {
		return c.JSON(fiber.Map{
			"SECRET_KEY": secret,
		})
	}

	return c.JSON(fiber.Map{
		"SECRET_KEY": "default_secret_key",
	})
}

type User struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

func handleLogin(c *fiber.Ctx) error {
	return fiber.ErrUnauthorized
}
