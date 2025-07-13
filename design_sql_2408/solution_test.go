package design_sql_2408

import (
	"reflect"
	"testing"
)

func TestSQL(t *testing.T) {
	type step struct {
		op    string // "ins", "sel", "rmv", "exp"
		table string
		row   []string
		id    int
		col   int
		want  interface{} // expected result (bool, string, []string)
	}

	tests := []struct {
		name  string
		names []string
		cols  []int
		steps []step
	}{
		{
			name:  "happy path insert/select/export/delete",
			names: []string{"Users"},
			cols:  []int{2},
			steps: []step{
				{op: "ins", table: "Users", row: []string{"Alice", "25"}, want: true},
				{op: "sel", table: "Users", id: 1, col: 1, want: "Alice"},
				{op: "sel", table: "Users", id: 1, col: 2, want: "25"},
				{op: "exp", table: "Users", want: []string{"1,Alice,25"}},
				{op: "rmv", table: "Users", id: 1},
				{op: "sel", table: "Users", id: 1, col: 1, want: "<null>"},
				{op: "exp", table: "Users", want: []string{}},
			},
		},
		{
			name:  "wrong column count",
			names: []string{"A"},
			cols:  []int{3},
			steps: []step{
				{op: "ins", table: "A", row: []string{"a", "b"}, want: false},
			},
		},
		{
			name:  "non-existent table",
			names: []string{"T"},
			cols:  []int{1},
			steps: []step{
				{op: "ins", table: "X", row: []string{"x"}, want: false},
				{op: "sel", table: "X", id: 1, col: 1, want: "<null>"},
				{op: "exp", table: "X", want: ([]string)(nil)},
			},
		},
		{
			name:  "multi-row export ordering",
			names: []string{"Nums"},
			cols:  []int{1},
			steps: []step{
				{op: "ins", table: "Nums", row: []string{"C"}, want: true},
				{op: "ins", table: "Nums", row: []string{"A"}, want: true},
				{op: "ins", table: "Nums", row: []string{"B"}, want: true},
				{op: "exp", table: "Nums", want: []string{"1,C", "2,A", "3,B"}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := NewSQL(tt.names, tt.cols)
			for i, s := range tt.steps {
				switch s.op {
				case "ins":
					got := db.Insert(s.table, s.row)
					if got != s.want.(bool) {
						t.Fatalf("step %d: Insert returned %v, want %v", i, got, s.want)
					}
				case "sel":
					got := db.Select(s.table, s.id, s.col)
					if got != s.want.(string) {
						t.Fatalf("step %d: Select returned %q, want %q", i, got, s.want)
					}
				case "rmv":
					db.Delete(s.table, s.id)
				case "exp":
					got := db.Export(s.table)
					if !reflect.DeepEqual(got, s.want) {
						t.Fatalf("step %d: Export returned %v, want %v", i, got, s.want)
					}
				}
			}
		})
	}
}
