package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type book struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

type errorMessage struct {
	Message string `json:"error"`
}

var bookNotFoundErrorMessage = errorMessage{"We could not find the book you requested"}
var outOfBookErrorMessage = errorMessage{"Unfortunately, we have run out of the book you requested. Please try again later."}
var bookAlreadyExistErrorMessage = errorMessage{"Id for the book already exists. Try a different id."}

var books = []book{
	{Id: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{Id: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{Id: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/book/:id", getBookById)
	router.POST("/book", createBook)
	router.PUT("book/checkout/:id", checkoutBook)
	router.PUT("book/return/:id", returnBook)
	router.DELETE("admin/book/:id", deleteBook)
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func findBookIndexById(id string) int {
	for i, b := range books {
		if b.Id == id {
			return i
		}
	}
	return -1
}

func findBookById(id string) *book {
	for i, b := range books {
		if b.Id == id {
			return &books[i]
		}
	}
	return nil
}

func getBookById(c *gin.Context) {
	var bkPtr = findBookById(c.Param("id"))

	if bkPtr == nil {
		c.IndentedJSON(http.StatusBadRequest, bookNotFoundErrorMessage)
	} else {
		c.IndentedJSON(http.StatusOK, *bkPtr)
	}
}

func createBook(c *gin.Context) {
	var newBook book

	var err = c.BindJSON(&newBook)
	if err != nil {
		return
	}

	if findBookIndexById(newBook.Id) == -1 {
		c.IndentedJSON(http.StatusBadRequest, bookAlreadyExistErrorMessage)
	} else {
		books = append(books, newBook)
		c.IndentedJSON(http.StatusOK, books)
	}
}

func checkoutBook(c *gin.Context) {
	var index = findBookIndexById(c.Param("id"))

	if index == -1 {
		c.IndentedJSON(http.StatusBadRequest, bookNotFoundErrorMessage)
	} else {
		var bk = books[index]
		if bk.Quantity < 1 {
			c.IndentedJSON(http.StatusAccepted, outOfBookErrorMessage)
		} else {
			bk.Quantity -= 1
			books[index] = bk
			c.IndentedJSON(http.StatusOK, bk)
		}
	}
}

func returnBook(c *gin.Context) {
	var index = findBookIndexById(c.Param("id"))

	if index == -1 {
		c.IndentedJSON(http.StatusBadRequest, bookNotFoundErrorMessage)
	} else {
		var bk = books[index]
		bk.Quantity += 1
		books[index] = bk
		c.IndentedJSON(http.StatusOK, bk)
	}
}

func deleteBook(c *gin.Context) {
	var index = findBookIndexById(c.Param("id"))

	if index == -1 {
		c.IndentedJSON(http.StatusBadRequest, bookNotFoundErrorMessage)
	} else {
		books = append(books[:index], books[index+1:]...)
		c.IndentedJSON(http.StatusOK, books)
	}
}
