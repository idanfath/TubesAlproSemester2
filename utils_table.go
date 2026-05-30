package main

type Table struct {
	header     []string
	rows       [][]string
	maxLengths []int
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
