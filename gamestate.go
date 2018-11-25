package tictactoe

type GameState uint32

const (
	Empty GameState = iota
	O
	X
)

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

const (
	TopMask       = 0x3F
	MiddleMask    = 0x530
	BottomMask    = 0x3F000
	LeftMask      = 0x30C3
	CenterMask    = 0xC30C
	RightMask     = 0x30C30
	AngleDownMask = 0x30303
	AngleUpMask   = 0x3330
	allO          = 0x15555
	allX          = 0x2AAAA
)

func (gs GameState) Winner() GameState {
	switch {
	case gs&TopMask == allO&TopMask:
		return O
	case gs&TopMask == allX&TopMask:
		return X
	case gs&MiddleMask == allO&MiddleMask:
		return O
	case gs&MiddleMask == allX&MiddleMask:
		return X
	case gs&BottomMask == allO&BottomMask:
		return O
	case gs&BottomMask == allX&BottomMask:
		return X
	case gs&LeftMask == allO&LeftMask:
		return O
	case gs&LeftMask == allX&LeftMask:
		return X
	case gs&CenterMask == allO&CenterMask:
		return O
	case gs&CenterMask == allX&CenterMask:
		return X
	case gs&RightMask == allO&RightMask:
		return O
	case gs&RightMask == allX&RightMask:
		return X
	case gs&AngleDownMask == allO&AngleDownMask:
		return O
	case gs&AngleDownMask == allX&AngleDownMask:
		return X
	case gs&AngleUpMask == allO&AngleUpMask:
		return O
	case gs&AngleUpMask == allX&AngleUpMask:
		return X
	default:
		return Empty
	}
}
