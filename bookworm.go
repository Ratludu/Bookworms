package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

func loadBookworms(filepath string) ([]Bookworm, error) {

	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var bookworms []Bookworm

	err = json.NewDecoder(f).Decode(&bookworms)

	if err != nil {
		return nil, err
	}

	return bookworms, nil

}

// registers all books for all bookworms
func booksCount(bookworms []Bookworm) map[Book]uint {

	count := make(map[Book]uint)

	for _, bookworm := range bookworms {
		for _, book := range bookworm.Books {
			count[book]++
		}
	}

	return count
}

// find common books returns books that are on more than one bookworms shelf.
func findCommonBooks(bookworms []Bookworm) []Book {
	booksOnShelves := booksCount(bookworms)

	var commonBooks []Book

	for book, count := range booksOnShelves {
		if count > 1 {
			commonBooks = append(commonBooks, book)
		}
	}

	return sortBooks(commonBooks)
}

// Implementing sort.Interface on Books
// byAuthor is a list of book
// Defining a custom type to implement sort.interface
type byAuthor []Book

// Len implements sort.Interface by returning the length of the BookbyAuthor
func (b byAuthor) Len() int { return len(b) }

// Swap implements sort.interfaces and dwaps two books
func (b byAuthor) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// Less implements the sort.Interface and returns books sorted by Author and then Title
func (b byAuthor) Less(i, j int) bool {
	if b[i].Author != b[j].Author {
		return b[i].Author < b[j].Author
	}

	return b[i].Title < b[j].Title
}

// sortBooks sorts the books by Author and then Title
func sortBooks(books []Book) []Book {

	sort.Sort(byAuthor(books))
	
	return books
}

// displayBooks print out the titles and authors of a list of books
func displayBooks(books []Book) {
	for _, book := range books {
		fmt.Println("-", book.Title, "by", book.Author)
	}
}
