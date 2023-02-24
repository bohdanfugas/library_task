package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"main/repository"
	"net/http"
)

func HandleBookRequests(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getBooks(w, r)
	} else if r.Method == http.MethodPost {
		postBooks(w, r)
	} else {
		fmt.Println("supporting only POST and GET requests")
	}
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	booksRepo := repository.NewBookRepo()
	books, err := booksRepo.GetBooks()

	if err != nil {
		println(err)
		return
	}

	// // sort by date
	// sort.Slice(books, func(i, j int) bool {
	// 	return books[i].Year < books[j].Year
	// })

	// write sorted books slice
	json.NewEncoder(w).Encode(books)
}

func postBooks(w http.ResponseWriter, r *http.Request) {
	booksRepo := repository.NewBookRepo()
	books, err := booksRepo.GetBooks()
	if err != nil {
		println(err)
		return
	}

	// unmarshall req body
	newBookByteValue, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	newBook := repository.CreateNewBook()
	json.Unmarshal(newBookByteValue, &newBook)

	// add new book to books slice
	books = append(books, newBook)

	err = booksRepo.SetBooks(books)
	if err != nil {
		println(err)
		return
	}
}
