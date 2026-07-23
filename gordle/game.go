package gordle

import (
	"fmt"
)

// Game holds all information we need to play a game of gorlde
type Game struct{}

// New reutrns a Game, which can be used to Play!
func New() *Game {
	g := &Game{}

	return g
}

// Play runs the game.
func (g *Game) Play() {
	fmt.Println("Welcome to Gordle")

	fmt.Printf("Enter a guess: \n")
}
