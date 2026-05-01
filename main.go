package main

import (
	"fmt"
)

// types and config
type AppConfig struct {
	Width int
}
type Option struct {
	Name string
}
type Options []Option

var App = AppConfig{
	Width: 150,
}

// center a string
func centeredString(s string, w int) string {
	return repeat(" ", (w-len(s))/2) + s
}
func center(s string) string {
	return centeredString(s, App.Width)
}

// strings util
func repeat(s string, count int) string {
	var result string
	var i int
	for i = 0; i < count; i++ {
		result += s
	}
	return result
}
func toString(s any) string {
	return fmt.Sprint(s)
}

// Tables stuffs
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

// options
func printOptions(options Options) {
	var longestLen = getOptionsLongestLength(options)
	var optionsAmount int = len(options)
	var i int
	var s string
	for i = 0; i < optionsAmount; i++ {
		s = fmt.Sprintf("[%d] %-*s", i, longestLen, options[i].Name)
		fmt.Println(center(s))
	}
}
func getOptionsLongestLength(options Options) int {
	var max int = 0
	var i int
	var optionsAmount int = len(options)
	for i = 0; i < optionsAmount; i++ {
		if len(options[i].Name) > max {
			max = len(options[i].Name)
		}
	}
	return max
}

func main() {
	fmt.Println(center("Hello, World!"))
	fmt.Println(center("My CLI App"))
	var test = Table{
		Header: []string{"Name", "Age", "City"},
		Rows: [][]any{
			{"Alice", 30, "New York"},
			{"Bob", "agus", "Los Angeles"},
			{"Charlie", 35, "Chicago"},
		},
	}
	printTable(test)
	var testOption = Options{
		{
			Name: "Bayar",
		},
		{
			Name: "Gaaaaaaaaaaaaa",
		},
		{
			Name: "Exit",
		},
	}
	printOptions(testOption)
}
