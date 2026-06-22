package main

import "fmt"

func main() {
	greeting := greet("id")
	fmt.Println(greeting)
}

type language string

var phrasebook = map[language] string {
	"el": "Χαίρετε Κόσμε",
	"en": "Hello world",
	"fr": "Bonjour le monde",
	"שלום עולם" :"he",
	"ur": " ہﯿﻠﻮ ",
	"vi": "Xin chào Thế Giới",
}

// greet returns a greeting to the world
func greet(l language) string {
	greeting, ok := phrasebook[l]

	// fmt.Printf("Ok is %v \n", ok)

	if !ok {
		return fmt.Sprintf("Unsupported language: %q", l)
	}

	return greeting
}

