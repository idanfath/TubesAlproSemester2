package main

import "fmt"

// ====================== TASK

type Task struct {
	id        int
	nama      string
	prioritas int // 1=Tinggi, 2=Sedang, 3=Rendah
	durasi    int
	date      Date
	selesai   bool
}
type Tasks [NMAX]Task

// mengarahkan user dalam bentuk menu dan navigasi halaman pencatatan tugas
func TaskPage() {
	var tasks = Tasks{
		{id: 1, nama: "belajar alpro", prioritas: 1, durasi: 90, date: Date{year: 2026, month: 6, day: 15}, selesai: false},
		{id: 2, nama: "ngerjain tubes", prioritas: 1, durasi: 120, date: Date{year: 2026, month: 6, day: 15}, selesai: false},
		{id: 3, nama: "olahraga pagi", prioritas: 2, durasi: 30, date: Date{year: 2026, month: 6, day: 16}, selesai: true},
		{id: 4, nama: "review materi kalkulus", prioritas: 1, durasi: 60, date: Date{year: 2026, month: 6, day: 16}, selesai: false},
		{id: 5, nama: "beli bahan makanan", prioritas: 3, durasi: 45, date: Date{year: 2026, month: 6, day: 16}, selesai: true},
		{id: 6, nama: "meeting koordinasi panitia", prioritas: 2, durasi: 120, date: Date{year: 2026, month: 6, day: 17}, selesai: true},
		{id: 7, nama: "fix bug extension", prioritas: 1, durasi: 180, date: Date{year: 2026, month: 6, day: 17}, selesai: false},
		{id: 8, nama: "push rank bareng teman", prioritas: 3, durasi: 150, date: Date{year: 2026, month: 6, day: 17}, selesai: true},
		{id: 9, nama: "streaming turnamen", prioritas: 3, durasi: 120, date: Date{year: 2026, month: 6, day: 18}, selesai: false},
		{id: 10, nama: "ngerjain modul praktikum", prioritas: 1, durasi: 90, date: Date{year: 2026, month: 6, day: 18}, selesai: false},
		{id: 11, nama: "update checklist notion", prioritas: 3, durasi: 15, date: Date{year: 2026, month: 6, day: 19}, selesai: true},
		{id: 12, nama: "slicing ui figma ke code", prioritas: 2, durasi: 180, date: Date{year: 2026, month: 6, day: 19}, selesai: false},
		{id: 13, nama: "deep cleaning kamar", prioritas: 2, durasi: 60, date: Date{year: 2026, month: 6, day: 20}, selesai: true},
		{id: 14, nama: "nonton film", prioritas: 3, durasi: 130, date: Date{year: 2026, month: 6, day: 20}, selesai: true},
		{id: 15, nama: "belajar git advanced", prioritas: 2, durasi: 90, date: Date{year: 2026, month: 6, day: 20}, selesai: false},
		{id: 16, nama: "persiapan materi tech talk", prioritas: 1, durasi: 120, date: Date{year: 2026, month: 6, day: 21}, selesai: false},
		{id: 17, nama: "baca buku bab 3-4", prioritas: 2, durasi: 45, date: Date{year: 2026, month: 6, day: 21}, selesai: true},
		{id: 18, nama: "laundry baju", prioritas: 3, durasi: 30, date: Date{year: 2026, month: 6, day: 22}, selesai: true},
		{id: 19, nama: "latihan soal ctf", prioritas: 1, durasi: 150, date: Date{year: 2026, month: 6, day: 22}, selesai: false},
		{id: 20, nama: "backup data ke cloud", prioritas: 2, durasi: 40, date: Date{year: 2026, month: 6, day: 22}, selesai: true},
		{id: 21, nama: "drafting proposal event", prioritas: 1, durasi: 180, date: Date{year: 2026, month: 6, day: 23}, selesai: false},
		{id: 22, nama: "eksperimen library baru", prioritas: 2, durasi: 90, date: Date{year: 2026, month: 6, day: 23}, selesai: true},
		{id: 23, nama: "cukur", prioritas: 3, durasi: 45, date: Date{year: 2026, month: 6, day: 24}, selesai: true},
		{id: 24, nama: "evaluasi mingguan tubes", prioritas: 1, durasi: 60, date: Date{year: 2026, month: 6, day: 24}, selesai: false},
		{id: 25, nama: "nonton serial experiment lain", prioritas: 3, durasi: 240, date: Date{year: 2026, month: 6, day: 25}, selesai: false},
	}
	var taskCount int = 25
	var taskSort = SortStatus{sorted: true, sortedby: "ID", asc: true}
	title("Halaman Pencatatan Tugas")
	var quitpage bool = false
	for !quitpage {
		var option int = getOptions(Opsi{"List", "Lihat", "Tandai Selesai", "Cari", "Tambah", "Edit", "Hapus", "Sorting", "Statistik", "Beranda"}, 10)
		switch option {
		case 1: // list
			optionListTask(tasks, taskCount, taskSort)
		case 2: // lihat
			optionShowTasks(tasks, taskCount, taskSort)
		case 3: // tandai selesai
			optionMarkAsDone(&tasks, taskCount, &taskSort)
		case 4: // search
			optionSearchTask(tasks, taskCount, taskSort)
		case 5: // add
			optionAddTask(&tasks, &taskCount, &taskSort)
		case 6: // edit
			optionEditTask(&tasks, taskCount, &taskSort)
		case 7: // hapus
			optionDeleteTask(&tasks, &taskCount, &taskSort)
		case 8: // sort
			optionSortTask(&tasks, taskCount, &taskSort)
		case 9: // statistik
			optionShowStatistikTask(tasks, taskCount)
		case 10:
			quitpage = true
		}
	}
}

