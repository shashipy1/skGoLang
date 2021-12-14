package main

import "fmt"

type Books struct {
	title string
	id    int
}

func main() {
	var book1 Books
	var book2 Books

	book1.title = "Golang"
	book1.id = 1

	book2.title = "SHASHI is learnin"
	book2.id =

		printBook(book1)
	printBook(book2)

}
func printBook(BookBooks) {
	fmt.Printf("book title %s\n", Book.title)
	fmt.Printf("book id %d\n", Book.id)

}
