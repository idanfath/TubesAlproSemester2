package main

import "fmt"

type Table struct {
	name             string
	header           []string
	rows             [][]any
	col_max_length   []int
	unique_col_index []int
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

	// FALLBACK: setidaknya print 1 row kosong, selama headernya ada, biar keliatan bentuk tabelnya
	if rowAmount == 0 {
		rowAmount = 1
		var x []any
		for i = 0; i < len(colLengths); i++ {
			x = append(x, "")
		}
		table.rows = [][]any{x}
	}

	// LOGIC UTAMA print baris per baris
	for i = 0; i < rowAmount; i++ {
		// build stringnya dari kolom kiri ke kanan
		s = "|"
		for j = 0; j < len(colLengths); j++ {
			// -1 karna kasih spasi di awal
			// TODO: tanya bapaknya boleh pake * dinamis ga
			s += fmt.Sprintf(" %-*s", colLengths[j]-1, truncate(toString(table.rows[i][j]), colLengths[j]-1))
			s += "|"
		}
		fmt.Println(cen(s))
	}
}

func colLengths(table Table) []int {
	var colAmount int = len(table.header)
	var rowAmount int = len(table.rows)
	var columnsLength []int
	var currLength int
	var col, row int
	// iterasi tiap column
	for col = 0; col < colAmount; col++ {
		columnsLength = append(columnsLength, len(toString(table.header[col]))) // mulai dari panjang header
		// iterasi tiap row di column untuk cari string terpanjang di columnya
		for row = 0; row < rowAmount; row++ {
			currLength = len(truncate(toString(table.rows[row][col]), table.col_max_length[col])) // panjang string di cell ini
			if currLength > columnsLength[col] {
				columnsLength[col] = currLength
			}
		}
		// exit iterasi row di column ini (max udah ketemu)

		// space tambahan
		columnsLength[col] = columnsLength[col] + 2
		// karna header akan di-center, harus fix paritas
		// kalu header genap, spacenya harus genap juga biar perfect center, begitu juga sebaliknya
		// basically paritasnya harus sama, kalo beda, tambahin 1 spasi lagi biar sama
		if len(table.header[col])%2 != columnsLength[col]%2 {
			columnsLength[col] = columnsLength[col] + 1
		}
	}
	return columnsLength
}