// memperlihatkan rasio persentase penyelesaian tugas harian
func optionShowStatistikTask(t Tasks, n int) {
	if n == 0 {
		show("Tidak ada data")
		return
	}

	var maxDay int
	for maxDay == 0 {
		showinline("Masukkan jumlah hari untuk statistik (Minimal 1): ")
		fmt.Scanln(&maxDay)
	}

	sortTask(&t, n, 5, 2)
	var currentDate Date = t[0].date
	var i int = 0
	var day int = 1
	var valid bool = true

	show("")
	for day <= maxDay && valid {
		var todaycount, todaydone int
		for i < n && t[i].date == currentDate {
			todaycount++
			if t[i].selesai {
				todaydone++
			}
			i++
		}
		// hitung
		var todaypercent float64 = 0
		if todaycount > 0 {
			todaypercent = float64(todaydone) / float64(todaycount) * 100
		}
		show(fmt.Sprintf("Hari ke-%d (%s) [%s]: %d task, %d selesai (%.2f%%)",
			day,
			ifstring(day == 1, "Hari ini", fmt.Sprintf("%d Hari lalu", day-1)),
			datestring(currentDate, true),
			todaycount, todaydone, todaypercent))
		if i >= n {
			valid = false
		}
		currentDate = getYesterday(currentDate)
		day++
	}
}

