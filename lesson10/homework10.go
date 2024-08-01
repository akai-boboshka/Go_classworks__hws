package main

import "fmt"

type Book struct {
	title  string
	author string
}

func (b Book) GetDetails() string {
	titleOfBook := b.title
	// authorOfBook := b.author
	return titleOfBook
	// return authorOfBook
	// return b.title
	// return b.author
}

type Library struct {
	CatalogOfBooks []string
}

func (l Library) AddBook(addBook []string) {
	addBook = append(addBook, "Таджики", "Пинокио")
	fmt.Println(addBook)
}

func main() {
	result := Book{"Harry Potter", "S.Ayiny"}
	fmt.Println(result.GetDetails())

	addNewBook := Library{}
	fmt.Println(addNewBook)
}
