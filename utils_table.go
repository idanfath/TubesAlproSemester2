package main

// return indexnya, buat akses ke tabel aslinya
func getTi(name string) int {
	for i := 0; i < len(tables); i++ {
		if tables[i].name == name {
			return i
		}
	}
	return -1
}

// return tablenya, buat display
func getT(name string) Table {
	var index = getTi(name)
	if index == -1 {
		return Table{}
	}
	return tables[index]
}

func insertT(name string, row []any) {
	var index = getTi(name)
	if index != -1 {
		tables[index].rows = append(tables[index].rows, row)
	}
}

func removeT(name string, rowIndex int) {
	var index = getTi(name)
	var rowCount = len(tables[index].rows)
	if index == -1 || rowIndex < 0 || rowIndex >= rowCount {
		return
	}
	var rowUntilIndex = tables[index].rows[0:rowIndex]            // row 0 sampai sebelum row yang dihapus
	var rowAfterIndex = tables[index].rows[rowIndex+1 : rowCount] // row setelah row yang dihapus sampai akhir
	tables[index].rows = append(rowUntilIndex, rowAfterIndex...)  // gabungin row-row yang udah dipisah tadi
}

func updateRowT(name string, rowIndex int, newRow []any) {
	var index = getTi(name)
	var rowCount = len(tables[index].rows)
	if index == -1 || rowIndex < 0 || rowIndex >= rowCount {
		return
	}
	tables[index].rows[rowIndex] = newRow
}

func updateCellT(name string, rowIndex int, colIndex int, newValue any) {
	var index = getTi(name)
	var rowCount = len(tables[index].rows)
	var colCount = len(tables[index].header)
	if index == -1 || rowIndex < 0 || rowIndex >= rowCount || colIndex < 0 || colIndex >= colCount {
		return
	}
	tables[index].rows[rowIndex][colIndex] = newValue
}

func findInTable(name string, searchValue any) TableIndex {
	var index = getTi(name)
	if index == -1 {
		return TableIndex{-1, -1}
	}
	var row, col int
	var rowCount = len(tables[index].rows)
	var colCount = len(tables[index].header)
	searchValue = lower(toString(searchValue))
	// iterasi ke semua sel, kalo ketemu yang dicari, return indexnya
	for row = 0; row < rowCount; row++ {
		for col = 0; col < colCount; col++ {
			if lower(toString(tables[index].rows[row][col])) == searchValue {
				return TableIndex{row, col}
			}
		}
	}
	return TableIndex{-1, -1}
}

func colValuesT(name string, colIndex int) []any {
	var index = getTi(name)
	var rowCount = len(tables[index].rows)
	if index == -1 || colIndex < 0 || colIndex >= len(tables[index].header) {
		return []any{}
	}
	var values []any
	for i := 0; i < rowCount; i++ {
		values = append(values, tables[index].rows[i][colIndex])
	}
	return values
}

// pake make.. semoga aman
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
			currLength = len(toString(table.rows[row][col])) // panjang string di cell ini
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
