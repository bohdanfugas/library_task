package repository

import (
	"encoding/json"
	"io"
	"main/entities"
	"os"
)

type BooksRepo interface {
	GetBooks() (entities.Books, error)
	SetBooks(newBooks entities.Books) error
	UnmarshalJSON(byteValue []byte) error
	MarshalJSON() ([]byte, error)
}

type booksRepo struct {
	books entities.Books
}

func NewBookRepo() BooksRepo {
	return &booksRepo{}
}

func CreateNewBook() entities.Book {
	return entities.Book{}
}

func (b *booksRepo) GetBooks() (entities.Books, error) {
	jsonFile, err := os.Open("database/books.json")
	defer jsonFile.Close()
	if err != nil {
		return nil, err
	}

	byteValue, _ := io.ReadAll(jsonFile)

	err = b.UnmarshalJSON(byteValue)
	if err != nil {
		return nil, err
	}

	return b.books, nil
}

func (b *booksRepo) SetBooks(newBooks entities.Books) error {
	b.books = newBooks
	byteArr, err := b.MarshalJSON()

	// update our JSON with users
	err = os.WriteFile("database/books.json", byteArr, 0666)
	if err != nil {
		return err
	}

	return nil
}

func (b *booksRepo) UnmarshalJSON(byteValue []byte) error {
	err := json.Unmarshal(byteValue, &b.books)
	if err != nil {
		return err
	}

	return nil
}

func (b *booksRepo) MarshalJSON() ([]byte, error) {
	byteArr, err := json.MarshalIndent(&b.books, "", "\t")
	if err != nil {
		return nil, err
	}

	return byteArr, nil
}