// mencari tugas berdasarkan kata kunci nama atau tanggal
func optionSearchTask(t Tasks, n int, ss SortStatus) {
	var i, resultCount int
	var targetText string
	var targetDate Date
	var result Tasks

	var searchByOption = Opsi{"Kata Kunci (Nama)", "Tanggal"}
	var searchBy int = getOptions(searchByOption, 2)

	switch searchBy {
	case 1:
		for len(targetText) == 0 {
			showinline("Masukkan keyword pencarian: ")
			fmt.Scanln(&targetText)
			targetText = replace(targetText, "_", " ")
		}
	case 2:
		for !isValidDate(targetDate) {
			showinline("Masukkan tanggal (contoh: 2045 12 17): ")
			fmt.Scanf("%d %d %d", &targetDate.year, &targetDate.month, &targetDate.day)
		}
	}

	if ss.sorted && searchBy == 2 && ss.sortedby == "Tanggal" {
		var left, right, mid, i int
		var startingidx int = -1
		left = 0
		right = n - 1
		// binary search
		for left <= right && startingidx == -1 {
			mid = left + (right-left)/2
			if t[mid].date == targetDate {
				startingidx = mid
			} else if isDateLarger(targetDate, t[mid].date) {
				if ss.asc {
					left = mid + 1
				} else {
					right = mid - 1
				}
			} else if isDateLarger(t[mid].date, targetDate) {
				if ss.asc {
					right = mid - 1
				} else {
					left = mid + 1
				}
			}
		}
		if startingidx == -1 {
			show("")
			show("Tidak ditemukan hasil apapun")
			return
		}
		// ketemu starting idx, cari upper dan lowerbound
		left = startingidx
		right = startingidx
		// lower bound
		for left > 0 && t[left-1].date == targetDate {
			left--
		}
		// upper bound
		for right < n-1 && t[right+1].date == targetDate {
			right++
		}
		// masukin ke result
		for i = left; i <= right; i++ {
			result[resultCount] = t[i]
			resultCount++
		}
		show("")
		show(fmt.Sprintf("Ketemu %d hasil menggunakan Binary Search (%s)", resultCount, ifstring(ss.asc, "Ascending", "Descending")))
	} else {
		// seq search
		for i = 0; i < n; i++ {
			if (searchBy == 1 && contains(lower(t[i].nama), lower(targetText))) || (searchBy == 2 && t[i].date == targetDate) {
				result[resultCount] = t[i]
				resultCount++
			}
		}
		if resultCount < 1 {
			show("")
			show("Tidak ditemukan hasil apapun")
			return
		}
		show("")
		show(fmt.Sprintf("Ketemu %d hasil menggunakan Sequential Search", resultCount))
	}
	//kalo arr yg dicari sorted, hasilnya udh pasti sorted juga (kyknya)
	optionListTask(result, resultCount, ss)
}

// sorting tugas berdasarkan field dan urutan dari pilihan user
func optionSortTask(t *Tasks, n int, ss *SortStatus) {
	var urutan int = getOptions(Opsi{"Ascending (A-Z)", "Descending (Z-A)"}, 2)
	// kalo ubah orderby, ubah bandinginTask juga
	var pilihanOrder = Opsi{"ID", "Nama", "Prioritas", "Durasi", "Tanggal", "Status"}
	var orderby int = getOptions(pilihanOrder, 6)
	// selection sort
	sortTask(t, n, orderby, urutan)
	ss.sorted = true
	ss.sortedby = pilihanOrder[orderby-1]
	ss.asc = urutan == 1
	show("")
	show(fmt.Sprintf("Berhasil diurutkan berdasarkan %s secara %s", ss.sortedby, ifstring(ss.asc, "Ascending", "Descending")))
}

// menginput atau menambahkan tugas baru ke array
func optionAddTask(t *Tasks, n *int, ss *SortStatus) {
	var task Task
	var highestIdTask = t[getHighestIdTaskIdx(*t, *n, *ss)]
	task.id = highestIdTask.id + 1
	for task.prioritas < 1 || task.prioritas > 3 {
		showinline("Masukkan prioritas task (1-3): ")
		fmt.Scanln(&task.prioritas)
	}
	for len(task.nama) == 0 {
		showinline("Masukkan nama task: ")
		fmt.Scanln(&task.nama)
	}
	for task.durasi <= 0 {
		showinline("Masukkan durasi task (menit): ")
		fmt.Scanln(&task.durasi)
	}
	for !isValidDate(task.date) {
		showinline("Masukkan tanggal pencatatan task (contoh: 2024 12 17): ")
		fmt.Scanf("%d %d %d", &task.date.year, &task.date.month, &task.date.day)
	}
	t[*n] = task
	*n++
	ss.sorted = false
	show("")
	show(fmt.Sprintf("Task [%d] %s berhasil ditambahkan", task.id, task.nama))
}

