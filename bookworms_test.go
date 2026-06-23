package main

import (
	"testing"

	"learngo-pockets/utils"
)

var (
	handmaidsTale = bookworms.Book{Author: "Margaret Atwood", Title: "The Handmaid's Tale"}
	oryxAndCrake = bookworms.Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}
	theBellJar = bookworms.Book{Author: "Sylvia Plath", Title: "The Bell Jar"}
	janeEyre = bookworms.Book{Author: "Charlotte Brontë", Title: "Jane Eyre"}
)

func TestLoadBookworms_Success(t *testing.T) {
	tests := map[string]struct {
		bookwormsFile string
		want          []bookworms.Bookworm
		wantErr       bool
	}{
		"file exists": {
			bookwormsFile: "testdata/bookworms.json",
			want: []bookworms.Bookworm{
				{Name: "Fadi", Books: []bookworms.Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []bookworms.Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			wantErr: false,
		},
		"file doesn't exist": {
			bookwormsFile: "testdata/no_file_here.json",
			want:          nil,
			wantErr:       true,
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := bookworms.LoadBookworms(testCase.bookwormsFile)
			if err != nil && !testCase.wantErr {
				t.Fatalf("unexpected error: %s", err.Error())
			}
			if err == nil && testCase.wantErr {
				t.Fatalf("expected an error, got none %s", err.Error())
			}
			if !equalBookworms(got, testCase.want) {
				t.Fatalf("different result: got %v, expected %v", got, testCase.want)
			}
		})
	}
}

func equalBookworms(got, target []bookworms.Bookworm) bool {
	if len(got) != len(target) {
		return false
	}
	for i := range got {
		if got[i].Name != target[i].Name {
			return false
		}
		if !equalBooks(got[i].Books, target[i].Books) {
			return false
		}
	}
	return true
}

func equalBooks(books, target []bookworms.Book) bool {
	if len(books) != len(target) {
		return false
	}

	for i := range books {
		if books[i] != target[i] {
			return false
		}
	}

	return true
}

func TestFindCommonBooks(t *testing.T) {
	tt := map[string]struct {
		input []bookworms.Bookworm
		want  []bookworms.Book
	}{
		"no common book": {
			input: []bookworms.Bookworm{
				{Name: "Fadi", Books: []bookworms.Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []bookworms.Book{oryxAndCrake, janeEyre}},
			},
			want: nil,
		},
		"one common book": {
			input: []bookworms.Bookworm{
				{Name: "Fadi", Books: []bookworms.Book{handmaidsTale}},
				{Name: "Peggy", Books: []bookworms.Book{handmaidsTale}},
			},
			want: []bookworms.Book{handmaidsTale},
		},
		"three bookworms share a book": {
			input: []bookworms.Bookworm{
				{Name: "Fadi", Books: []bookworms.Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []bookworms.Book{handmaidsTale, janeEyre}},
				{Name: "Pip", Books: []bookworms.Book{handmaidsTale, oryxAndCrake}},
			},
			want: []bookworms.Book{handmaidsTale},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := bookworms.FindCommonBooks(tc.input)
			if !equalBooks(got, tc.want) {
				t.Fatalf("got a different list of books: %v, expected %v", got, tc.want)
			}
		})
	}
}

func TestBookCount(t *testing.T) {
	tt := map[string]struct {
		input []bookworms.Bookworm
		want  map[bookworms.Book]uint
	}{
		"nominal use case": {
			input: []bookworms.Bookworm{
				{Name: "Fadi", Books: []bookworms.Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []bookworms.Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			want: map[bookworms.Book]uint{handmaidsTale: 2, theBellJar: 1, oryxAndCrake: 1, janeEyre: 1},
		},
		"no bookworms": {
			input: []bookworms.Bookworm{},
			want:  map[bookworms.Book]uint{},
		},
		"bookworm without books": {
			input: []bookworms.Bookworm{
				{Name: "Fadi", Books: []bookworms.Book{}},
			},
			want: map[bookworms.Book]uint{},
		},
		"bookworm with twice the same book": {
			input: []bookworms.Bookworm{
				{Name: "Fadi", Books: []bookworms.Book{handmaidsTale, handmaidsTale}},
			},
			want: map[bookworms.Book]uint{handmaidsTale: 2},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := bookworms.BooksCount(tc.input)
			if !equalCounts(got, tc.want) {
				t.Fatalf("got a different list of books: %v, expected %v", got, tc.want)
			}
		})
	}
}

func equalCounts(got, want map[bookworms.Book]uint) bool {
	if len(got) != len(want) {
		return false
	}

	for book, targetCount := range want {
		count, ok := got[book]
		if !ok || targetCount != count {
			return false
		}
	}
	return true
}
