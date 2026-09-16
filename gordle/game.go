package gordle

import (
	"fmt"
	"bufio"
	"io"
	"os"
	"strings"
	"slices"
)

// Game holds all information we need to play a game of gorlde
type Game struct{
	reader *bufio.Reader
	solution []rune
	maxAttempts int
}

// New reutrns a Game, which can be used to Play!
func New(playerInput io.Reader, solution string, maxAttempts int) *Game {
	g := &Game{
		reader: bufio.NewReader(playerInput),
		solution: splitToUppercaseCharacters(solution),
		maxAttempts: maxAttempts,
	}

	return g
}

// Play runs the game.
func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")

	for currentAttempt := 1; currentAttempt <= g.maxAttempts; currentAttempt++ {
		//ask for a valid word
		guess := g.ask()

		if slices.Equal(guess, g.solution) {
			fmt.Printf("You won! You found it in %d guess(es)!", len(guess))
			return
		}
	}

	fmt.Printf("You've lost! The solution was: %s. \n", string(g.solution))
}

const solutionLength = 5

// errInvalidWordLength is returned when the guess has the wrong number of chracters
var errInvalidWordLength = fmt.Errorf("invalid guess, word doesn't have the correct number of letters")

// validateGuess ensures the guess is valid enough.
func (g *Game) validateGuess(guess []rune) error {
	if len(guess) != solutionLength {
		return fmt.Errorf("expected %d, got %d, %w", solutionLength, len(guess), errInvalidWordLength)
	}
	return nil
}
  
// ask reads input until a valid suggestion is made (and returned).
func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", len(g.solution))

	for {
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Gordle failed to read your guess: %s\n", err)
			return nil
		}

		guess := splitToUppercaseCharacters(string(playerInput))

		err = g.validateGuess(guess)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle's solution: ")
		} else {
			return guess
		}
	}
}

// plitToUppercaseCharacters is a naive implementation to turn a string int
func splitToUppercaseCharacters(input string) []rune {
	return []rune(strings.ToUpper(input))
}
