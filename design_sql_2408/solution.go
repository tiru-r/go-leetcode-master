package design_sql_2408

// SQL implements the tiny database required by LeetCode 2408.
type SQL struct{ tbl map[string]*table }

type table struct {
	colCnt int
	rows   map[int][]string // rowID -> cells
}

// Constructor creates the tables.
func Constructor(names []string, cols []int) SQL {
	tbl := make(map[string]*table, len(names))
	for i, name := range names {
		tbl[name] = &table{colCnt: cols[i], rows: make(map[int][]string)}
	}
	return SQL{tbl: tbl}
}

// Ins inserts a row and returns success.
func (q *SQL) Ins(name string, row []string) bool {
	t, ok := q.tbl[name]
	if !ok || len(row) != t.colCnt {
		return false
	}
	id := len(t.rows) + 1 // LeetCode guarantees unique ids
	t.rows[id] = row
	return true
}

// Rmv removes a row (idempotent).
func (q *SQL) Rmv(name string, id int) {
	if t, ok := q.tbl[name]; ok {
		delete(t.rows, id)
	}
}

// Sel returns the requested cell or "<null>".
func (q *SQL) Sel(name string, id, col int) string {
	t, ok := q.tbl[name]
	if !ok {
		return "<null>"
	}
	row, ok := t.rows[id]
	if !ok || col < 1 || col > len(row) {
		return "<null>"
	}
	return row[col-1]
}

// Exp returns every row in the format "id,col1,col2,…".
func (q *SQL) Exp(name string) []string {
	t, ok := q.tbl[name]
	if !ok {
		return nil
	}
	out := make([]string, 0, len(t.rows))
	for id, row := range t.rows {
		// Pre-size buffer: id digits + commas + all column bytes
		size := len(row)
		for _, c := range row {
			size += len(c)
		}
		buf := make([]byte, 0, size+1) // +1 for id comma
		buf = appendInt(buf, id)
		for _, c := range row {
			buf = append(buf, ',')
			buf = append(buf, c...)
		}
		out = append(out, string(buf))
	}
	return out
}

// appendInt appends the decimal representation of i to buf.
func appendInt(buf []byte, i int) []byte {
	if i == 0 {
		return append(buf, '0')
	}
	// Write digits in reverse then flip.
	start := len(buf)
	for i > 0 {
		buf = append(buf, byte(i%10+'0'))
		i /= 10
	}
	// Reverse in place.
	for l, r := start, len(buf)-1; l < r; l, r = l+1, r-1 {
		buf[l], buf[r] = buf[r], buf[l]
	}
	return buf
}
