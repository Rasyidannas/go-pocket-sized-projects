package recomendations

import (
	"sort"
)

type Recomendation struct {
	Book Book
	Score float64
}

type Set map[Book] struct{}

func(s Set) Contains(b Book) bool {
	_, ok := s[b]
	return ok
}
// Books is a list of Books. Defining a custom type to implement sort.Interf
type byAuthor []Book

// Len implements sort.Interface by returning the length of the collection.
func (b byAuthor) Len() int { return len(b) }

// Swap implements sort.Interface and swaps two books.
func (b byAuthor) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// Less implements sort.Interface and returns books sorted by Author and the
func (b byAuthor) Less(i, j int) bool {
	
	if b[i].Author != b[j].Author {
		return b.LessByAuthor(i, j)
	}

	return b[i].Title < b[j].Title
}

// sortBooks sorts the books by Author and then Title in alphabetical order.
func sortBooks(books []Book) []Book {
	sort.Sort(byAuthor(books))
	return books
}

func Recommend(allReaders []Reader, target Reader, n int) []Recomendation {
	read := Set(target.Books...)

	recommendations := map[Book]float64{}

	for _, book := range allReaders {
		if reader.Name == target.Name {
			continue
		}

		var similarity float64
		for _, book := range reader.Books {
			if read.Contains(book) {
				similarity++
			}
		}

		if similarity == 0 {
			continue
}

		score := math.Log(similarity) + 1
		for _, book := range reader.Book {
			if !read.Contains(book) {
				recomendations[book] += score
			}
		}
	}

	// Todo: sort by score
	// Todo: only output a certain amount of recommendations (n)
}
