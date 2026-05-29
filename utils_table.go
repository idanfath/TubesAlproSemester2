package main

import "fmt"

// return indexnya, buat akses langsung ke tabel aslinya, untuk operasi kyk write, update, delete, dll
func getTi(name string) int {
	for i := 0; i < len(tables); i++ {
		if tables[i].name == name {
			return i
		}
	}
	return -1
}

// return tablenya, buat display atau sekedar akses data
func getT(name string) Table {
	var index = getTi(name)
	if index == -1 {
		return Table{}
	}
	return tables[index]
}

func insertT(name string, row []any) {
	var index = getTi(name)
	var table = getT(name)
	// cek kalo ada unique_col_index
	if len(table.unique_col_index) > 0 {
		var i, j int
		// untuk tiap tiap index unique
		for i = 0; i < len(table.unique_col_index); i++ {
			// ambil semua value di column yang dicek, dan value di row baru yang mau dicek
			var colValues = colValuesT(name, table.unique_col_index[i]) // contoh (id): [1, 2, 3, 4]
			var newRowColValue = row[table.unique_col_index[i]]         // contoh: 3
			// cek value di row baru sama dengan value di column yang dicek, kalo ada yang sama, jangan insert
			for j = 0; j < len(colValues); j++ {
				if colValues[j] == newRowColValue {
					TempRenderQ = []RenderData{{text: fmt.Sprintf("Gagal menambahkan data! Kolom %s harus unik, tapi ada data lain yang punya nilai %v di kolom tersebut.", table.header[table.unique_col_index[i]], newRowColValue)}}
					return
				}
			}
		}
	}

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
	var i int
	var rowCount = len(tables[index].rows)
	if index == -1 || rowIndex < 0 || rowIndex >= rowCount {
		return
	}
	// cek value tiap column di row baru, kalu kosong, pakai value lamanya
	for i = 0; i < len(newRow); i++ {
		if newRow[i] == "" {
			newRow[i] = tables[index].rows[rowIndex][i]
		}
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

// return index row yang columnya (colIndex) punya value yang dicari (searchValue), kalo ga ketemu return -1
func findFirstInTableCol(name string, searchValue any, colIndex int) int {
	var table = getT(name)
	if table.name == "" {
		return -1
	}
	var row int
	var rowCount = len(table.rows)
	searchValue = lower(toString(searchValue))
	for row = 0; row < rowCount; row++ {
		if lower(toString(table.rows[row][colIndex])) == searchValue {
			return row
		}
	}
	return -1
}

func colValuesT(name string, colIndex int) []any {
	var table = getT(name)
	if table.name == "" {
		return []any{}
	}
	var rowCount = len(table.rows)
	if colIndex < 0 || colIndex >= len(table.header) {
		return []any{}
	}
	var values []any
	for i := 0; i < rowCount; i++ {
		values = append(values, table.rows[i][colIndex])
	}
	return values
}
