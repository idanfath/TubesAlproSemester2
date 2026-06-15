package main

import "fmt"

type Task struct {
	id        int
	nama      string
	prioritas int // 1=Tinggi, 2=Sedang, 3=Rendah
	durasi    int // menit
	selesai   bool
}

const MAX_TASK = 9999

var tasks = [MAX_TASK]Task{
	{id: 1, nama: "Belajar Alpro", prioritas: 1, durasi: 90, selesai: false},
	{id: 2, nama: "Ngerjain tubes", prioritas: 1, durasi: 120, selesai: false},
	{id: 3, nama: "Olahraga pagi", prioritas: 2, durasi: 30, selesai: true},
}

var taskCount int = 3

func labelPrioritas(p int) string {
	switch p {
	case 1:
		return "Tinggi"
	case 2:
		return "Sedang"
	case 3:
		return "Rendah"
	default:
		return "?"
	}
}

func labelSelesai(s bool) string {
	if s {
		return "Selesai"
	} else {
		return "Belum Selesai"
	}
}

func showTasks() {
	var i int
	for i = 0; i < taskCount; i++ {
		fmt.Printf("[%d] %s | Prioritas: %s | %d mnt | %s\n", tasks[i].id,
			truncate(tasks[i].nama, (App.width/3)*2),
			labelPrioritas(tasks[i].prioritas),
			tasks[i].durasi,
			labelSelesai(tasks[i].selesai),
		)
	}
	show("")
}

func pencatatanTask(showItems bool) {
	title("Halaman Pencatatan Task")
	if showItems {
		showTasks()
	}
	var option int = getOptions([]string{"Tandai Selesai", "Lihat", "Tambah", "Update", "Hapus", "Beranda"})
	switch option {
	case 1:
		optionSelesai()
	case 2:
		optionLihatTask()
	case 3:
		optionTambahTask()
	case 4:
		optionUbahTaks()
	case 5:
		optionHapusTask()
	case 6:
		return
	}
}

func optionSelesai() {
	var idx int = -1
	for idx == -1 {
		idx = getTaskIdxFromId()
	}
	tasks[idx].selesai = true
	pencatatanTask(true)
}

func optionUbahTaks() {
	var task, oldTask Task
	var idx int = -1
	for idx == -1 {
		idx = getTaskIdxFromId()
	}
	oldTask = tasks[idx]
	task.id = oldTask.id

	show("Kosongkan untuk menggunakan data lama")
	showinline(fmt.Sprintf("Nama tugas (contoh: belajar_alpro): (%s) ", oldTask.nama))
	fmt.Scanln(&task.nama)
	if len(task.nama) == 0 {
		task.nama = oldTask.nama
	}

	for task.prioritas < 1 || task.prioritas > 3 {
		showinline(fmt.Sprintf("Prioritas (1=Tinggi, 2=Sedang, 3=Rendah): (%d) ", oldTask.prioritas))
		fmt.Scanln(&task.prioritas)
		if task.prioritas == 0 {
			task.prioritas = oldTask.prioritas
		}
	}

	for task.durasi <= 0 {
		showinline(fmt.Sprintf("Durasi (menit): (%d) ", oldTask.durasi))
		fmt.Scanln(&task.durasi)
		if task.durasi == 0 {
			task.durasi = oldTask.durasi
		}
	}

	var x string
	for lower(x) != "y" && lower(x) != "n" {
		showinline(fmt.Sprintf("Sudah selesai? (y/n): (%s) ", ifstring(oldTask.selesai, "y", "n")))
		fmt.Scanln(&x)
		if len(x) == 0 {
			task.selesai = oldTask.selesai
			break //boleh lah ya
		} else if lower(x) == "y" {
			task.selesai = true
		} else if lower(x) == "n" {
			task.selesai = false
		}
	}

	updateTask(idx, task)
	pencatatanTask(true)
}

// func optionUbahTaks() {
// 	var idx int = -1
// 	for idx == -1 {
// 		idx = getTaskIdxFromId()
// 	}

