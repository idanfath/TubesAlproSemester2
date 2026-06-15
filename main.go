package main

func main() {
	for !SIGKILL {
		title("Selamat datang di " + App.title)
		var operation int = getOptions([]string{"Pencatatan Mood", "Pencatatan Tugas", "Keluar"})
		switch operation {
		case 1:
			pencatatanMood(true)
		case 2:
			pencatatanTask(true)
		case 3:
			show("Goodbye world")
			exit()
		}
	}
}
