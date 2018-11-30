package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/matthias-stone/tictactoe"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	rl := tictactoe.NewReinforcementLearning()

	trainingOpponent := Corners{}

	printHeader()
	testResults(rl)
	train(rl, trainingOpponent, 1000)
	testResults(rl)
	train(rl, trainingOpponent, 10000)
	testResults(rl)
	train(rl, trainingOpponent, 20000)
	train(rl, tictactoe.MiniMaxSometimesRandom(0.5), 20000)
	testResults(rl)
}

func train(rl *tictactoe.ReinforcementLearning, second tictactoe.Player, rounds int) {
	var p1, p2 tictactoe.Player
loop:
	for j := 0; j < rounds; j++ {
		rlt := tictactoe.NewReinforcementLearningTrainer(rl)
		active, inactive := tictactoe.X, tictactoe.O
		currentPlayer, nextPlayer := 1, 2
		if j&1 == 1 {
			p1, p2 = rlt, second
		} else {
			p1, p2 = second, rlt
			currentPlayer, nextPlayer = nextPlayer, currentPlayer
		}

		gs := tictactoe.GameState(0)
		for i := 0; i < 9; i++ {
			m := p1.Move(gs, active)
			gs |= m
			if gs.Winner() != tictactoe.Empty {
				_, ok := p1.(*tictactoe.ReinforcementLearningTrainer)
				rlt.RecordWin(ok)
				continue loop
			}
			p1, p2 = p2, p1
			active, inactive = inactive, active
			currentPlayer, nextPlayer = nextPlayer, currentPlayer
		}
	}
}

func compete(first, second tictactoe.Player, rounds int) [3]int {
	var p1, p2 tictactoe.Player
	results := [3]int{}
loop:
	for j := 0; j < rounds; j++ {
		active, inactive := tictactoe.X, tictactoe.O
		currentPlayer, nextPlayer := 1, 2
		if j&1 == 1 {
			p1, p2 = first, second
		} else {
			p1, p2 = second, first
			currentPlayer, nextPlayer = nextPlayer, currentPlayer
		}

		gs := tictactoe.GameState(0)
		for i := 0; i < 9; i++ {
			m := p1.Move(gs, active)
			gs |= m
			if gs.Winner() != tictactoe.Empty {
				results[currentPlayer] += 1
				continue loop
			}
			p1, p2 = p2, p1
			active, inactive = inactive, active
			currentPlayer, nextPlayer = nextPlayer, currentPlayer
		}
		results[0] += 1
	}
	return results
}

func printHeader() {
	fmt.Printf(" rl v corners     rl v minimax    corners v minimax\n")
	fmt.Printf(" win/loss/draw    win/loss/draw    win/loss/draw\n")
}

func testResults(rl *tictactoe.ReinforcementLearning) {
	rounds := 1000
	var (
		r1 = compete(rl, Corners{}, rounds)
		r2 = compete(rl, tictactoe.MiniMax{}, rounds)
		r3 = compete(Corners{}, tictactoe.MiniMax{}, rounds)
	)

	fmt.Printf("%4d %4d %4d   %4d %4d %4d   %4d %4d %4d\n",
		r1[1], r1[2], r1[0],
		r2[1], r2[2], r2[0],
		r3[1], r3[2], r3[0],
	)
}
