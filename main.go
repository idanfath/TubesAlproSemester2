package mindflow

import (
	"fmt"
)

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
