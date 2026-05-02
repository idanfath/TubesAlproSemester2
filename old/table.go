package main

import "fmt"

type Table struct {
	header []string
	rows   [][]any
}
type TablesStorage struct {
	name  string
	table Table
}

var tablesStorage []TablesStorage

func printTable(table Table) {
	var colLengths []int = getColLongestString(table)
	var colAmount = len(table.header)
	var rowAmount = len(table.rows)
	printTableBorderHorizontal(colAmount, colLengths)
	printTableContentHeader(table, colAmount, colLengths)
	printTableBorderHorizontal(colAmount, colLengths)
	printTableContentRows(table, colAmount, rowAmount, colLengths)
	printTableBorderHorizontal(colAmount, colLengths)
}

func printTableContentHeader(table Table, colAmount int, colLengths []int) {
	var i int
	var s string = "|"
	for i = 0; i < colAmount; i++ {
		s += fmt.Sprintf("%-*s", colLengths[i], centeredString(table.header[i], colLengths[i])) + "|"
	}
	fmt.Println(center(s))
}

func printTableContentRows(table Table, colAmount int, rowAmount int, colLengths []int) {
	var row, col int
	var s string
	if rowAmount == 0 {
		rowAmount = 1
		var x []any
		var i int
		for i = 0; i < colAmount; i++ {
			x = append(x, "")
		}
		table.rows = [][]any{x}
	}
	for row = 0; row < rowAmount; row++ {
		s = "|"
		for col = 0; col < colAmount; col++ {
			// -1 karna kasih spasi di awal
			s += fmt.Sprintf(" %-*s", colLengths[col]-1, toString(table.rows[row][col]))
			s += "|"
		}
		fmt.Println(center(s))
	}
}

func printTableBorderHorizontal(colAmount int, colLengths []int) {
	var i int
	var s string = "+"
	for i = 0; i < colAmount; i++ {
		s = s + repeat("-", colLengths[i]) + "+"
	}
	fmt.Println(center(s))
}

func getColLongestString(table Table) []int {
	var colAmount int = len(table.header)
	var rowAmount int = len(table.rows)
	var columnsLength []int = make([]int, colAmount)
	var col, row int
	var currentRowsColLength int
	for col = 0; col < colAmount; col++ {
		columnsLength[col] = len(table.header[col])
		for row = 0; row < rowAmount; row++ {
			currentRowsColLength = len(toString(table.rows[row][col]))
			if currentRowsColLength > columnsLength[col] {
				columnsLength[col] = currentRowsColLength
			}
		}
		// nambahin spacing dikit, + weird ahh stuff (biar perfect center headernya)
		columnsLength[col] = columnsLength[col] + 2
		if (columnsLength[col]%2 != 0 && len(table.header[col])%2 == 0) || (columnsLength[col]%2 == 0 && len(table.header[col])%2 == 1) {
			columnsLength[col] = columnsLength[col] + 1
		}
	}
	return columnsLength
}

func getColumns(table Table) []string {
	var colAmount int = len(table.header)
	var columns []string
	var i int
	for i = 0; i < colAmount; i++ {
		columns = append(columns, table.header[i])
	}
	return columns
}
func getColumnsValue(table Table, columnIndex int) []any {
	var rowAmount int = len(table.rows)
	var values []any
	var i int
	for i = 0; i < rowAmount; i++ {
		values = append(values, table.rows[i][columnIndex])
	}
	return values
}
func newRow(table Table, values []any) {
	table.rows = append(table.rows, values)
}
func getTableByName(name string) Table {
	var i int
	var tablesAmount int = len(tablesStorage)
	for i = 0; i < tablesAmount; i++ {
		if tablesStorage[i].name == name {
			return tablesStorage[i].table
		}
	}
	return Table{}
}

// temp
func find(table Table, value any, col_index int) bool {
	var values = getColumnsValue(table, col_index)
	var i int
	for i = 0; i < len(values); i++ {
		if values[i] == value {
			return true
		}
	}
	return false
}

func getTableIndexFromTableStorage(tableName string) int {
	var i int
	for i = 0; i < len(tablesStorage); i++ {
		if tablesStorage[i].name == tableName {
			return i
		}
	}
	return 0
}

func appendToTable(tableName string, values []any) {
	var table Table = getTableByName(tableName)
	table.rows = append(table.rows, values)
	tablesStorage[getTableIndexFromTableStorage(tableName)].table = table
}

func removeRow(tableName string, rowIndex int) {
	var table Table = getTableByName(tableName)
	var newRows [][]any
	var i int
	for i = 0; i < len(table.rows); i++ {
		if i == rowIndex {
			continue
		}
		newRows = append(newRows, table.rows[i])
	}
	tablesStorage[getTableIndexFromTableStorage(tableName)].table.rows = newRows
}
