package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/matthias-stone/tictactoe"
)

func main() {
	players := []tictactoe.Player{
		tictactoe.Random{},
		tictactoe.RandomOpportunistic{},
		tictactoe.RandomSpoiler{},
		tictactoe.RandomOpportunisticSpoiler{},
		tictactoe.MiniMax{},
		tictactoe.MiniMaxSometimesRandom(0.95),
		tictactoe.MiniMaxSometimesRandom(0.75),
		tictactoe.MiniMaxSometimesRandom(0.5),
	}
	rand.Seed(time.Now().UnixNano())

	fmt.Println(" win/loss/draw  Players")
	wg := sync.WaitGroup{}
	for i, p1 := range players {
		for _, p2 := range players[i:] {
			wg.Add(1)
			go func(p1, p2 tictactoe.Player) {
				r := compete(p1, p2, 100)
				fmt.Printf("%4d %4d %4d  %s vs %s\n", r[1], r[2], r[0], p1.Name(), p2.Name())
				wg.Done()
			}(p1, p2)
		}
	}
	wg.Wait()
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

// func Println(gs tictactoe.GameState) {
// 	fmt.Printf(strings.Replace(gs.String(), "\n", " ", 2))
// }
