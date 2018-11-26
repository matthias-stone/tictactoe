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
	}
	rand.Seed(time.Now().UnixNano())

	wg := sync.WaitGroup{}
	for _, p1 := range players {
		for _, p2 := range players {
			wg.Add(1)
			go func(p1, p2 tictactoe.Player) {
				r := compete(p1, p2, 1000)
				fmt.Printf("Draw: %d, %s: %d, %s: %d\n", r[0], p1.Name(), r[1], p2.Name(), r[2])
				wg.Done()
			}(p1, p2)
		}
	}
	wg.Wait()
}

func compete(first, second tictactoe.Player, rounds int) [3]int {
	var p1, p2 tictactoe.Player
	results := [3]int{}
	j := 0
loop:
	for j < rounds {
		j++
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
				goto loop
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