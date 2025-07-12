package design_sql_2408

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDesignSQL(t *testing.T) {
	t.Run("basic flow", func(t *testing.T) {
		sql := Constructor(
			[]string{"users", "orders"},
			[]int{3, 2},
		)

		assert.True(t, sql.Ins("users", []string{"alice", "25", "NY"}))
		assert.True(t, sql.Ins("users", []string{"bob", "30", "CA"}))
		assert.True(t, sql.Ins("orders", []string{"100", "2025-07-12"}))

		// select existing data
		assert.Equal(t, "alice", sql.Sel("users", 1, 1))
		assert.Equal(t, "30", sql.Sel("users", 2, 2))
		assert.Equal(t, "2025-07-12", sql.Sel("orders", 1, 2))

		// invalid selects
		assert.Equal(t, "<null>", sql.Sel("users", 99, 1))
		assert.Equal(t, "<null>", sql.Sel("users", 1, 0))
		assert.Equal(t, "<null>", sql.Sel("unknown", 1, 1))

		// export
		exp := sql.Exp("users")
		assert.Len(t, exp, 2)
		assert.Contains(t, exp, "1,alice,25,NY")
		assert.Contains(t, exp, "2,bob,30,CA")

		// remove & verify gone
		sql.Rmv("users", 1)
		assert.Equal(t, "<null>", sql.Sel("users", 1, 1))
		assert.Len(t, sql.Exp("users"), 1)
	})

	t.Run("edge cases", func(t *testing.T) {
		sql := Constructor([]string{"t"}, []int{1})

		// bad insert (wrong column count)
		assert.False(t, sql.Ins("t", []string{}))

		// empty export
		assert.Empty(t, sql.Exp("t"))
	})
}
