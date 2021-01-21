package model

// Book struct
type Book struct {
	ID     int    `json: "id"`
	Title  string `json: "title"`
	Author string `json: "author"`
	Year   string `json: "year"`
}

// Books struct
type Books struct {
	Books []Book `json: "books"`
}
