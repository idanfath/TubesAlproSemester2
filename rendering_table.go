// todo: optimize, mending tableContent(cols[//tes, test2]) otomatis print centered, logika truncate pas build rows

package main

import "fmt"

type Table struct {
	header     []string
	rows       [][]string
	maxLengths []int
}

// print table
func printTable(table Table, override_pagination bool) {
	var lengthsCol []int = columnLengths(table)
	var rowAmount = len(table.rows)
	tableBorderCorner(lengthsCol)
	tableContentHeader(table, lengthsCol)
	tableBorderCorner(lengthsCol)
	tableContentRows(table, rowAmount, override_pagination, lengthsCol)
	tableBorderCorner(lengthsCol)
	// pagination
	if rowAmount < App.pagination || override_pagination {
		return
	}
	tablePagination("center", lengthsCol, rowAmount)
	tableBorderCorner(lengthsCol)
}

func tablePagination(alignment string, colLengths []int, rowAmount int) {
	if len(colLengths) == 0 {
		return
	}
	var content string = fmt.Sprintf("Page %d of %d (%d data)", Temp.table_page, rowAmount/App.pagination+ifi(rowAmount%App.pagination == 0, 0, 1), rowAmount)
	var innerWidth int = sum(colLengths) + (len(colLengths) - 1)
	var text string = truncate(toString(content), innerWidth)
	var pad int = innerWidth - len(text)

	var s string = "|"
	switch alignment {
	case "right":
		s += repeat(" ", pad) + text
	case "center":
		var left int = pad / 2
		var right int = pad - left
		s += repeat(" ", left) + text + repeat(" ", right)
	default:
		s += text + repeat(" ", pad)
	}
	s += "|"

	fmt.Println(cen(s))
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

func tableContentRows(table Table, rowAmount int, override_pagination bool, colLengths []int) {
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

	// logika paging
	var offset int = ifi(override_pagination, 0, (Temp.table_page-1)*App.pagination)
	var end int = ifi(override_pagination, rowAmount, min(offset+App.pagination, rowAmount))
	// LOGIC UTAMA print baris per baris
	for i = offset; i < end; i++ {
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

func columnLengths(table Table) []int {
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
			currLength = len(truncate(toString(table.rows[row][col]), table.maxLengths[col])) // panjang string di cell ini
			if currLength > columnsLength[col] {
				columnsLength[col] = currLength
			}
		}
		// exit iterasi row di column ini (max udah ketemu)

		// space tambahan, biar ga nempel ke border
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