// 	show("")
// 	show(fmt.Sprintf("[%d] %s | Prioritas: %s | %d menit | %s",
// 		tasks[idx].id,
// 		tasks[idx].nama,
// 		labelPrioritas(tasks[idx].prioritas),
// 		tasks[idx].durasi,
// 		labelSelesai(tasks[idx].selesai),
// 	))
// 	show("")

// 	var field int = getOptions([]string{"nama", "prioritas", "durasi", "selesai"})
// 	switch field {
// 	case 1:
// 		showinline("Nama baru: ")
// 		var nama string
// 		fmt.Scanln(&nama)
// 		tasks[idx].nama = nama
// 	case 2:
// 		showinline("Prioritas baru (1=Tinggi, 2=Sedang, 3=Rendah): ")
// 		var prioritas int
// 		fmt.Scanln(&prioritas)
// 		tasks[idx].prioritas = prioritas
// 	case 3:
// 		showinline("Durasi baru (menit): ")
// 		var durasi int
// 		fmt.Scanln(&durasi)
// 		tasks[idx].durasi = durasi
// 	case 4:
// 		showinline("Sudah selesai? (1=Ya, 0=Tidak): ")
// 		var selesaiInt int
// 		fmt.Scanln(&selesaiInt)
// 		tasks[idx].selesai = selesaiInt == 1
// 	}
// 	pencatatanTask(true)
// }

func optionTambahTask() {
	var task Task
	var highestIdTask = tasks[getHighestIdTaskIdx()]
	show(toString(highestIdTask))
	task.id = highestIdTask.id + 1

	for len(task.nama) == 0 {
		showinline("Nama tugas (contoh: belajar_alpro): ")
		fmt.Scanln(&task.nama)
		task.nama = replace(task.nama, "_", " ")
	}

	for task.prioritas < 1 || task.prioritas > 3 {
		showinline("Prioritas (1=Tinggi, 2=Sedang, 3=Rendah): ")
		fmt.Scanln(&task.prioritas)
	}

	for task.durasi <= 0 {
		showinline("Durasi (menit): ")
		fmt.Scanln(&task.durasi)
	}

	task.selesai = false
	insertTask(task)
	pencatatanTask(true)
}

func insertTask(t Task) {
	tasks[taskCount] = t
	taskCount++
}

func optionHapusTask() {
	var idx int = -1
	for idx == -1 {
		idx = getTaskIdxFromId()
	}
	deleteTasks(idx)
	pencatatanTask(true)
}

func optionLihatTask() {
	var idx int = -1
	for idx == -1 {
		idx = getTaskIdxFromId()
	}

	show("")
	show(fmt.Sprintf("[%d] Prioritas %s | %d Menit | %s",
		tasks[idx].id,
		labelPrioritas(tasks[idx].prioritas),
		tasks[idx].durasi,
		labelSelesai(tasks[idx].selesai),
	))
	show(tasks[idx].nama)
	show("")
	pencatatanTask(false)
}

func findTask(id int) int {
	var i int
	for i = 0; i < taskCount; i++ {
		if tasks[i].id == id {
			return i
		}
	}
	return -1
}

func getHighestIdTaskIdx() int {
	var max_idx = 0
	var i int
	for i = 0; i < taskCount; i++ {
		if tasks[i].id > tasks[max_idx].id {
			max_idx = i
		}
	}
	return max_idx
}

func getTaskIdxFromId() int {
	showinline("Masukkan ID Task: ")
	var x int
	fmt.Scanln(&x)
	return findTask(x)
}

func deleteTasks(idx int) {
	var newTask [MAX_TASK]Task
	var j, i int
	j = 0
	for i = 0; i < taskCount; i++ {
		if i == idx {
			continue
		}
		newTask[j] = tasks[i]
		j++
	}
	tasks = newTask
	taskCount = j
}

func updateTask(idx int, task Task) {
	tasks[idx] = task
}
