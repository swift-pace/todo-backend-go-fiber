package main

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Book struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

func handleGetBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func handleGetBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for _, book := range books {
		if book.Id == id {
			return c.JSON(book)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}

func handleCreateBook(c *fiber.Ctx) error {
	newBook := new(Book)
	if err := c.BodyParser(newBook); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if newBook.Name == "" {
		return c.Status(fiber.StatusBadRequest).SendString("book.name is required")
	}

	newBook.Id = books[len(books)-1].Id + 1
	books = append(books, *newBook)

	return c.JSON(newBook)
}

func handleUpdateBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	updateBook := new(Book)
	if err := c.BodyParser(updateBook); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if updateBook.Name == "" {
		return c.Status(fiber.StatusBadRequest).SendString("book.name is required")
	}

	for i := range books {
		if books[i].Id == bookId {
			books[i].Name = updateBook.Name
			books[i].Author = updateBook.Author

			return c.JSON(books[i])
		}
	}

	return c.Status(fiber.StatusNotFound).SendString("book not found")
}

func handleDeleteBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i := range books {
		if books[i].Id == bookId {
			books = append(books[:i], books[i+1:]...)

			return c.SendStatus(fiber.StatusNoContent)
		}
	}

	return c.Status(fiber.StatusNotFound).SendString("book not found")
}

func handleUploadImage(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = c.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}
