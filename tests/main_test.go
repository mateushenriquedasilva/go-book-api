package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
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
	api.DB.Create(&book)
	return book
}

func TestCreateBook(t *testing.T) {
	setupTestDB()

	router := gin.Default()
	router.POST("/book", api.CreateBook)

	book := api.Book{
		Title:  "Demo Book name",
		Author: "Demo Author name",
		Year:   2021,
	}

	jsonValue, _ := json.Marshal(book)
	req, _ := http.NewRequest("POST", "/book", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, status)
	}

	var response api.JsonResponse
	json.NewDecoder(w.Body).Decode(&response)

	if response.Data == nil {
		t.Error("Expected book data, got nil")
	}
}

func TestGetBooks(t *testing.T) {
	setupTestDB()
	addBook()

	router := gin.Default()
	router.GET("/books", api.GetBooks)

	req, _ := http.NewRequest("GET", "/books", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}

	var response api.JsonResponse
	json.NewDecoder(w.Body).Decode(&response)

	if len(response.Data.([]interface{})) == 0 {
		t.Errorf("Expected non-empty books list")
	}
}

func TestGetBook(t *testing.T) {
	setupTestDB()
	book := addBook()
	router := gin.Default()
	router.GET("/book/:id", api.GetBook)

	req, _ := http.NewRequest("GET", "/book/"+strconv.Itoa(int(book.ID)), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}

	var response api.JsonResponse
	json.NewDecoder(w.Body).Decode(&response)

	if response.Data == nil || response.Data.(map[string]interface{})["id"] != float64(book.ID) {
		t.Errorf("Expected book ID %d, got nil or wrong ID", book.ID)
	}
}
