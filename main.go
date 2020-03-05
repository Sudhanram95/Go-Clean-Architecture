package main

import (
	"fmt"
	"go-clean-architecture/controller"
	"go-clean-architecture/driver"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	db := driver.ConnectDB()

	controller := controller.Controller{}
	router := mux.NewRouter()

	router.HandleFunc("/books", controller.GetAllBooksFromDB(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBookById(db)).Methods("GET")
	router.HandleFunc("/book", controller.AddNewBookToDB(db)).Methods("POST")
	router.HandleFunc("/book", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.DeleteBookFromDB(db)).Methods("DELETE")

	fmt.Println("Server is running on port 8000")
	http.ListenAndServe(":8000", router)
}
