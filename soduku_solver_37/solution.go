package soduku_solver_37

type Solver struct {
	board      [][]byte
	rowMask    [9]uint16
	colMask    [9]uint16
	boxMask    [9]uint16
	emptyCells [][2]int
}

func solveSudoku(board [][]byte) {
	s := &Solver{board: board}
	s.init()
	s.solve(0)
}

func (s *Solver) init() {
	s.emptyCells = make([][2]int, 0, 81)
	for r := range 9 {
		for c := range 9 {
			if s.board[r][c] == '.' {
				s.emptyCells = append(s.emptyCells, [2]int{r, c})
			} else {
				d := uint16(s.board[r][c] - '0')
				s.set(r, c, d, true)
			}
		}
	}
}

func (s *Solver) set(r, c int, d uint16, on bool) {
	bit := uint16(1) << (d - 1)
	bi := (r/3)*3 + c/3
	if on {
		s.rowMask[r] |= bit
		s.colMask[c] |= bit
		s.boxMask[bi] |= bit
	} else {
		s.rowMask[r] &^= bit
		s.colMask[c] &^= bit
		s.boxMask[bi] &^= bit
	}
}

func (s *Solver) candidates(r, c int) uint16 {
	bi := (r/3)*3 + c/3
	return (1<<9 - 1) &^ (s.rowMask[r] | s.colMask[c] | s.boxMask[bi])
}

func (s *Solver) solve(pos int) bool {
	if pos >= len(s.emptyCells) {
		return true
	}

	r, c := s.emptyCells[pos][0], s.emptyCells[pos][1]
	cands := s.candidates(r, c)
	if cands == 0 {
		return false
	}

	for d := uint16(1); d <= 9; d++ {
		if cands&(1<<(d-1)) == 0 {
			continue
		}

		s.board[r][c] = byte('0' + d)
		s.set(r, c, d, true)

		if s.solve(pos+1) {
			return true
		}

		s.set(r, c, d, false)
		s.board[r][c] = '.'
	}
	return false
}
