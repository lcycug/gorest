package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// Book Struct
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books as a slice of Book Strcut
var books []Book

// Get all books
func getBooks(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(books)
}

// Get a book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Get params
	params := mux.Vars(r)
	// Loop through all books to find a matched ID
	for _, b := range books {
		if b.ID == params["id"] {
			_ = json.NewEncoder(w).Encode(b)
			return
		}
	}
	_ = json.NewEncoder(w).Encode(&Book{})
}

// Create a book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	_ = json.NewEncoder(w).Encode(book)
}

// Update a book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Get params
	params := mux.Vars(r)
	// Loop through all books to find a matched ID
	for i, b := range books {
		if b.ID == params["id"] {
			books = append(books[:i], books[i+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			_ = json.NewEncoder(w).Encode(book)
			return
		}
	}
	_ = json.NewEncoder(w).Encode(&Book{})
}

// Delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Get params
	params := mux.Vars(r)
	// Loop through all books to find a matched ID
	for i, b := range books {
		if b.ID == params["id"] {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}
	_ = json.NewEncoder(w).Encode(books)
}

func main() {

	// Mock data
	books = append(books, Book{
		ID:    "1",
		Isbn:  "123",
		Title: "Go 101",
		Author: &Author{
			Firstname: "Tapir",
			Lastname:  "Liu",
		},
	})
	books = append(books, Book{
		ID:    "2",
		Isbn:  "567",
		Title: "The Little Go Book",
		Author: &Author{
			Firstname: "Karl",
			Lastname:  "Seguin",
		},
	})

	// Init the router
	r := mux.NewRouter()

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
