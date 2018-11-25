package tictactoe

import "testing"
import "github.com/stretchr/testify/assert"

func TestGameStateString(t *testing.T) {
	tests := []struct {
		name   string
		input  GameState
		output string
	}{
		{"empty", 0, "---\n---\n---\n"},
		{"centerO", 0x100, "---\n-O-\n---\n"},
		{"centerX", 0x200, "---\n-X-\n---\n"},
		{"intersperced", 0x21212, "X-O\n-X-\nO-X\n"},
		{"allO", 0x15555, "OOO\nOOO\nOOO\n"},
		{"allX", 0x2AAAA, "XXX\nXXX\nXXX\n"},
		{"complete", 0x1A966, "XOX\nOOX\nXXO\n"},
	}
	for _, test := range tests {
		assert.Equal(t, test.output, test.input.String(), test.name)
	}
}
