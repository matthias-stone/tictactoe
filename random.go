package tictactoe

import "math/rand"

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