// menyunting atribut dari data tugas yang telah tersimpan di array
func optionEditTask(t *Tasks, n int, ss *SortStatus) {
	var task, oldTask Task
	var idx int = getTaskIdxFromId(*t, n, *ss)

	if idx == -1 {
		show("ID Tidak Ditemukan")
		return
	}

	oldTask = t[idx]
	task.id = oldTask.id

	show("Kosongkan untuk menggunakan data lama")
	showinline(fmt.Sprintf("Nama tugas (contoh: belajar_alpro): (%s) ", oldTask.nama))
	fmt.Scanln(&task.nama)
	if len(task.nama) == 0 {
		task.nama = oldTask.nama
	}
	task.nama = replace(task.nama, "_", " ")

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

	for !isValidDate(task.date) {
		showinline(fmt.Sprintf("Masukkan tanggal pencatatan task (contoh: 2029 12 17): (%s) ", datestring(oldTask.date, false)))
		fmt.Scanf("%d", &task.date.year)
		if task.date.year == 0 {
			task.date = oldTask.date
		} else {
			fmt.Scanf("%d %d", &task.date.month, &task.date.day)
		}
	}

	var x string
	var selesaiValid bool
	for !selesaiValid {
		showinline(fmt.Sprintf("Sudah selesai? (y/n): (%s) ", ifstring(oldTask.selesai, "y", "n")))
		fmt.Scanln(&x)
		if len(x) == 0 {
			task.selesai = oldTask.selesai
			selesaiValid = true
		} else if lower(x) == "y" {
			task.selesai = true
			selesaiValid = true
		} else if lower(x) == "n" {
			task.selesai = false
			selesaiValid = true
		}
	}

	t[idx] = task
	ss.sorted = false
	show("")
	show(fmt.Sprintf("Task [%d] %s berhasil diubah", task.id, task.nama))
}

// menghapus data tugas
func optionDeleteTask(t *Tasks, n *int, ss *SortStatus) {
	var idx int = getTaskIdxFromId(*t, *n, *ss)
	if idx == -1 {
		show("ID Tidak Ditemukan")
		return
	}
	var deletedId int = t[idx].id
	var i int
	for i = idx; i < *n-1; i++ {
		t[i] = t[i+1]
	}
	t[*n-1] = Task{}
	*n = *n - 1
	ss.sorted = false
	show("")
	show(fmt.Sprintf("Task [%d] berhasil dihapus", deletedId))
}

// memberikan label status pengerjaan tugas menjadi "Selesai"
func optionMarkAsDone(t *Tasks, n int, ss *SortStatus) {
	var idx int = getTaskIdxFromId(*t, n, *ss)
	if idx == -1 {
		show("ID Tidak Ditemukan")
		return
	}
	if t[idx].selesai {
		show(fmt.Sprintf("\nTask [%d] %s sudah ditandai selesai", t[idx].id, t[idx].nama))
	} else {
		t[idx].selesai = true
		show(fmt.Sprintf("\nTask [%d] %s berhasil ditandai selesai", t[idx].id, t[idx].nama))
		ss.sorted = false
	}
}

// menampilkan detail satu tugas berdasarkan ID
func optionShowTasks(t Tasks, n int, ss SortStatus) {
	var idx int = getTaskIdxFromId(t, n, ss)
	if idx == -1 {
		show("ID Tidak Ditemukan")
		return
	}
	show("")
	show(fmt.Sprintf("[%d] Prioritas: %s | %d mnt | %s | %s", t[idx].id, labelPrioritas(t[idx].prioritas), t[idx].durasi, datestring(t[idx].date, true), labelSelesai(t[idx].selesai)))
	show(t[idx].nama)
}

