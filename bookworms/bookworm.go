package bookworms

import (
	"os"
	"encoding/json"
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
