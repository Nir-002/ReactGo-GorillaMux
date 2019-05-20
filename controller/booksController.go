package controllers

import (
	"database/sql"
 	"encoding/json"
	"log"
	"net/http"
	"books-list/utils"
	"strconv"

	
	"books-list/repository/book"
	"books-list/model"
	"github.com/gorilla/mux"
)

type Controller struct {}

var books []models.Book 
var db *sql.DB

// For error capturing......
func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// For enabling CORS..........
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}


func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	
	return func (w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error

		books = []models.Book{}
		bookRepo := bookRepository.BookRepository{}
		books, err := bookRepo.GetBooks(db, book, books)

		enableCors(&w)

		if err != nil {
			error.Message = "Server Error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, books)

		json.NewEncoder(w).Encode(books)
	
		// use without DB.....................................
		//	json.NewEncoder(w).Encode(books)
	}
}

func(c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func( w http.ResponseWriter, r *http.Request) {
			var book models.Book
			var error models.Error
	
			params := mux.Vars(r)
	
			books = []models.Book{}
			bookRepo := bookRepository.BookRepository{}
	
			id, _ :=strconv.Atoi(params["id"])
		
			books, err := bookRepo.GetBook(db, book, id)
			
			enableCors(&w)

			if err != nil {
			
				if err == sql.ErrNoRows {
					error.Message = "Not Found"
					utils.SendError(w, http.StatusNotFound, error)
					return
				} else {
				
				error.Message = "Server Error"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			
				}			
			}
	
			w.Header().Set("Content-Type", "application/json")
			utils.SendSuccess(w, books)
	
			json.NewEncoder(w).Encode(books)
		
	}
}

func(c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var bookID int
		var error models.Error

		enableCors(&w)

		json.NewDecoder(r.Body).Decode(&book)

		if book.Author == "" || book.Title == "" || book.Year == "" {
			error.Message = "Enter missing fields"
			utils.SendError(w, http.StatusBadRequest, error)
			return 
		}

		bookRepo := bookRepository.BookRepository{}
		bookID, err := bookRepo.AddBook(db, book)

		if err != nil {
			error.Message = "Server Error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, bookID)
	}
}
func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error
		
		enableCors(&w)

		json.NewDecoder(r.Body).Decode(&book)

		if book.ID == 0 || book.Author == "" || book.Title == "" || book.Year == "" {
			error.Message = "All fields are required"
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		bookRepo := bookRepository.BookRepository{}
		rowsUpdated, err := bookRepo.UpdateBook(db, book)

		if err != nil {
			error.Message = "Server Error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, rowsUpdated)
	}
}

func (c Controller) RemoveBook (db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var error models.Error
		
		params := mux.Vars(r)

		bookRepo := bookRepository.BookRepository{}
		id, _ := strconv.Atoi(params["id"])

		enableCors(&w)
		
		rowsDeleted, err := bookRepo.RemoveBook(db, id)

		if err != nil {
			error.Message = "Server Error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		if rowsDeleted == 0 {
			error.Message = "Not Found"
			utils.SendError(w, http.StatusNotFound, error)
			return 
		}

	}
}

