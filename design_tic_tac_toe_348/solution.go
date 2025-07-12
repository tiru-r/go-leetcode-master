package design_tic_tac_toe_348

type TicTacToe struct {
	row []int
	col []int
	d1  int // main diagonal (i == j)
	d2  int // anti-diagonal (i+j == n-1)
	n   int
}

func Constructor(n int) TicTacToe {
	return TicTacToe{
		row: make([]int, n),
		col: make([]int, n),
		n:   n,
	}
}

// Move returns the winner after the move (0, 1, 2).
func (t *TicTacToe) Move(r, c, player int) int {
	delta := 1
	if player == 2 {
		delta = -1
	}

	t.row[r] += delta
	t.col[c] += delta
	if r == c {
		t.d1 += delta
	}
	if r+c == t.n-1 {
		t.d2 += delta
	}

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	target := t.n
	if abs(t.row[r]) == target || abs(t.col[c]) == target ||
		abs(t.d1) == target || abs(t.d2) == target {
		return player
	}
	return 0
}
