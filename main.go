package main

import (
	"fmt"
	"os"
	"learngo-pockets/bookworms"
)

func main () {
	bookwormsList, err := bookworms.LoadBookworms("testdata/bookworms.json")

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load bookworms: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(bookwormsList)
}
