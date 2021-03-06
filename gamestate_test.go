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
		{"AllO", 0x15555, "OOO\nOOO\nOOO\n"},
		{"AllX", 0x2AAAA, "XXX\nXXX\nXXX\n"},
		{"complete", 0x1A966, "XOX\nOOX\nXXO\n"},
		{"TopMask", (Pos1 | Pos2 | Pos3) & AllO, "OOO\n---\n---\n"},
		{"MiddleMask", (Pos4 | Pos5 | Pos6) & AllX, "---\nXXX\n---\n"},
		{"BottomMask", (Pos7 | Pos8 | Pos9) & AllO, "---\n---\nOOO\n"},
		{"LeftMask", (Pos1 | Pos4 | Pos7) & AllX, "X--\nX--\nX--\n"},
		{"CenterMask", (Pos2 | Pos5 | Pos8) & AllO, "-O-\n-O-\n-O-\n"},
		{"RightMask", (Pos3 | Pos6 | Pos9) & AllX, "--X\n--X\n--X\n"},
		{"AngleDownMask", (Pos1 | Pos5 | Pos9) & AllO, "O--\n-O-\n--O\n"},
		{"AngleUpMask", (Pos3 | Pos5 | Pos7) & AllX, "--X\n-X-\nX--\n"},
	}
	for _, test := range tests {
		assert.Equal(t, test.output, test.input.String(), test.name)
		assert.Equal(t, test.input, GameStateFromString(test.output), test.name)
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
		{"AllO", 0x15555, O},
		{"AllX", 0x2AAAA, X},
		{"complete", 0x1A966, Empty},
		{"Top", 0x2A, X},
		{"Middle", 0x540, O},
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

func BenchmarkWinner(b *testing.B) {
	cases := []GameState{
		0,
		0x100,
		0x200,
		0x1,
		0x2,
		0x21212,
		0x15555,
		0x2AAAA,
		0x1A966,
		(Pos1 | Pos2 | Pos3) & AllO,
		(Pos4 | Pos5 | Pos6) & AllX,
		(Pos7 | Pos8 | Pos9) & AllO,
		(Pos1 | Pos4 | Pos7) & AllX,
		(Pos2 | Pos5 | Pos8) & AllO,
		(Pos3 | Pos6 | Pos9) & AllX,
		(Pos1 | Pos5 | Pos9) & AllO,
		(Pos3 | Pos5 | Pos7) & AllX,
	}

	for i := 0; i < b.N; i++ {
		for _, gs := range cases {
			gs.Winner()
		}
	}
}
