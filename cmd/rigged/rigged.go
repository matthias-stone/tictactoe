package main

import (
	"math/rand"

	"github.com/matthias-stone/tictactoe"
	"github.com/matthias-stone/tictactoe/bots"
)

// Corners tries to take the corners.
type Corners struct{}

func (r Corners) Name() string { return "random" }

const (
	CornerMask = tictactoe.Pos1 | tictactoe.Pos3 | tictactoe.Pos7 | tictactoe.Pos9
	Xcorners   = tictactoe.AllX & CornerMask
	Ocorners   = tictactoe.AllO & CornerMask
)

func (r Corners) Move(gs, player tictactoe.GameState) tictactoe.GameState {
	playerMask := tictactoe.GameState(tictactoe.AllO)
	opponentCorners := Xcorners
	if player == tictactoe.X {
		playerMask = tictactoe.AllX
		opponentCorners = Ocorners
	}
	// if the other player played a corner, call minimax
	if gs&(opponentCorners) != 0 {
		return bots.MiniMax{}.Move(gs, player)
	}

	moves := gs.AvailableMoves()
	// Otherwise, corners, then random
	for _, m := range moves {
		if m&CornerMask != 0 {
			return m & playerMask
		}
	}

	return moves[rand.Int()%len(moves)] & playerMask
}
