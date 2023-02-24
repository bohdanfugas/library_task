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
	UnmarshalJson() error
	MarshalJson(newBooks entities.Books) ([]byte, error)
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
	err := b.UnmarshalJson()

	if err != nil {
		return nil, err
	}

	return b.books, nil
}

func (b *booksRepo) SetBooks(newBooks entities.Books) error {
	byteArr, err := b.MarshalJson(newBooks)

	// update our JSON with users
	err = os.WriteFile("database/books.json", byteArr, 0666)
	if err != nil {
		return err
	}

	return nil
}

func (b *booksRepo) UnmarshalJson() error {
	jsonFile, err := os.Open("database/books.json")
	defer jsonFile.Close()

	if err != nil {
		return err
	}
	byteValue, _ := io.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &b.books)
	if err != nil {
		return err
	}

	return nil
}

func (b *booksRepo) MarshalJson(newBooks entities.Books) ([]byte, error) {
	byteArr, err := json.MarshalIndent(&newBooks, "", "\t")

	if err != nil {
		return nil, err
	}

	return byteArr, nil
}
