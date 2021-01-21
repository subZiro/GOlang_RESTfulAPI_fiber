package main

import (
	"github.com/gofiber/fiber"

	// "/controllers/books"
	// "/controllers/books"
	// "./routes"
	"./database"

	"./routes"
)

func setupRoutes(app *fiber.App) {
	// moved from main method

	// api group
	api := app.Group("/api")

	// give response when at /api
	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the api endpoint 😉",
		})
	})

	// register books endpoints
	routes.BooksRoute(api.Group("/books"))
}

func main() {
	var err error
	if err = database.ConnectDB("database/sqlite3.db"); err != nil {
		panic(err)
	}

	app := fiber.New()
	setupRoutes(app)
	err = app.Listen(":8080")
	// handle error
	if err != nil {
		panic(err)
	}
	defer database.DB.Close()
}
