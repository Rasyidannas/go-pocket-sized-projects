package gordle

import (
	"errors"
	"strings"
	"testing"

	"golang.org/x/exp/slices"
)

func TestGameValidateGuess(t *testing.T) {
	tt := map[string]struct {
		word []rune
		expected error
	}{
		"nominal": {
			word: []rune("GUESS"),
			expected: nil,
		},
		"too long": {
			word: []rune("POCKET"),
			expected: errInvalidWordLength,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g := New(nil, string(tc.word), 0)

			err := g.validateGuess(tc.word)
			if !errors.Is(err, tc.expected) {
				t.Errorf("%c, expected %q, got %q", tc.word, tc.expected, err)
			}	
		})
	}
}

func TestGameAsk(t *testing.T) {
	tt := map[string]struct {
		input string
		want []rune
	}{
		"5 characters in english": {
			input: "HELLO\n",
			want: []rune("HELLO"),
		},
		"5 characters in arabic": {
			input: "مرحبا\n",
			want: []rune("مرحبا"),
		},
		"5 characters in japanese": {
			input: "こんにちは\n",
			want: []rune("こんにちは"),
		},
		"3 characters in japanese": {
			input: "こんに\nこんにちは\n",
			want: []rune("こんにちは"),
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g := New(strings.NewReader(tc.input), string(tc.want), 0)

			got := g.ask()
			if !slices.Equal(got, tc.want) {
				t.Errorf("got = %v, want %v", string(got), string(tc.want))
			}
		})
	}
}

func TestComputeFeedback(t *testing.T) {
	tt := map[string]struct {
		guess            string
		solution         string
		expectedFeedback feedback
	}{
		"nominal": {
			guess:    "hello",
			solution: "hello",
			expectedFeedback: feedback{correctPosition, correctPosition, correctPosition, correctPosition, correctPosition},
		},
		"double character": {
			guess:    "heelo",
			solution: "hello",
			expectedFeedback: feedback{correctPosition, correctPosition, absentCharacter, correctPosition, correctPosition},
		},
		"double characters with wrong answer": {
			guess:    "hhelo",
			solution: "hello",
			expectedFeedback: feedback{correctPosition, absentCharacter, wrongPosition, correctPosition, correctPosition},
		},
		"two identical, but not in the right position (from left to right)": {
			guess:    "hlleo",
			solution: "hello",
			expectedFeedback: feedback{correctPosition, wrongPosition, correctPosition, wrongPosition, correctPosition},
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			fb := computeFeedback([]rune(tc.guess), []rune(tc.solution))
			if !tc.expectedFeedback.Equal(fb) {
				t.Errorf("guess: %q, got the wrong feedback, expected %v, got %v", tc.guess, tc.expectedFeedback, fb)
			}
		})
	}
}
