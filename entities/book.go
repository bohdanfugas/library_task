package entities

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Year   int    `json:"year,string"`
}
