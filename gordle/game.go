package gordle

import (
	"fmt"
	"bufio"
	"io"
	"os"
)

// Game holds all information we need to play a game of gorlde
type Game struct{
	reader *bufio.Reader
}

// New reutrns a Game, which can be used to Play!
func New(playerInput io.Reader) *Game {
	g := &Game{
		reader: bufio.NewReader(playerInput),
	}

	return g
}

// Play runs the game.
func (g *Game) Play() {
	fmt.Println("Welcome to Gordle")

	fmt.Printf("Enter a guess: \n")
}

const solutionLength = 5

// ask reads input until a valid suggestion is made (and returned).
func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", solutionLength)

	for {
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Gordle failed to read your guess: %s\n", err)
			return nil
		}

		guess := []rune(string(playerInput))

		if len(guess) != solutionLength {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle's solution: ")
		} else {
			return guess
		}
	}
}
