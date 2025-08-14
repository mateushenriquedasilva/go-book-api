package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mateushenriquedasilva/go-book-api/api"
)

func main() {
	api.InitDB()
	r := gin.Default()

	//routes
	r.GET("/books", api.GetBooks)
	r.POST("/books", api.CreateBook)
	r.GET("/book/:id", api.GetBook)
	r.PUT("/book/:id", api.UpdateBook)
	r.DELETE("/book/:id", api.DeleteBook)

	r.Run(":8080")
}
