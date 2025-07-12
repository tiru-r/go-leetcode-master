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
	// Pre-compute empty cells and initialize masks
	s.emptyCells = make([][2]int, 0, 81)
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
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

// Constraint propagation: returns false if conflict found
func (s *Solver) propagate(r, c int, d uint16) bool {
	// Check related cells for conflicts
	for _, pos := range s.emptyCells {
		if pos[0] == r || pos[1] == c || ((pos[0]/3)*3+pos[1]/3) == (r/3)*3+c/3 {
			if pos[0] == r && pos[1] == c {
				continue
			}
			cands := s.candidates(pos[0], pos[1])
			if cands == 0 {
				return false
			}
		}
	}
	return true
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

	// Try candidates in order of increasing value
	for d := uint16(1); d <= 9; d++ {
		if cands&(1<<(d-1)) == 0 {
			continue
		}

		s.board[r][c] = byte('0' + d)
		s.set(r, c, d, true)

		if s.propagate(r, c, d) && s.solve(pos+1) {
			return true
		}

		s.set(r, c, d, false)
		s.board[r][c] = '.'
	}
	return false
}
