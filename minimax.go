package tictactoe

import (
	"math/rand"
)

type MiniMax struct{}

func (mm MiniMax) Name() string { return "minimax" }

func (mm MiniMax) Move(gs, player GameState) GameState {
	playerMask, opponentMask := GameState(allO), GameState(allX)
	opponent := X
	if player == X {
		playerMask, opponentMask = opponentMask, playerMask
		opponent = O
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

func (mm MiniMax) recurse(gs, player, opponent, playerMask, opponentMask GameState) int {
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
