package bookworms

import (
	"sort"
	"math"
)

type Recommendation struct {
	Book Book
	Score float64
}

type Set map[Book] struct{}

func NewSet(books ...Book) Set {
	s := make(Set)
	for _, b := range books {
			s[b] = struct{}{}
	}
	return s
}

func(s Set) Contains(b Book) bool {
	_, ok := s[b]
	return ok
}

func Recommend(allReaders []Bookworm, target Bookworm, n int) []Recommendation {
	read := NewSet(target.Books...)

	scores := map[Book]float64{}

	for _, reader := range allReaders {
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
			for _, book := range reader.Books {
					if !read.Contains(book) {
							scores[book] += score
					}
			}
	}

	// Convert map to slice and sort by score descending
	recommendations := make([]Recommendation, 0, len(scores))
	for book, score := range scores {
			recommendations = append(recommendations, Recommendation{Book: book, Score: score})
	}

	sort.Slice(recommendations, func(i, j int) bool {
			return recommendations[i].Score > recommendations[j].Score
	})

	if n > len(recommendations) {
			n = len(recommendations)
	}

	return recommendations[:n]
}
