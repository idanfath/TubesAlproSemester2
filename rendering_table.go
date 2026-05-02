package main

import "fmt"

type Table struct {
	name   string
	header []string
	rows   [][]any
}
type TableIndex struct {
	rowIndex int
	colIndex int
}

var tables []Table

func printT(table Table) {
	var lengthsCol []int = colLengths(table)
	var rowAmount = len(table.rows)
	tableBorderCorner(lengthsCol)
	tableContentHeader(table, lengthsCol)
	tableBorderCorner(lengthsCol)
	tableContentRows(table, rowAmount, lengthsCol)
	tableBorderCorner(lengthsCol)
}

func tableBorderCorner(colLengths []int) {
	if len(colLengths) == 0 {
		return
	}
	var i int
	var s string = "+"
	for i = 0; i < len(colLengths); i++ {
		s += repeat("-", colLengths[i]) + "+"
	}
	fmt.Println(cen(s))
}

func tableContentHeader(table Table, colLengths []int) {
	if len(table.header) == 0 {
		return
	}
	var i int
	var s string = "|"
	for i = 0; i < len(colLengths); i++ {
		s += fmt.Sprintf("%-*s", colLengths[i], centerString(table.header[i], colLengths[i])) + "|"
	}
	fmt.Println(cen(s))
}

func tableContentRows(table Table, rowAmount int, colLengths []int) {
	var i, j int
	var s string
	if len(table.header) == 0 {
		return
	}
	// setidaknya print 1 row kosong, selama headernya ada, biar keliatan bentuk tabelnya
	if rowAmount == 0 {
		rowAmount = 1
		var x []any
		for i = 0; i < len(colLengths); i++ {
			x = append(x, "")
		}
		table.rows = [][]any{x}
	}
	// print baris per baris
	for i = 0; i < rowAmount; i++ {
		// build stringnya dari kolom kiri ke kanan
		s = "|"
		for j = 0; j < len(colLengths); j++ {
			// -1 karna kasih spasi di awal
			s += fmt.Sprintf(" %-*s", colLengths[j]-1, toString(table.rows[i][j]))
			s += "|"
		}
		fmt.Println(cen(s))
	}
}
