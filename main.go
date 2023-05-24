package main

import (
	"errors"
	"net/http"

	/*"errors"*/
	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Genre    string `json:"genre"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Five Tuesdays in Winter", Author: "Lily King", Genre: "Romance Novel", Quantity: 15},
	{ID: "2", Title: "Trust", Author: "Hernan Diaz", Genre: "Historical Fiction", Quantity: 13},
	{ID: "3", Title: "Now I am here", Author: "Chidi Ebere", Genre: "Dystopian Fiction", Quantity: 9},
	{ID: "4", Title: "Maps of Our Spectacular Bodies", Author: "Maddie Mortimer", Genre: "Domestic Fiction", Quantity: 7},
	{ID: "5", Title: "What You Need From The Night", Author: "Laurent Petitmangin", Genre: "Political Fiction", Quantity: 5},
	{ID: "6", Title: "Devotion", Author: "Hannah Kent", Genre: "Literary Fiction", Quantity: 3},
	{ID: "7", Title: "Western Lane", Author: "Chetna Maroo", Genre: "Coming-of-age", Quantity: 2},
	{ID: "8", Title: "The Cat Who Saved Books", Author: "Sosuke Natsukawa", Genre: "Fantasy Fiction", Quantity: 1},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func bookByID(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := getBookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}
	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)

}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := getBookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not availble"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func getBookByID(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil

		}
	}
	return nil, errors.New("that book wasn't found")
}

func addBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookByID)
	router.POST("/books", addBook)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("return", returnBook)
	router.Run("localhost:8080")
}
