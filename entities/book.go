package entities

type Books []Book

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Year   int    `json:"year,string"`
}
