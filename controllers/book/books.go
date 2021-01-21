package book

import (
	"database/sql"
	"fmt"

	// "../../database"
	"../../database"

	// "../../model"
	"../../model"

	"github.com/gofiber/fiber"
)

// получение массива всех книг из бд
func GetBooks(c *fiber.Ctx) error {
	rows, err := database.DB.Query("SELECT * FROM tb_book;")
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err})
	}
	defer rows.Close()

	result := model.Books{}
	for rows.Next() { // итерации по строкам ответа запроса
		b := model.Book{}
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Year); err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"error":   err})
		}

		result.Books = append(result.Books, b) // добавление книги в массив книг
	}

	// 200, корректный результат
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "Books returnet",
		"data":    result.Books})
}

// получение книги по id
func GetBook(c *fiber.Ctx) error {

	// извлечение переданного параметра запроса
	id := c.Params("id")
	// запрос получения 1 строки по фильтру id
	row := database.DB.QueryRow("SELECT * FROM tb_book WHERE id = $1", id)

	b := new(model.Book)
	// кейс возможных ошибок результата запроса
	switch err := row.Scan(&b.ID, &b.Title, &b.Author, &b.Year); err {
	case sql.ErrNoRows: // не получено строк
		fmt.Println("No rows were returned!")
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err})
	case nil: // нет ошибок
		fmt.Println(b.ID, b.Title, b.Author, b.Year)
	default: // дефолтно ошибка и возврат 500
		// panic(err)
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err})
	}

	// 200, корректный результат
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "Book returnet",
		"data":    b})
}

// метод добавления новой книги
func AddBook(c *fiber.Ctx) error {

	// создание структуры книги
	b := new(model.Book)
	//  парсинг переданного тела post запроса
	if err := c.BodyParser(b); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err})
	}

	// запрос вставки новой строки в таблицк tb_book
	query := "INSERT INTO tb_book (title, author, year) VALUES ($1, $2, $3);"
	res, err := database.DB.Exec(query, b.Title, b.Author, b.Year)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err})
	}

	// 200, корректный результат
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"data":    res,
		"message": "Book successfully created"})
}

// метод обновления книги по id
func UpdateBook(c *fiber.Ctx) error {

	// создание структуры книги
	b := new(model.Book)
	//  парсинг переданного тела post запроса
	if err := c.BodyParser(b); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err})
	}

	// // получение номера переданной книги
	id := c.Params("id")
	// запром удаления киниги по фильтру id
	query := `UPDATE tb_book SET title = $1, author = $2, year = $3 WHERE id = $4;`
	res, err := database.DB.Exec(query, b.Title, b.Author, b.Year, id)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err})
	}

	// 200, корректный результат
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "book is updated successfully",
		"data":    res})
}

// метод удаления книги по переданному параметру id
func DeleteBook(c *fiber.Ctx) error {

	// получение номера переданной книги
	id := c.Params("id")
	// запром удаления киниги по фильтру id
	res, err := database.DB.Exec("DELETE FROM tb_book WHERE id = $1", id)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err})
	}

	// return 200 in JSON format
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "book is deleted successfully",
		"data":    res})
}
