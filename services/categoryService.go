package services

import (
	"errors"

	"bookstore-management-api/models"
	"bookstore-management-api/repositories"
)

func GetCategories() ([]models.Category, error) {
	return repositories.GetAllCategories()
}

func GetCategoryByID(id int) (models.Category, error) {
	return repositories.GetCategoryByID(id)
}

func CreateCategory(name string) (*models.Category, error) {
	if name == "" {
		return nil, errors.New("category name is required")
	}

	return repositories.CreateCategory(name)
}

func DeleteCategory(id int) error {
	return repositories.DeleteCategory(id)
}

func GetBooksByCategory(categoryID int) (interface{}, error) {
	_, err := repositories.GetCategoryByID(categoryID)
	if err != nil {
		return nil, errors.New("category not found")
	}

	books, err := repositories.GetBooksByCategoryID(categoryID)
	if err != nil {
		return nil, err
	}

	return books, nil
}
