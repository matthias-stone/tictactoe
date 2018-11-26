package tictactoe

import "testing"

func BenchmarkNewReinforcementLearning(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewReinforcementLearning()
	}
}
