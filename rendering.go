package main

import "fmt"

func show(input string) {
	fmt.Println(cen(input))
}

func showinline(input string) {
	fmt.Print(cen(input))
}

// loop inputs, ubah ke string, satukan lalu print
// func showMany(inputs []any) {
// 	var i int
// 	var s string
// 	for i = 0; i < len(inputs); i++ {
// 		s = s + " " + toString(inputs[i])
// 	}
// 	show(s)
// }

func getOptions(options []string) int {
	var i int
	// print options
	for i = 0; i < len(options); i++ {
		// contoh: [1] Keluar
		show(fmt.Sprintf("[%d] %s", i+1, options[i]))
	}
	// wait for input
	var x int = 0
	showinline("Pilih opsi: ")
	fmt.Scan(&x)
	return x
}

func title(name string) {
	show("")
	show("-- " + name + " --")
	show("")
}
