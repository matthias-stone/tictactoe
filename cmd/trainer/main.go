package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/matthias-stone/tictactoe"
	"github.com/matthias-stone/tictactoe/bots"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	rl := bots.NewReinforcementLearning()

	printHeader()
	testResults(rl)
	train(rl, bots.MiniMaxSometimesRandom(0.5), 1000)
	testResults(rl)
	train(rl, bots.MiniMaxSometimesRandom(0.5), 10000)
	testResults(rl)
	train(rl, bots.MiniMaxSometimesRandom(0.5), 50000)
	testResults(rl)
}

func train(rl *bots.ReinforcementLearning, second tictactoe.Player, rounds int) {
	var p1, p2 tictactoe.Player
loop:
	for j := 0; j < rounds; j++ {
		rlt := bots.NewReinforcementLearningTrainer(rl)
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
				_, ok := p1.(*bots.ReinforcementLearningTrainer)
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
	fmt.Printf("     random            self         0.1 minimax        minimax\n")
	fmt.Printf(" win/loss/draw    win/loss/draw    win/loss/draw    win/loss/draw\n")
}

func testResults(rl *bots.ReinforcementLearning) {
	rounds := 1000
	var (
		r1, _ = timeCompeteAndDuration(rl, bots.RandomOpportunisticSpoiler{}, rounds)
		r2, _ = timeCompeteAndDuration(rl, rl, rounds)
		r3, _ = timeCompeteAndDuration(rl, bots.MiniMaxSometimesRandom(0.1), rounds)
		r4, _ = timeCompeteAndDuration(rl, bots.MiniMax{}, rounds)
	)

	fmt.Printf("%4d %4d %4d   %4d %4d %4d   %4d %4d %4d   %4d %4d %4d\n",
		r1[1], r1[2], r1[0],
		r2[1], r2[2], r2[0],
		r3[1], r3[2], r3[0],
		r4[1], r4[2], r4[0],
	)
}

func timeCompeteAndDuration(p1, p2 tictactoe.Player, rounds int) ([3]int, time.Duration) {
	t := time.Now()
	r := compete(p1, p2, rounds)
	return r, time.Now().Sub(t)
}
