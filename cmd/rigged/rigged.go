package main

import (
	"math/rand"

	"github.com/matthias-stone/tictactoe"
)

// Corners tries to take the corners.
type Corners struct{}

func (r Corners) Name() string { return "random" }

const (
	allO       = 0x15555
	allX       = 0x2AAAA
	CornerMask = tictactoe.Pos1 | tictactoe.Pos3 | tictactoe.Pos7 | tictactoe.Pos9
	Xcorners   = allX & CornerMask
	Ocorners   = allO & CornerMask
)

func (r Corners) Move(gs, player tictactoe.GameState) tictactoe.GameState {
	playerMask := tictactoe.GameState(allO)
	opponentCorners := Xcorners
	if player == tictactoe.X {
		playerMask = allX
		opponentCorners = Ocorners
	}
	// if the other player played a corner, call minimax
	if gs&(opponentCorners) != 0 {
		return tictactoe.MiniMax{}.Move(gs, player)
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
