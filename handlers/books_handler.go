package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	entities "main/entities"
	"net/http"
	"os"
	"sort"
)

func HandleBookRequests(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		handleBooksGet(w, r)
	} else if r.Method == "POST" {
		handleBooksPost(w, r)
	} else {
		fmt.Println("supporting only POST and GET requests")
	}
}

func handleBooksGet(w http.ResponseWriter, r *http.Request) {
	// unmarshal our JSON with books
	var books []entities.Book
	jsonFile := unmarshalBooksJson(&books)
	defer jsonFile.Close()

	// sort by date
	sort.Slice(books, func(i, j int) bool {
		return books[i].Year < books[j].Year
	})

	// write sorted books slice
	json.NewEncoder(w).Encode(books)
}

func handleBooksPost(w http.ResponseWriter, r *http.Request) {
	// unmarshall our books JSON
	var books []entities.Book
	jsonFile := unmarshalBooksJson(&books)
	defer jsonFile.Close()

	// unmarshall req body
	newBookByteValue, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	var newBook entities.Book
	json.Unmarshal(newBookByteValue, &newBook)

	// add new book to books slice
	books = append(books, newBook)

	// marshal books to JSON
	byteArr, err := json.MarshalIndent(books, "", "\t")
	if err != nil {
		fmt.Println(err)
	}

	// update our JSON with users
	err = os.WriteFile("database/books.json", byteArr, 0666)
	if err != nil {
		fmt.Println(err)
	}
}

func unmarshalBooksJson(outSlice *[]entities.Book) *os.File {
	jsonFile, err := os.Open("database/books.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, outSlice)

	return jsonFile
}
