package tictactoe

import "math/rand"

// Random makes a random move.
type Random struct{}

func (r Random) Name() string { return "random" }

func (r Random) Move(gs, player GameState) GameState {
	playerMask := GameState(allO)
	if player == X {
		playerMask = allX
	}
	moves := gs.AvailableMoves()
	return moves[rand.Int()%len(moves)] & playerMask
}

// RandomOpportunistic tries to win with the current move, failing that play randomly.
type RandomOpportunistic struct{}

func (ro RandomOpportunistic) Name() string { return "random-opportunistic" }

func (ro RandomOpportunistic) Move(gs, player GameState) GameState {
	playerMask := GameState(allO)
	if player == X {
		playerMask = allX
	}
	moves := gs.AvailableMoves()
	for _, m := range gs.AvailableMoves() {
		if ((playerMask & m) | gs).Winner() == player {
			return m & playerMask
		}
	}
	return moves[rand.Int()%len(moves)] & playerMask
}

// RandomSpoiler tries to take a winning opportunity from the other player with the current move, failing that play randomly.
type RandomSpoiler struct{}

func (ro RandomSpoiler) Name() string { return "random-spoiler" }

func (ro RandomSpoiler) Move(gs, player GameState) GameState {
	playerMask, opponentMask := GameState(allO), GameState(allX)
	if player == X {
		playerMask, opponentMask = opponentMask, playerMask
	}
	moves := gs.AvailableMoves()
	for _, m := range gs.AvailableMoves() {
		if ((opponentMask & m) | gs).Winner() == player {
			return m & playerMask
		}
	}
	return moves[rand.Int()%len(moves)] & playerMask
}

// RandomOpportunisticSpoiler tries to win with the current move, failing that tries to take a winning opportunity from the other player with the current move, failing that play randomly.
type RandomOpportunisticSpoiler struct{}

func (ro RandomOpportunisticSpoiler) Name() string { return "random-opportunistic-spoiler" }

func (ro RandomOpportunisticSpoiler) Move(gs, player GameState) GameState {
	playerMask, opponentMask := GameState(allO), GameState(allX)
	if player == X {
		playerMask, opponentMask = opponentMask, playerMask
	}
	moves := gs.AvailableMoves()
	for _, m := range gs.AvailableMoves() {
		if ((playerMask & m) | gs).Winner() == player {
			return m & playerMask
		}
	}
	for _, m := range gs.AvailableMoves() {
		if ((opponentMask & m) | gs).Winner() == player {
			return m & playerMask
		}
	}
	return moves[rand.Int()%len(moves)] & playerMask
}
