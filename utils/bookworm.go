package bookworms

import (
	"os"
	"encoding/json"
	"sort"
	"fmt"
)


type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

// LoadBookworms reads the file and returns the list of bookworms
func LoadBookworms(filePath string) ([]Bookworm, error) {
	f, err := os.Open(filePath)
	
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var bookworms []Bookworm

	// Decode the file and store the content in the variable book
	err = json.NewDecoder(f).Decode(&bookworms)
	
	if err != nil {
		return nil, err
	}

	return bookworms, nil
}

// FindCommonBooks returns books that are on more than one bookworm's shelf.
func FindCommonBooks(bookworms []Bookworm) []Book {
	booksOnShelves := BooksCount(bookworms)

	var commonBooks []Book
	for book, count := range booksOnShelves {
		if count > 1 {
			commonBooks = append(commonBooks, book)
		}
	}

	return sortBooks(commonBooks)
}

// BooksCount registers all the books and their occurrences from the bookworm
func BooksCount(bookworms []Bookworm) map[Book]uint {
	count := make(map[Book]uint)

	for _, bookworm := range bookworms {
		for _, book := range bookworm.Books {
			count[book]++
		} 
	}

	return count
}

// sortBooks sorts the books by Author and then Title.
func sortBooks(books []Book) []Book {
	sort.Slice(books, func(i, j int) bool {
		if books[i].Author != books[j].Author {
			return books[i].Author < books[j].Author
		}
			return books[i].Title < books[j].Title
	})

	return books
}

// displayBooks prints out the titles and authors of a list of books
func DisplayBooks(books []Book) {
	for _, book := range books {
		fmt.Println("-", book.Title, "by", book.Author)
	}
}
