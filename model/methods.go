package model

import "database/sql"

// "../..database"
// "../database"

type BookModel struct {
	db *sql.DB
}

func (bookmodel BookModel) Update(book *Book) (int64, error) {
	query := `UPDATE tb_book SET title = $1, author = $2, year = $3 WHERE id = $4;`
	result, err := BookModel.db.Exec(query, book.Title, book.Author, book.Year, book.ID)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}
