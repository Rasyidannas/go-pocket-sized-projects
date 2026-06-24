package main

import (
	"fmt"
	"os"
	"learngo-pockets/utils"
)

func main () {
	bookwormsList, err := bookworms.LoadBookworms("testdata/bookworms.json")

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load bookworms: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(bookwormsList)

	if len(bookwormsList) > 0 {
    target := bookwormsList[0]
    fmt.Printf("\nRecommendations for %s:\n", target.Name)
    recommendations := bookworms.Recommend(bookwormsList, target, 2)
    for _, r := range recommendations {
        fmt.Printf("- %s by %s (score: %.2f)\n", r.Book.Title, r.Book.Author, r.Score)
    }
	}

	commonBooks := bookworms.FindCommonBooks(bookwormsList)

	fmt.Println("Here are the books in common:")
	bookworms.DisplayBooks(commonBooks)
}
