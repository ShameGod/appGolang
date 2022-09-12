package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json: "id"`
	Title    string `json: "title"`
	Author   string `json: "author"`
	Quantity int    `json: "quantity"`
}

var books = []book{
	{ID: "1", Title: "Antigone", Author: "Stephan KING", Quantity: 10},
	{ID: "2", Title: "Harry Potter", Author: "Rida KEJJI", Quantity: 10},
	{ID: "3", Title: "Hobbit", Author: "Safa KEJJI", Quantity: 10}}

func getBooks(c *gin.Context) { // here we do not do a deep copy of the parameters
	c.IndentedJSON(http.StatusOK, books)
}

func addBook(c *gin.Context) {
	var newBook book
	err := c.BindJSON(&newBook)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, "the book was added")
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/addbook", addBook)
	router.Run("localhost:8080")
}
