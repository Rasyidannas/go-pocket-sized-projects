package main

import (
	"testing"

	"learngo-pockets/bookworms"
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

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := bookworms.LoadBookworms(tc.bookwormsFile)
			if tc.wantErr && err == nil {
				t.Fatal("expected an error, got none")
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if !equalBookworms(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
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
