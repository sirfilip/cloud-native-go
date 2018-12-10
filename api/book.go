package api

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Description string `json:"description,omitempty"`
}

func (b Book) ToJSON() []byte {
	data, err := json.Marshal(b)

	if err != nil {
		panic(err)
	}

	return data
}

func BookFromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}

var Books = map[string]Book{
	"123": Book{Title: "Going out of the dimensions", Author: "Filip K", ISBN: "123"},
	"456": Book{Title: "Programming Go", Author: "Walter S", ISBN: "456"},
}

func BooksHandlerFunc(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodGet:
		books := AllBooks()
		writeJSON(w, books)
	case http.MethodPost:
		book := Book{}
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		isbn, created := CreateBook(book)
		if created {
			w.Header().Add("Location", "/api/books/"+isbn)
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusConflict)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}
}

func BookHandlerFunc(w http.ResponseWriter, r *http.Request) {
	isbn := r.URL.Path[len("/api/books/"):]
	switch method := r.Method; method {
	case http.MethodGet:
		if book, found := GetBook(isbn); found {
			writeJSON(w, book)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodPut:
		book := Book{}
		if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		if updated := UpdateBook(isbn, book); !updated {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Add("Location", "/api/books/"+isbn)
	case http.MethodDelete:
		if deleted := DeleteBook(isbn); !deleted {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}
}

func AllBooks() []Book {
	books := make([]Book, 0, len(Books))
	for _, book := range Books {
		books = append(books, book)
	}
	return books
}

func GetBook(isbn string) (Book, bool) {
	if _, found := Books[isbn]; !found {
		return Book{}, false
	} else {
		return Books[isbn], true
	}
}

func CreateBook(book Book) (string, bool) {
	if _, found := Books[book.ISBN]; found {
		return "", false
	}
	Books[book.ISBN] = book
	return book.ISBN, true
}

func UpdateBook(isbn string, book Book) bool {
	_, found := Books[isbn]
	if !found {
		return false
	}
	Books[isbn] = book
	return true
}

func DeleteBook(isbn string) bool {
	if _, found := Books[isbn]; !found {
		return false
	}
	delete(Books, isbn)
	return true
}

func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(data)
}
