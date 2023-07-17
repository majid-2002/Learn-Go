package controllers

import (
	"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/majid-2002/go-bookstore/pkg/utils"
	"github.com/majid-2002/go-bookstore/pkg/models"
	"net/http"
	"strconv"
)


var NewBook models.Book //? create a new book variable of type Book from the models package


func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks() //? get all books from the models package
	
	res, _ := json.Marshal(newBooks) //? convert the books to json that we get from the sql database

	w.Header().Set("Content-Type", "pkglication/json") //? set the header to json

	w.WriteHeader(http.StatusOK) //? set the status to ok

	w.Write(res) //? write the response

}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //? get the params

	bookId := vars["bookId"] //? get the bookId

	ID, err := strconv.ParseInt(bookId, 0, 0) //? parse the bookId to int

	if err != nil {
		fmt.Println("Error while parsing")
	}

	bookDetails, _ := models.GetBookById(ID) //? get the book by id from the models package

	res, _ := json.Marshal(bookDetails) //? convert the book to json that we get from the sql database

	w.Header().Set("Content-Type", "application/json") //? set the header to json

	w.WriteHeader(http.StatusOK) //? set the status to ok

	w.Write(res) //? write the response

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{} //? create a new book of type Book from the models package

	utils.ParseBody(r, CreateBook) //? parse the body

	b := CreateBook.CreateBook() //? create the book

	res, _ := json.Marshal(b) //? convert the book to json

	w.WriteHeader(http.StatusOK) //? set the status to ok

	w.Write(res) //? write the response

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //? get the params

	bookId := vars["bookId"] //? get the bookId

	ID, err := strconv.ParseInt(bookId, 0, 0) //? parse the bookId to int

	if err != nil {
		fmt.Println("Error while parsing")
	}

	book := models.DeleteBook(ID) //? delete the book

	res, _ := json.Marshal(book) //? convert the book to json

	w.Header().Set("Content-Type", "application/json") //? set the header to json

	w.WriteHeader(http.StatusOK) //? set the status to ok

	w.Write(res) //? write the response

}


func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{} //? create a new book of type Book from the models package

	utils.ParseBody(r, updateBook) //? parse the body

	vars := mux.Vars(r) //? get the params

	bookId := vars["bookId"] //? get the bookId from the params

	ID, err := strconv.ParseInt(bookId, 0, 0) //? parse the bookId to int

	if err != nil {
		fmt.Println("Error while parsing")
	}
	
	bookDetails, db := models.GetBookById(ID) //? get the book by id from the models package

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}

	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}

	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails) //? save the book

	res, _ := json.Marshal(bookDetails) //? convert the book to json

	w.Header().Set("Content-Type", "application/json") //? set the header to json

	w.WriteHeader(http.StatusOK) //? set the status to ok

	w.Write(res) //? write the response
}
