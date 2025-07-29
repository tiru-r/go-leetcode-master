package design_sql_2408

import (
	"slices"
	"sort"
	"strconv"
	"strings"
)

type SQL struct {
	tables map[string]*Table
}

type Table struct {
	columns int
	rows    map[int][]string
	nextID  int
}

func NewSQL(names []string, columns []int) *SQL {
	tables := make(map[string]*Table, len(names))
	for i, name := range names {
		tables[name] = &Table{
			columns: columns[i],
			rows:    make(map[int][]string),
			nextID:  1,
		}
	}
	return &SQL{tables: tables}
}

func (sql *SQL) Insert(tableName string, row []string) bool {
	table, exists := sql.tables[tableName]
	if !exists || len(row) != table.columns {
		return false
	}
	id := table.nextID
	table.nextID++
	table.rows[id] = slices.Clone(row)
	return true
}

func (sql *SQL) Delete(tableName string, rowID int) {
	if table, exists := sql.tables[tableName]; exists {
		delete(table.rows, rowID)
	}
}

func (sql *SQL) Select(tableName string, rowID, columnID int) string {
	table, exists := sql.tables[tableName]
	if !exists {
		return "<null>"
	}
	row, exists := table.rows[rowID]
	if !exists || columnID < 1 || columnID > len(row) {
		return "<null>"
	}
	return row[columnID-1]
}

func (sql *SQL) Export(tableName string) []string {
	table, exists := sql.tables[tableName]
	if !exists || len(table.rows) == 0 {
		return []string{}
	}
	
	// Get sorted row IDs for deterministic output
	ids := make([]int, 0, len(table.rows))
	for id := range table.rows {
		ids = append(ids, id)
	}
	sort.Ints(ids)
	
	result := make([]string, 0, len(table.rows))
	for _, id := range ids {
		row := table.rows[id]
		var builder strings.Builder
		builder.WriteString(strconv.Itoa(id))
		for _, cell := range row {
			builder.WriteByte(',')
			builder.WriteString(cell)
		}
		result = append(result, builder.String())
	}
	return result
}
