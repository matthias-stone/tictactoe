package tictactoe

import "testing"
import "github.com/stretchr/testify/assert"

func TestGameState_String(t *testing.T) {
	tests := []struct {
		name   string
		input  GameState
		output string
	}{
		{"empty", 0, "---\n---\n---\n"},
		{"centerO", 0x100, "---\n-O-\n---\n"},
		{"centerX", 0x200, "---\n-X-\n---\n"},
		{"topLeftO", 0x1, "O--\n---\n---\n"},
		{"topLeftX", 0x2, "X--\n---\n---\n"},
		{"intersperced", 0x21212, "X-O\n-X-\nO-X\n"},
		{"allO", 0x15555, "OOO\nOOO\nOOO\n"},
		{"allX", 0x2AAAA, "XXX\nXXX\nXXX\n"},
		{"complete", 0x1A966, "XOX\nOOX\nXXO\n"},
	}
	for _, test := range tests {
		assert.Equal(t, test.output, test.input.String(), test.name)
	}
}

func TestGameState_Winner(t *testing.T) {
	tests := []struct {
		name   string
		input  GameState
		output GameState
	}{
		{"empty", 0, Empty},
		{"centerO", 0x100, Empty},
		{"centerX", 0x200, Empty},
		{"topLeftO", 0x1, Empty},
		{"topLeftX", 0x2, Empty},
		{"intersperced", 0x21212, X},
		{"allO", 0x15555, O},
		{"allX", 0x2AAAA, X},
		{"complete", 0x1A966, Empty},
		{"Top", 0x2A, X},
		{"Middle", 0x510, O},
		{"Bottom", 0x2A000, X},
		{"Left", 0x1041, O},
		{"Center", 0x8208, X},
		{"Right", 0x10410, O},
		{"Angle\\", 0x20202, X},
		{"Angle/", 0x1110, O},
	}
	for _, test := range tests {
		assert.Equal(t, test.output, test.input.Winner(), test.name)
	}
}
