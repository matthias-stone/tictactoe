package tictactoe

type Player interface {
	Move(gs, player GameState) GameState
	Name() string
}
