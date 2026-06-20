package main

import "fmt"

// fungsi untuk menampilkan string di tengah layar (centered)
func show(input string) {
	fmt.Println(cen(input))
}

// fungsi untuk menampilkan string di tengah layar tanpa newline
func showinline(input string) {
	fmt.Print(cen(input))
}

// fungsi untuk menampilkan dan mengembalikan index+1 option yang dipilih user
func getOptions(options [NMAX]string, n int) int {
	var i int
	// print options
	show("")
	for i = 0; i < n; i++ {
		// contoh: [1] Keluar
		show(fmt.Sprintf("[%d] %s", i+1, options[i]))
	}
	// wait for input
	var x int = 0
	var valid bool = false
	for !valid {
		showinline("Pilih opsi: ")
		fmt.Scanln(&x)
		if x >= 1 && x <= n {
			valid = true
		} else {
			show("Opsi tidak valid, silakan pilih lagi")
		}
	}
	return x
}

// reusable untuk menampilkan judul halaman
func title(name string) {
	show("")
	show("-- " + name + " --")
}
