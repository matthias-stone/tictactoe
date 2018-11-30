package bots

import (
	"math/rand"

	"github.com/matthias-stone/tictactoe"
)

// Random makes a random move.
type Random struct{}

func (r Random) Name() string { return "random" }

func (r Random) Move(gs, player tictactoe.GameState) tictactoe.GameState {
	playerMask := tictactoe.GameState(tictactoe.AllO)
	if player == tictactoe.X {
		playerMask = tictactoe.AllX
	}
	moves := gs.AvailableMoves()
	return moves[rand.Int()%len(moves)] & playerMask
}

// RandomOpportunistic tries to win with the current move, failing that play randomly.
type RandomOpportunistic struct{}

func (ro RandomOpportunistic) Name() string { return "random-opportunistic" }

func (ro RandomOpportunistic) Move(gs, player tictactoe.GameState) tictactoe.GameState {
	playerMask := tictactoe.GameState(tictactoe.AllO)
	if player == tictactoe.X {
		playerMask = tictactoe.AllX
	}
	moves := gs.AvailableMoves()
	for _, m := range moves {
		if ((playerMask & m) | gs).Winner() == player {
			return m & playerMask
		}
	}
	return moves[rand.Int()%len(moves)] & playerMask
}

// RandomSpoiler tries to take a winning opportunity from the other player with the current move, failing that play randomly.
type RandomSpoiler struct{}

func (ro RandomSpoiler) Name() string { return "random-spoiler" }

func (ro RandomSpoiler) Move(gs, player tictactoe.GameState) tictactoe.GameState {
	playerMask, opponentMask := tictactoe.GameState(tictactoe.AllO), tictactoe.GameState(tictactoe.AllX)
	if player == tictactoe.X {
		playerMask, opponentMask = opponentMask, playerMask
	}
	moves := gs.AvailableMoves()
	for _, m := range moves {
		if ((opponentMask & m) | gs).Winner() == player {
			return m & playerMask
		}
	}
	return moves[rand.Int()%len(moves)] & playerMask
}

// RandomOpportunisticSpoiler tries to win with the current move, failing that tries to take a winning opportunity from the other player with the current move, failing that play randomly.
type RandomOpportunisticSpoiler struct{}

func (ro RandomOpportunisticSpoiler) Name() string { return "random-opportunistic-spoiler" }

func (ro RandomOpportunisticSpoiler) Move(gs, player tictactoe.GameState) tictactoe.GameState {
	playerMask, opponentMask := tictactoe.GameState(tictactoe.AllO), tictactoe.GameState(tictactoe.AllX)
	if player == tictactoe.X {
		playerMask, opponentMask = opponentMask, playerMask
	}
	moves := gs.AvailableMoves()
	for _, m := range moves {
		if ((playerMask & m) | gs).Winner() == player {
			return m & playerMask
		}
	}
	for _, m := range moves {
		if ((opponentMask & m) | gs).Winner() == player {
			return m & playerMask
		}
	}
	return moves[rand.Int()%len(moves)] & playerMask
}
