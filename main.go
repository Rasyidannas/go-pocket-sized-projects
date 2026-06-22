package main

import (
	"flag"
	"fmt"
)

func main() {
	var lang string
	flag.StringVar(&lang, "lang", "en", "The required language, e.g. en, ur, el, fr, he and vi.")
	flag.Parse()

	// fmt.Printf("lang: %v", lang)
	
	greeting := greet(language(lang))
	fmt.Println(greeting)
}

type language string

var phrasebook = map[language] string {
	"el": "Χαίρετε Κόσμε",
	"en": "Hello world",
	"fr": "Bonjour le monde",
	"he": "שלום עולם",
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

