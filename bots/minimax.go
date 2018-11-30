package bots

import (
	"fmt"
	"math/rand"

	"github.com/matthias-stone/tictactoe"
)

type MiniMax struct{}

func (mm MiniMax) Name() string { return "minimax" }

func (mm MiniMax) Move(gs, player tictactoe.GameState) tictactoe.GameState {
	playerMask, opponentMask := tictactoe.GameState(tictactoe.AllO), tictactoe.GameState(tictactoe.AllX)
	opponent := tictactoe.X
	if player == tictactoe.X {
		playerMask, opponentMask = opponentMask, playerMask
		opponent = tictactoe.O
	}
	moves := gs.AvailableMoves()
	values := make([]int, len(moves))
	for i, m := range moves {
		newMove := gs | (m & playerMask)
		if newMove.Winner() == player {
			return playerMask & m
		}
		values[i] = mm.recurse(newMove, opponent, player, opponentMask, playerMask)
		if values[i] == -1 {
			return playerMask & m
		}
	}
	for i, v := range values {
		if v == 0 {
			return moves[i] & playerMask
		}
	}

	return moves[rand.Int()%len(moves)] & playerMask
}

func (mm MiniMax) recurse(gs, player, opponent, playerMask, opponentMask tictactoe.GameState) int {
	moves := gs.AvailableMoves()
	if len(moves) == 0 {
		return 0
	}
	// If we can win, do so!
	for _, m := range moves {
		if ((playerMask & m) | gs).Winner() == player {
			return 1
		}
	}

	values := make([]int, len(moves))
	for i, m := range moves {
		values[i] = mm.recurse(gs|(m&playerMask), opponent, player, opponentMask, playerMask)
		if values[i] == -1 {
			return 1
		}
	}
	for _, v := range values {
		if v == 0 {
			return 0
		}
	}
	return -1
}

// MiniMaxSometimesRandom will play randomly the fraction specified.
// 0 = never, 1 = always.
type MiniMaxSometimesRandom float64

func (mmsr MiniMaxSometimesRandom) Name() string {
	return fmt.Sprintf("minimax-sometimes-random-%0.3f", mmsr)
}

func (mmsr MiniMaxSometimesRandom) Move(gs, player tictactoe.GameState) tictactoe.GameState {
	playerMask := tictactoe.GameState(tictactoe.AllO)
	if player == tictactoe.X {
		playerMask = tictactoe.GameState(tictactoe.AllX)
	}
	if float64(mmsr) > rand.Float64() {
		moves := gs.AvailableMoves()
		return moves[rand.Int()%len(moves)] & playerMask
	}
	return MiniMax{}.Move(gs, player)
}
