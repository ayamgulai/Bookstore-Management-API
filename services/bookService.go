package services

import (
	"errors"

	"bookstore-management-api/models"
	"bookstore-management-api/repositories"
)

func GetBooks() ([]models.Book, error) {
	return repositories.GetBooks()
}

func GetBookByID(id int) (*models.Book, error) {
	book, err := repositories.GetBookByID(id)
	if err != nil {
		return nil, err
	}

	if book == nil {
		return nil, errors.New("book not found")
	}

	return book, nil
}

func CreateBook(book models.Book) (*models.Book, error) {
	if book.Title == "" {
		return nil, errors.New("title is required")
	}

	exists, err := repositories.IsCategoryExists(book.CategoryID)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, errors.New("category not found")
	}
	// VALIDASI RELEASE YEAR
	if book.ReleaseYear < 1980 || book.ReleaseYear > 2024 {
		return nil, errors.New("release_year must be between 1980 and 2024")
	}

	// KONVERSI THICKNESS
	if book.TotalPage > 100 {
		book.Thickness = "tebal"
	} else {
		book.Thickness = "tipis"
	}

	return repositories.CreateBook(book)
}

func DeleteBook(id int) error {
	deleted, err := repositories.DeleteBook(id)
	if err != nil {
		return err
	}

	if !deleted {
		return errors.New("book not found")
	}

	return nil
}
