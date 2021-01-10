package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Book struct {
	ID     string
	Title  string
	Author string
	Year   string
}

var books []Book

func main() {
	books = append(books, Book{"1", "Book1", "Author1", "2001"},
		Book{"2", "Book2", "Author2", "2003"},
		Book{"3", "Book3", "Author3", "2001"})

	router := mux.NewRouter()
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))

}

func getBooks(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode(books)
}
func getBook(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(writer).Encode(book)
		}
	}
	log.Println(params)
}
func addBook(writer http.ResponseWriter, request *http.Request) {
	var book Book
	json.NewDecoder(request.Body).Decode(book)

}
func updateBook(writer http.ResponseWriter, request *http.Request) {
	log.Println("Look up book in DB and update it")
}
func removeBook(writer http.ResponseWriter, request *http.Request) {
	log.Println("Remove book from DB")
}
