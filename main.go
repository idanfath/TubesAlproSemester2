package main

func main() {
	// inisialisasi aplikasi, termasuk halaman, konfigurasi, dll
	initializeApp()
	// do while, karna mau render dulu sekali sebelum stop loop
	for {
		render()
		if KILLSIG {
			render()
			break
		}
	}
}
