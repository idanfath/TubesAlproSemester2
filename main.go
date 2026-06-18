package main

// ====================== MAIN

type Opsi [NMAX]string
type SortStatus struct {
	sorted   bool
	asc      bool
	sortedby string
}

// fungsi utama program, menampilkan menu utama
func main() {
	var stop bool
	for !stop {
		title("Selamat datang di Mindflow")
		var operation int = getOptions(Opsi{"Pencatatan Mood", "Pencatatan Tugas", "Keluar"}, 3)
		switch operation {
		case 1:
			MoodPage()
		case 2:
			TaskPage()
		case 3:
			show("Goodbye world")
			stop = true
		}
	}
}
