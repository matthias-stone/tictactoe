package tictactoe

type GameState uint32

const O = 1
const X = 2

func (gs GameState) String() string {
	s := make([]byte, 12)
	for i := range s {
		if i&0x3 == 3 {
			s[i] = '\n'
		} else {
			switch gs & 3 {
			case O:
				s[i] = 'O'
			case X:
				s[i] = 'X'
			default:
				s[i] = '-'
			}
			gs = gs >> 2
		}
	}

	return string(s)
}
