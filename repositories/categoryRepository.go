package repositories

import (
	"bookstore-management-api/configs"
	"bookstore-management-api/models"
)

func GetAllCategories() ([]models.Category, error) {
	rows, err := configs.DB.Query(`
		SELECT id, name, created_at
		FROM categories
		ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category

	for rows.Next() {
		var c models.Category
		err := rows.Scan(&c.ID, &c.Name, &c.CreatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	return categories, nil
}

func GetCategoryByID(id int) (models.Category, error) {
	var c models.Category

	err := configs.DB.QueryRow(`
		SELECT id, name, created_at
		FROM categories
		WHERE id = $1
	`, id).Scan(&c.ID, &c.Name, &c.CreatedAt)

	return c, err
}

func CreateCategory(name string) (*models.Category, error) {
	query := `
		INSERT INTO categories (name)
		VALUES ($1)
		RETURNING id, name, created_at
	`

	var c models.Category

	err := configs.DB.QueryRow(query, name).Scan(
		&c.ID,
		&c.Name,
		&c.CreatedAt,
	)

	return &c, err
}

func DeleteCategory(id int) error {
	_, err := configs.DB.Exec(`
		DELETE FROM categories WHERE id = $1
	`, id)

	return err
}

func IsCategoryExists(categoryID int) (bool, error) {
	var exists bool
	err := configs.DB.QueryRow(`
		SELECT EXISTS (
			SELECT 1 FROM categories WHERE id = $1
		)
	`, categoryID).Scan(&exists)

	return exists, err
}
