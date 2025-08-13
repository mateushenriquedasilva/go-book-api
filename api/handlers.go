package api

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	dns := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := DB.AutoMigrate(&Book{}); err != nil {
		log.Fatal("Failed to migrate schema:", err)
	}
}

func CreateBook(c *gin.Context) {
	var book Book

	if err := c.ShouldBindJSON(&book); err != nil {
		ResponseJSON(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}
	DB.Create(&book)
	ResponseJSON(c, http.StatusCreated, "Book created successfully", book)
}

func GetBooks(c *gin.Context) {
	var books []Book
	DB.Find(&books)
	ResponseJSON(c, http.StatusOK, "Books retrieved successfully", books)
}
