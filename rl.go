package tictactoe

import (
	"math/rand"
)

// Note that states where playing as O, and playing as X, never overlap.
type ReinforcementLearning [Pos9]float32

func NewReinforcementLearning() *ReinforcementLearning {
	rl := ReinforcementLearning{}
	for i := range rl[:] {
		rl[i] = 0.5
	}
	return &rl
}
func (rl *ReinforcementLearning) Name() string { return "reinforcement-learning" }

func (rl *ReinforcementLearning) Move(gs, player GameState) GameState {
	playerMask := GameState(allO)
	if player == X {
		playerMask = GameState(allX)
	}

	moves := gs.AvailableMoves()
	values := make([]float32, len(moves))
	for i, m := range moves {
		newMove := gs | (m & playerMask)
		values[i] = rl[newMove]
	}

	// Find the most valuable move and do that
	var moveValue float32
	move := 0
	for i, v := range values {
		if v > moveValue {
			move, moveValue = i, v
		}
	}

	return moves[move] & playerMask
}

const (
	explorationRate = 0.10
	stepSize        = 0.10
)

type ReinforcementLearningTrainer struct {
	rl       *ReinforcementLearning
	lastMove GameState
}

func NewReinforcementLearningTrainer(rl *ReinforcementLearning) *ReinforcementLearningTrainer {
	return &ReinforcementLearningTrainer{rl, Empty}
}

func (rlt *ReinforcementLearningTrainer) Name() string { return "reinforcement-learning-training" }

func (rlt *ReinforcementLearningTrainer) Move(gs, player GameState) GameState {
	playerMask := GameState(allO)
	if player == X {
		playerMask = GameState(allX)
	}
	moves := gs.AvailableMoves()

	// Pick best
	values := make([]float32, len(moves))
	for i, m := range moves {
		newMove := gs | (m & playerMask)
		values[i] = rlt.rl[newMove]
	}

	var moveValue float32
	move := 0
	for i, v := range values {
		if v > moveValue {
			move, moveValue = i, v
		}
	}

	if len(moves) > 1 && rlt.explore() {
		newMove := rand.Int() % (len(moves) - 1)
		if newMove >= move {
			newMove++
		}
		move = newMove
	} else {
		rlt.update(gs | (moves[move] & playerMask))
	}
	// Maybe explore?

	return moves[move] & playerMask
}

func (rlt *ReinforcementLearningTrainer) RecordWin(win bool) {
	var result float32
	if win {
		result = 1
	}
	rlt.rl[rlt.lastMove] = rlt.rl[rlt.lastMove] + stepSize*(result-rlt.rl[rlt.lastMove])
}

func (rlt *ReinforcementLearningTrainer) update(nextState GameState) {
	rlt.rl[rlt.lastMove] = rlt.rl[rlt.lastMove] + stepSize*(rlt.rl[nextState]-rlt.rl[rlt.lastMove])
	rlt.lastMove = nextState
}

func (rlt *ReinforcementLearningTrainer) explore() bool {
	return explorationRate > rand.Float32()
}
