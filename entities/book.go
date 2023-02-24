package entities

type Books []Book

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Year   int    `json:"year,string"`
}

func (b Books) Len() int {
	return len(b)
}

func (b Books) Less(i, j int) bool {
	return int32(b[i].Year) < int32(b[j].Year)
}

func (b Books) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
