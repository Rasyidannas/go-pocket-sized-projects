package gordle

import (
	"fmt"
	"os"
	"strings"
)

// hint describes the validity of a character in a word.
type hint byte

// feedback is a list of hints, one per character of the word.
type feedback []hint

const (
	absentCharacter hint = iota
	wrongPosition
	correctPosition
)

// String implements the Stringer interface
func (h hint) String() string {
	switch h {
		case absentCharacter:
			return "⬜" // grey square
		case wrongPosition:
			return "🟡" // yellow circle
		case correctPosition:
			return "💚" //green heart
		default:
			// This should never happen
			return "💔"
	}
}

// StringConcat is a naive implementation to build feedback as a string.
// It is used only to benchmark it against the strings.Builder version.
func (fb feedback) StringConcat() string {
	var output string
	for _, h := range fb {
		output += h.String()
	}

	return output
}

// String implements the Stringer interface for a slice of hints.
func (fb feedback) String() string {
	sb := strings.Builder{}
	for _, h := range fb {
		sb.WriteString(h.String())
	}

	return sb.String()
}

// computeFeedback verifies every character of the guess against
func computeFeedback(guess, solution []rune) feedback {
	// initialise holder for marks
	result := make(feedback, len(guess))
	used := make([]bool, len(solution))

	if len(guess) != len(solution) {
		_, _ = fmt.Fprintf(os.Stderr, "Internal error! Guess and solution have equal of count characters")
		return result
	}

	// check for correct letters
	for posInGuess, character := range guess {
		if character == solution[posInGuess] {
			result[posInGuess] = correctPosition
			used[posInGuess] = true
		}
	}
	
	//look for letters in the wrong position
	for posInGuess, character := range guess {
		if result[posInGuess] != absentCharacter {
			// The character has already been marked, ignore it
			continue
		}

		for posInSolution, target := range solution {
			if used[posInSolution] {
				// The letter of the solution is already assigned to a letter of
				// Skip to the next letter of the solution.}
				continue
			}
			if character == target {
				result[posInGuess] = wrongPosition
				used[posInSolution] = true
				//Skip to the next letter of the guess
				break
			}
		}
	}

	return result
}

// Equal determines equality of two feedbacks
func (fb feedback) Equal(other feedback) bool {
	if len(fb) != len(other) {
		return false
	}

	for index, value := range fb {
		if value != other[index] {
			return false
		}
	}
	return true
}