// menampilkan seluruh daftar tugas di layar
func optionListTask(t Tasks, n int, ss SortStatus) {
	var i int
	show("")
	for i = 0; i < n; i++ {
		fmt.Printf("[%d] %s | Prioritas: %s | %d mnt | %s | %s\n", t[i].id,
			truncate(t[i].nama, 50),
			labelPrioritas(t[i].prioritas),
			t[i].durasi,
			datestring(t[i].date, true),
			labelSelesai(t[i].selesai),
		)
	}
	if ss.sorted && n > 0 {
		show(fmt.Sprintf("Diurutkan berdasarkan %s secara %s", ss.sortedby, ifstring(ss.asc, "Ascending", "Descending")))
	}
}

// ====================== TASK (UTILITY)

// "ID", "Nama", "Prioritas", "Durasi", "Status"
// Membandingkan properti dua tugas untuk sorting
func bandinginTask(t1 Task, t2 Task, orderby int, biggerthan bool) bool {
	var result bool
	switch orderby {
	case 1:
		result = t1.id > t2.id
	case 2:
		result = t1.nama > t2.nama
	case 3:
		result = t1.prioritas > t2.prioritas
	case 4:
		result = t1.durasi > t2.durasi
	case 5:
		result = isDateLarger(t1.date, t2.date)
	case 6:
		result = t1.selesai && !t2.selesai
	default:
		result = false
	}
	if biggerthan {
		return result
	}
	return !result
}

// mengurutkan array data tugas dengan Selection Sort
func sortTask(t *Tasks, n int, orderby int, urutan int) {
	var i, j, extreme int
	var temp Task
	for i = 0; i < n; i++ {
		extreme = i
		for j = i + 1; j < n; j++ {
			// kalo descending, kita moods[j] hrs lbh gede dri moods[extreme]
			if bandinginTask(t[j], t[extreme], orderby, urutan == 2) {
				extreme = j
			}
		}
		temp = t[i]
		t[i] = t[extreme]
		t[extreme] = temp
	}
}

// mencari indeks tugas dengan ID terbesar saat ini
func getHighestIdTaskIdx(t Tasks, n int, ss SortStatus) int {
	if ss.sorted && ss.sortedby == "ID" {
		if ss.asc {
			return n - 1
		} else {
			return 0
		}
	}
	var max_idx = 0
	var i int
	for i = 0; i < n; i++ {
		if t[i].id > t[max_idx].id {
			max_idx = i
		}
	}
	return max_idx
}

// meminta input ID tugas dari user lalu mencari indeksnya
func getTaskIdxFromId(t Tasks, n int, ss SortStatus) int {
	showinline("Masukkan ID Task: ")
	var x int
	fmt.Scanln(&x)
	return findTask(t, n, ss, x)
}

// Mencari indeks tugas berdasarkan ID dengan binary/sequential
func findTask(t Tasks, n int, ss SortStatus, id int) int {
	var i int
	if ss.sorted && ss.sortedby == "ID" {
		// binary search kalo mendukung
		var left, right, mid int
		left = 0
		right = n - 1
		for left <= right {
			mid = left + (right-left)/2
			if t[mid].id == id {
				return mid
			} else if t[mid].id < id {
				if ss.asc {
					left = mid + 1
				} else {
					right = mid - 1
				}
			} else {
				if ss.asc {
					right = mid - 1
				} else {
					left = mid + 1
				}
			}
		}
	} else {
		for i = 0; i < n; i++ {
			if t[i].id == id {
				return i
			}
		}
	}

	return -1
}

// mengconvert prioritas tugas (1/2/3) menjadi string
func labelPrioritas(p int) string {
	switch p {
	case 1:
		return "Tinggi"
	case 2:
		return "Sedang"
	case 3:
		return "Rendah"
	default:
		return "Tidak valid"
	}
}

// mengconvert status boolean pengerjaan menjadi string
func labelSelesai(s bool) string {
	if s {
		return "Selesai"
	} else {
		return "Belum Selesai"
	}
}
