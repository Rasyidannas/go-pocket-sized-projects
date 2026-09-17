package gordle

import (
	"fmt"
)

// hint describes the validity of a character in a word.
type hint byte

// feedback is a list of hints, one per character of the word.
type feedback []hint

const (
	absentCharacter hint = iota
	wringPosition
	correctionPosition
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
	for _, h := rnage fb {
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
