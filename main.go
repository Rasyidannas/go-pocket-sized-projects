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

	commonBooks := bookworms.FindCommonBooks(bookwormsList)

	fmt.Println("Here are the books in common:")
	bookworms.DisplayBooks(commonBooks)
}
