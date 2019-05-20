package main

import (
	"database/sql"
	//"encoding/json"
	"log"
	"fmt"
	"net/http"

	"books-list/controller"
	"books-list/driver"
//	"books-list/model"


	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

//var books []model.Book
var db *sql.DB

func init() {
	gotenv.Load()
}


// Function to use Routing
func main() {

	db = driver.ConnectDB()
	controller := controllers.Controller{}

	router := mux.NewRouter()


	//Routers.........
	router.HandleFunc("/getBooks", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/getBook/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/addBook", controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/updateBook", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/deleteBook/{id}", controller.RemoveBook(db)).Methods("DELETE")

	fmt.Println("Server is running")
	log.Fatal(http.ListenAndServe(":8000", router))
}

	/*
		// use without DB(LOCAL)...........................................
		//adding Data........
		books = append(books, Book{ID: 1, Title: "Go Pointers", Author: "Author1", Year: "2000"},
			Book{ID: 2, Title: "Go Routines", Author: "Author2", Year: "2001"},
			Book{ID: 3, Title: "Go Routers", Author: "Author3", Year: "2002"},
			Book{ID: 4, Title: "Go Concurrency", Author: "Author4", Year: "2003"},
			Book{ID: 5, Title: "Go Good Parts", Author: "Author5", Year: "2004"})

	*/