package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	entities "main/entities"
	handlers "main/handlers"
	utils "main/utils"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/books", handlers.HandleBookRequests)

	go func() {
		runTest()
	}()

	http.ListenAndServe(":8080", mux)
}

func runTest() {
	fmt.Println("=================================")
	fmt.Println("OUR JSON ON START")
	fmt.Println("=================================")
	getBooksData()

	fmt.Println("=================================")
	fmt.Println("ADDING NEW BOOK")
	fmt.Println("=================================")
	newBook := entities.Book{
		Name:   "Refactoring",
		Author: "Martin Fowler",
		Year:   1999,
	}
	postBooksData(newBook)

	fmt.Println(" ")
	fmt.Println("=================================")
	fmt.Println("OUR JSON AFTER ADDING NEW BOOK")
	fmt.Println("=================================")
	getBooksData()
}

func getBooksData() {
	requestURL := "http://localhost:8080/books"
	client := &http.Client{}

	r, _ := http.NewRequest(http.MethodGet, requestURL, nil)
	resp, err := client.Do(r)
	if err != nil {
		fmt.Println("HTTP call failed:", err)
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	prettyJSONString, err := utils.PrettyString(string(body))
	fmt.Println(prettyJSONString)
}

func postBooksData(newBook entities.Book) {
	byteArr, err := json.MarshalIndent(newBook, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	// timeout 500 ms
	time.Sleep(500 * time.Millisecond)

	client := http.Client{}
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/books", bytes.NewBuffer(byteArr))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Unable to reach the server.")
	} else {
		fmt.Println(resp.Status)
		fmt.Printf("Name: %s, Author: %s, Year: %d.\n", newBook.Name, newBook.Author, newBook.Year)
	}
}
