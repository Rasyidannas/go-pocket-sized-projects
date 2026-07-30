package gordle

import (
	"strings"
	"testing"

	"golang.org/x/exp/slices"
)

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
			g := New(strings.NewReader(tc.input))

			got := g.ask()
			if !slices.Equal(got, tc.want) {
				t.Errorf("got = %v, want %v", string(got), string(tc.want))
			}
		})
	}
}

