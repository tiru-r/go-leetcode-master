package design_tic_tac_toe_348

type TicTacToe struct {
	rows         []int
	cols         []int
	mainDiag     int
	antiDiag     int
	size         int
}

func NewTicTacToe(n int) *TicTacToe {
	return &TicTacToe{
		rows: make([]int, n),
		cols: make([]int, n),
		size: n,
	}
}

func (t *TicTacToe) Move(row, col, player int) int {
	delta := 1
	if player == 2 {
		delta = -1
	}
	
	t.rows[row] += delta
	t.cols[col] += delta
	
	if row == col {
		t.mainDiag += delta
	}
	if row+col == t.size-1 {
		t.antiDiag += delta
	}
	
	target := t.size
	if abs(t.rows[row]) == target || abs(t.cols[col]) == target ||
		abs(t.mainDiag) == target || abs(t.antiDiag) == target {
		return player
	}
	
	return 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
