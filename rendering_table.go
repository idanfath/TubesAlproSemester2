package main

import "fmt"

// print table
func printTable(table Table) {
	var lengthsCol []int = columnLengths(table)
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
		var x []string
		for i = 0; i < len(colLengths); i++ {
			x = append(x, "")
		}
		table.rows = [][]string{x}
	}

	// LOGIC UTAMA print baris per baris
	for i = 0; i < rowAmount; i++ {
		// build stringnya dari kolom kiri ke kanan
		s = "|"
		for j = 0; j < len(colLengths); j++ {
			// -2 karena kasih spasi di awal dan di akhir agar tidak nempel dinding
			var width = colLengths[j] - 2
			var text = truncate(toString(table.rows[i][j]), width)
			// | stringxxxx | dimana x itu spasi buat isi sisa spacenya, misal width 10
			s += fmt.Sprintf(" %s ", text+repeat(" ", width-len(text)))
			s += "|"
		}
		fmt.Println(cen(s))
	}
}
