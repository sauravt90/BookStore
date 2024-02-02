package controllers

import (
	"encoding/json"
	"fmt"
	"go-bookstore/pkg/models"
	"go-bookstore/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var newBook models.Book

type ErrorResp struct {
	StatusCode  int16
	ErrorReason string
	Method      string
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBook := models.GetAllBooks()
	res, _ := json.Marshal(newBook)
	w.Header().Set("Content-Type", "pkglicatin/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	bookId := vars["bookId"]

	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(Id)

	res, _ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "pkglicatin/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}

	utils.ParseBody(r, CreateBook)

	b := CreateBook.CreateBook()

	res, _ := json.Marshal(b)

	w.Header().Set("Content-Type", "pkglicatin/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	bookId := vars["bookId"]

	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
		errResp := ErrorResp{ErrorReason: err.Error(), StatusCode: http.StatusBadRequest, Method: "DELETE"}
		res, _ := json.Marshal(errResp)
		w.Header().Set("Content-Type", "pkglicatin/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}
	bookDetails := models.DeleteBook(Id)

	res, _ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "pkglicatin/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)

	vars := mux.Vars(r)
	bookId := vars["bookId"]

	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, db := models.GetBookById(Id)

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)

	res, _ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "pkglicatin/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
