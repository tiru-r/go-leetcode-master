package soduku_solver_37

const (
	boardSize = 9
	boxSize   = 3
	allBitsMask = (1 << boardSize) - 1 // 0b111111111
)

// solveSudoku fills the 9×9 Sudoku board in-place using backtracking with bitmask optimization.
func solveSudoku(board [][]byte) {
	solver := &sudokuSolver{board: board}
	solver.initialize()
	_ = solver.backtrack(0) // Solution guaranteed to exist per problem constraints
}

// sudokuSolver uses bitmasks to efficiently track digit constraints
type sudokuSolver struct {
	board      [][]byte
	rowMask    [boardSize]uint16 // Bitmask of used digits in each row
	colMask    [boardSize]uint16 // Bitmask of used digits in each column  
	boxMask    [boardSize]uint16 // Bitmask of used digits in each 3x3 box
	emptyCells [][2]int          // Pre-computed list of empty cell positions
}

// initialize scans the board to set up bitmasks and collect empty cells
func (s *sudokuSolver) initialize() {
	// Pre-allocate with estimated capacity (typical Sudoku has ~40-60 empty cells)
	s.emptyCells = make([][2]int, 0, 50)
	
	for row := range boardSize {
		for col := range boardSize {
			if s.board[row][col] == '.' {
				s.emptyCells = append(s.emptyCells, [2]int{row, col})
			} else {
				// Convert ASCII digit to number and set constraints
				digit := uint16(s.board[row][col] - '0')
				s.updateConstraints(row, col, digit, true)
			}
		}
	}
}

// updateConstraints sets or unsets digit constraints using bitmasks
func (s *sudokuSolver) updateConstraints(row, col int, digit uint16, enable bool) {
	digitBit := uint16(1) << (digit - 1)
	boxIndex := (row/boxSize)*boxSize + col/boxSize
	
	if enable {
		// Set bit to mark digit as used
		s.rowMask[row] |= digitBit
		s.colMask[col] |= digitBit
		s.boxMask[boxIndex] |= digitBit
	} else {
		// Clear bit to mark digit as available
		s.rowMask[row] &^= digitBit
		s.colMask[col] &^= digitBit
		s.boxMask[boxIndex] &^= digitBit
	}
}

// getValidDigits returns bitmask of digits that can be placed at (row, col)
func (s *sudokuSolver) getValidDigits(row, col int) uint16 {
	boxIndex := (row/boxSize)*boxSize + col/boxSize
	// XOR with all bits to get available digits (flip used->available)
	usedDigits := s.rowMask[row] | s.colMask[col] | s.boxMask[boxIndex]
	return allBitsMask &^ usedDigits
}

// backtrack recursively fills empty cells using constraint propagation
func (s *sudokuSolver) backtrack(position int) bool {
	// Base case: all empty cells filled successfully
	if position == len(s.emptyCells) {
		return true
	}
	
	// Get current empty cell coordinates
	row, col := s.emptyCells[position][0], s.emptyCells[position][1]

	// Get bitmask of valid digits for this position
	validDigits := s.getValidDigits(row, col)
	
	// Try each possible digit (1-9)
	for digit := uint16(1); digit <= boardSize; digit++ {
		digitBit := uint16(1) << (digit - 1)
		
		// Skip if digit is not valid for this position
		if validDigits&digitBit == 0 {
			continue
		}

		// Place digit and update constraints
		s.board[row][col] = byte('0' + digit)
		s.updateConstraints(row, col, digit, true)
		
		// Recursively try to fill remaining cells
		if s.backtrack(position + 1) {
			return true // Solution found
		}
		
		// Backtrack: undo placement and constraints
		s.updateConstraints(row, col, digit, false)
		s.board[row][col] = '.'
	}
	
	return false // No valid digit found for this position
}
