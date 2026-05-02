package main

import "fmt"

func main() {
	initializeApp()
	fmt.Println(cen(App.title))
	fmt.Println("initialized table:")
	printT(getT("Food"))
	fmt.Println("\nafter inserting new row:")
	insertT("Food", []any{4, "David", "Pasta"})
	printT(getT("Food"))
	fmt.Println("\nafter removing row:")
	removeT("Food", 1)
	printT(getT("Food"))
	fmt.Println("\nafter updating row:")
	updateRowT("Food", 1, []any{5, "Charlie", "Salad"})
	printT(getT("Food"))
	fmt.Println("\nafter updating cell:")
	updateCellT("Food", 0, 2, "Spaghetti")
	printT(getT("Food"))
	fmt.Println("\nafter finding value:")
	var searchResult = findInTable("Food", "Charlie")
	fmt.Printf("Found Charlie at rowIndex %d, columnIndex %d\n", searchResult.rowIndex, searchResult.colIndex)
	fmt.Println("\nafter finding value that doesn't exist:")
	var searchResult2 = findInTable("Food", "Eve")
	fmt.Printf("Found Eve at rowIndex %d, columnIndex %d\n", searchResult2.rowIndex, searchResult2.colIndex)
	fmt.Println("\nColumn Values ID (index 0):")
	fmt.Println(colValuesT("Food", 0))
}
