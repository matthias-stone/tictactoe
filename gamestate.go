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
	Pos1          = 0x3
	Pos2          = 0xC
	Pos3          = 0x30
	Pos4          = 0xC0
	Pos5          = 0x300
	Pos6          = 0xC00
	Pos7          = 0x3000
	Pos8          = 0xC000
	Pos9          = 0x30000
	TopMask       = Pos1 | Pos2 | Pos3 // 0x3C
	MiddleMask    = Pos4 | Pos5 | Pos6 // 0xFC0
	BottomMask    = Pos7 | Pos8 | Pos9 // 0x3F000
	LeftMask      = Pos1 | Pos4 | Pos7 // 0x30C3
	CenterMask    = Pos2 | Pos5 | Pos8 // 0xC30C
	RightMask     = Pos3 | Pos6 | Pos9 // 0x30C30
	AngleDownMask = Pos1 | Pos5 | Pos9 // 0x30303
	AngleUpMask   = Pos3 | Pos5 | Pos7 // 0x3330
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
