package routes

import (
	"../controllers/book"
	"github.com/gofiber/fiber"
)

func BooksRoute(route fiber.Router) {
	route.Get("/", book.GetBooks)
	route.Get("/:id", book.GetBook)
	route.Post("/", book.AddBook)
	route.Put("/:id", book.UpdateBook)
	route.Delete("/:id", book.DeleteBook)
}
