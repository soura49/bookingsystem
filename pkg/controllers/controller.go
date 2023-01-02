package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sour49/bookingsystem/pkg/models"
	"github.com/sour49/bookingsystem/pkg/utils"
)

var NewBook models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := &models.Book{}
	createDBRecord := newBook.CreateBook()
	fmt.Println(newBook)
	res, _ := json.Marshal(createDBRecord)
	fmt.Println(res)
	w.WriteHeader(http.StatusOK)
	log.Println(w.Write(res))
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		panic(err)
	}
	bookDetails, _ := models.GetBookById(ID)

	res, err := json.Marshal(bookDetails)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var bookNeededToBeUpdated = &models.Book{}
	utils.ParseBody(r, bookNeededToBeUpdated)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while Parsing")
	}
	bookDetails, db := models.GetBookById(ID)

	if bookNeededToBeUpdated.Name != "" {
		bookDetails.Name = bookNeededToBeUpdated.Name
	}

	if bookNeededToBeUpdated.Author != "" {
		bookDetails.Author = bookNeededToBeUpdated.Author
	}

	if bookNeededToBeUpdated.Publication != "" {
		bookDetails.Publication = bookNeededToBeUpdated.Publication
	}

	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	w.Write(res)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookIdToBeDeleted := vars["id"]
	ID, err := strconv.ParseInt(bookIdToBeDeleted, 0, 0)
	if err != nil {
		fmt.Println("Error Parsing the Data")
	}
	book := models.DeleteBook(ID)

	res, err := json.Marshal(book)

	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
