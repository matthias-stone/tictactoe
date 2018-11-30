package bots

import (
	"math/rand"
	"testing"
	"time"

	"github.com/matthias-stone/tictactoe"
	"github.com/stretchr/testify/assert"
)

func TestMiniMax_Move(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{"easy", "X-O\nO--\nX--\n", "---\n---\n--X\n"},
		{"block", "OO-\nXO-\nOXX\n", "--X\n---\n---\n"},
	}

	rand.Seed(time.Now().UnixNano())

	m := MiniMax{}
	for _, test := range tests {
		assert.Equal(t, test.output, m.Move(tictactoe.GameStateFromString(test.input), tictactoe.X).String(), test.name)
	}
}

func BenchmarkMiniMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MiniMax{}.Move(0, tictactoe.X)
	}
}
