package repositories

import (
	"bookstore-management-api/configs"
	"bookstore-management-api/models"
	"database/sql"
)

// GET ALL
func GetBooks() ([]models.Book, error) {
	rows, err := configs.DB.Query(`
		SELECT id, title, description, image_url, release_year,
		       price, total_page, thickness, category_id,
		       created_at, created_by
		FROM books
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book

	for rows.Next() {
		var book models.Book
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Description,
			&book.ImageURL,
			&book.ReleaseYear,
			&book.Price,
			&book.TotalPage,
			&book.Thickness,
			&book.CategoryID,
			&book.CreatedAt,
			&book.CreatedBy,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func GetBookByID(id int) (*models.Book, error) {
	row := configs.DB.QueryRow(`
		SELECT id, title, description, image_url, release_year,
		       price, total_page, thickness, category_id,
		       created_at, created_by
		FROM books WHERE id = $1
	`, id)

	var book models.Book
	err := row.Scan(
		&book.ID,
		&book.Title,
		&book.Description,
		&book.ImageURL,
		&book.ReleaseYear,
		&book.Price,
		&book.TotalPage,
		&book.Thickness,
		&book.CategoryID,
		&book.CreatedAt,
		&book.CreatedBy,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &book, nil
}

func CreateBook(book models.Book) (*models.Book, error) {
	query := `
	INSERT INTO books (
	title, description, image_url, release_year,
	price, total_page, thickness, category_id, created_by
	) 
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
	RETURNING id, title, description, image_url, release_year, 
	price, total_page, thickness, category_id, 
	created_at, created_by`

	var createdBook models.Book

	err := configs.DB.QueryRow(query,
		book.Title,
		book.Description,
		book.ImageURL,
		book.ReleaseYear,
		book.Price,
		book.TotalPage,
		book.Thickness,
		book.CategoryID,
		book.CreatedBy,
	).Scan(
		&createdBook.ID,
		&createdBook.Title,
		&createdBook.Description,
		&createdBook.ImageURL,
		&createdBook.ReleaseYear,
		&createdBook.Price,
		&createdBook.TotalPage,
		&createdBook.Thickness,
		&createdBook.CategoryID,
		&createdBook.CreatedAt,
		&createdBook.CreatedBy,
	)

	return &createdBook, err
}

func DeleteBook(id int) (bool, error) {
	res, err := configs.DB.Exec(`DELETE FROM books WHERE id = $1`, id)
	if err != nil {
		return false, err
	}

	affected, _ := res.RowsAffected()
	return affected > 0, nil
}

func GetBooksByCategoryID(categoryID int) ([]models.Book, error) {
	rows, err := configs.DB.Query(`
		SELECT id, title, description, image_url, release_year,
		       price, total_page, thickness, category_id, created_at
		FROM books
		WHERE category_id = $1
	`, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book

	for rows.Next() {
		var book models.Book
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Description,
			&book.ImageURL,
			&book.ReleaseYear,
			&book.Price,
			&book.TotalPage,
			&book.Thickness,
			&book.CategoryID,
			&book.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}
