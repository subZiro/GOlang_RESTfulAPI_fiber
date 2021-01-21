package database

import (
	"database/sql"
	"fmt"
	"os"

	// "../model"
	"../model"

	_ "github.com/mattn/go-sqlite3"
)

// DB глобальная переменная для подключения к базе данны
var DB *sql.DB

// возвращает соединение к бд
func ConnectDB(name string) error {
	if !ExistsDB(name) {
		fmt.Println("DataBase not found")
		createDB(name)
		initDB(name)
	}
	return connect(name)
}

// подключение
func connect(name string) error {
	var err error
	DB, err = sql.Open("sqlite3", name)
	if err != nil {
		return err
	}
	// defer DB.Close()

	if err = DB.Ping(); err != nil {
		return err
	}

	fmt.Println("Connection Opened to Database")
	return nil
}

// возвращает true если существует файл name, иначе false
func ExistsDB(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// создание файла базы данных
func createDB(name string) {
	fmt.Println("Creating", name, "...")
	file, err := os.Create(name) // Create SQLite file
	checkErr(err)
	file.Close()
	fmt.Println(name, "created")
}

// создание таблиц и заполнение данными
func initDB(name string) {
	err := connect(name) // подключение
	checkErr(err)

	createTable() // создание таблицы
	insertData()  // заполнение данными

	// defer DB.Close() // закрытие соединения
}

// обработка ошибок
func checkErr(err error) {
	if err != nil {
		fmt.Printf("Fatal Error: %s\n", err)
	}
}

// создание таблицы tb_book
func createTable() {
	// SQL Statement for Create Table
	query := `
	CREATE TABLE IF NOT EXISTS tb_book (
		id integer NOT NULL PRIMARY KEY AUTOINCREMENT, 
		title TEXT, 
		author TEXT, 
		year TEXT
	);`

	fmt.Println("Create book table ...")
	statement, err := DB.Prepare(query) // Prepare statement
	checkErr(err)
	_, err = statement.Exec()
	checkErr(err)
	fmt.Println("book table created")
}

// вставка данных в таблицу tb_book
func insertData() {

	books := new(model.Books)
	books.Books = append(
		books.Books,
		model.Book{ID: 1, Title: "Golang pointers", Author: "Mr. Blue", Year: "2020"},
		model.Book{ID: 2, Title: "Go", Author: "Mr. Green", Year: "2021"},
		model.Book{ID: 3, Title: "Golang api routers", Author: "Mr. Apiner", Year: "2021"},
		model.Book{ID: 4, Title: "Golang rest ful api", Author: "Mr. Current", Year: "2019"},
		model.Book{ID: 5, Title: "Golang good test", Author: "Mr. Good", Year: "2018"},
	)

	fmt.Println("Inserting book record ...")
	query := `INSERT INTO tb_book(title, author, year) VALUES (?, ?, ?)`
	statement, err := DB.Prepare(query) // Prepare statement
	checkErr(err)

	for _, b := range books.Books {
		fmt.Println("Вставка :", b)
		_, err = statement.Exec(b.Title, b.Author, b.Year)
		checkErr(err)
	}

}
