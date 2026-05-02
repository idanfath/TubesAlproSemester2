package mindflow

import "fmt"

type Table struct {
	Header []string
	Rows   [][]any
}

func printTable(table Table) {
	var colLengths []int = getColLongestString(table)
	var colAmount = len(table.Header)
	var rowAmount = len(table.Rows)
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
		s += fmt.Sprintf("%-*s", colLengths[i], centeredString(table.Header[i], colLengths[i])) + "|"
	}
	fmt.Println(center(s))
}

func printTableContentRows(table Table, colAmount int, rowAmount int, colLengths []int) {
	var row, col int
	var s string
	for row = 0; row < rowAmount; row++ {
		s = "|"
		for col = 0; col < colAmount; col++ {
			// -1 karna kasih spasi di awal
			s += fmt.Sprintf(" %-*s", colLengths[col]-1, toString(table.Rows[row][col]))
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
	var colAmount int = len(table.Header)
	var rowAmount int = len(table.Rows)
	var columnsLength []int = make([]int, colAmount)
	var col, row int
	var currentRowsColLength int
	for col = 0; col < colAmount; col++ {

		columnsLength[col] = len(table.Header[col])
		for row = 0; row < rowAmount; row++ {
			currentRowsColLength = len(toString(table.Rows[row][col]))
			if currentRowsColLength > columnsLength[col] {
				columnsLength[col] = currentRowsColLength
			}
		}
		// nambahin spacing dikit, + weird ahh stuff (biar perfect center headernya)
		columnsLength[col] = columnsLength[col] + 2
		if (columnsLength[col]%2 != 0 && len(table.Header[col])%2 == 0) || (columnsLength[col]%2 == 0 && len(table.Header[col])%2 == 1) {
			columnsLength[col] = columnsLength[col] + 1
		}
	}
	return columnsLength
}
