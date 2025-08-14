package tests

import (
	"github.com/mateushenriquedasilva/go-book-api/api"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() {
	var err error
	api.DB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect test database")
	}
	api.DB.AutoMigrate(&api.Book{})
}

func addBook() api.Book {
	book := api.Book{Title: "Go Programming", Author: "John Doe", Year: 2023}
	api.DB.Create(book)
	return book
}
