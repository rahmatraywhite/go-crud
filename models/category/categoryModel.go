package categoryModel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

func GetAll() []entities.Category {
	rows, err := config.DB.Query(`SELECT * FROM categories`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var categories []entities.Category

	for rows.Next() {
		var category entities.Category
		if err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpddatedAt); err != nil {
			panic(err)
		}

		categories = append(categories, category)
	}

	return categories
}

func Create(category entities.Category) bool {
	result, err := config.DB.Exec(`
		INSERT INTO categories (name, created_at, updated_at) VALUES (?, ?, ?)
	`, category.Name, category.CreatedAt, category.UpddatedAt)

	if err != nil {
		panic(err)
	}

	lastInsterId, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	return lastInsterId > 0
}
