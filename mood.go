package main

import "fmt"

// ====================== MOOD

type Mood struct {
	id          int
	description string
	score       int
	date        Date
}

const NMAX = 9999

type Moods [NMAX]Mood

// mengarahkan menu dan navigasi halaman pencatatan mood
func MoodPage() {
	// ID, Deskripsi, Skor, Tanggal
	var moods = Moods{
		{id: 1, description: "gila gw happy banget!", score: 8, date: Date{year: 2026, month: 6, day: 1}},
		{id: 2, description: "ga belajar tapi tetep lulus coy", score: 10, date: Date{year: 2026, month: 6, day: 2}},
		{id: 3, description: "coba ga belajar lagi, mau tes hoki, dpt 20", score: 3, date: Date{year: 2026, month: 6, day: 3}},
		{id: 4, description: "Biasa aja sih", score: 5, date: Date{year: 2026, month: 6, day: 4}},
		{id: 5, description: "lagi apes hari ini", score: 3, date: Date{year: 2026, month: 6, day: 5}},
		{id: 6, description: "Project tubes akhirnya kelar", score: 8, date: Date{year: 2026, month: 6, day: 6}},
		{id: 7, description: "dapet SKIN vandallll", score: 9, date: Date{year: 2026, month: 6, day: 7}},
		{id: 8, description: "Badan agak kurang enak", score: 4, date: Date{year: 2026, month: 6, day: 8}},
		{id: 9, description: "cukup happy", score: 7, date: Date{year: 2026, month: 6, day: 9}},
		{id: 10, description: "Gila menang HACKATHON", score: 10, date: Date{year: 2026, month: 6, day: 10}},
		{id: 11, description: "lagi kurang enak badan", score: 2, date: Date{year: 2026, month: 6, day: 11}},
		{id: 12, description: "happy sih cuma quiz kalkulus ga sesuai ekspetasi", score: 6, date: Date{year: 2026, month: 6, day: 12}},
		{id: 13, description: "jalanan telkom macettt", score: 3, date: Date{year: 2026, month: 6, day: 13}},
		{id: 14, description: "gila paper rax go to ewc", score: 9, date: Date{year: 2026, month: 6, day: 14}},
		{id: 15, description: "damn qots basis data se ez itu", score: 8, date: Date{year: 2026, month: 6, day: 15}},
	}
	var moodCount int = 15
	var moodSort = SortStatus{sorted: true, asc: true, sortedby: "ID"} //wajib set sorted ke false tiap modifikasi (crud) moodnya
	title("Halaman Pencatatan Mood")
	var quitpage bool = false
	for !quitpage {
		var option int = getOptions(Opsi{"List", "Lihat", "Cari", "Tambah", "Edit", "Hapus", "Sorting", "Statistik", "Beranda"}, 9)
		switch option {
		case 1: // list
			optionListMoods(moods, moodCount, moodSort)
		case 2: // lihat
			optionShowMoods(moods, moodCount, moodSort)
		case 3: // search
			optionSearchMoods(moods, moodCount, moodSort)
		case 4: // add
			optionAddMood(&moods, &moodCount, &moodSort)
		case 5: // edit
			optionEditMood(&moods, moodCount, &moodSort)
		case 6: // hapus
			optionDeleteMood(&moods, &moodCount, &moodSort)
		case 7: // sort
			optionSortMood(&moods, moodCount, &moodSort)
		case 8: // statistikkkkkkkkkk
			optionShowStatistikMood(moods, moodCount)
		case 9:
			quitpage = true
		}
	}
}

// menampilkan seluruh daftar mood
func optionListMoods(t Moods, n int, ss SortStatus) {
	var i int
	show("")
	for i = 0; i < n; i++ {
		fmt.Printf("[%d] Mood %d/10 | %s | %s\n", t[i].id, t[i].score, datestring(t[i].date, true), t[i].description)
	}
	if ss.sorted && n > 0 {
		show(fmt.Sprintf("Diurutkan berdasarkan %s secara %s", ss.sortedby, ifstring(ss.asc, "Ascending", "Descending")))
	}
}

// mencari mood berdasarkan keyword atau tanggal
func optionSearchMoods(t Moods, n int, ss SortStatus) {
	var i, resultCount int
	var targetDesc string
	var targetDate Date
	var result Moods

	var searchByOption = Opsi{"Kata Kunci (Deskripsi)", "Tanggal"}
	var searchBy int = getOptions(searchByOption, 2)

	switch searchBy {
	case 1:
		for len(targetDesc) == 0 {
			showinline("Masukkan keyword pencarian: ")
			fmt.Scanln(&targetDesc)
			targetDesc = replace(targetDesc, "_", " ")
		}
	case 2:
		for !isValidDate(targetDate) {
			showinline("Masukkan tanggal (contoh: 2069 12 17): ")
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
		// ketemu starting idx, cari upper dan lowerbound, reuse variable
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
			if (searchBy == 1 && contains(lower(t[i].description), lower(targetDesc))) || (searchBy == 2 && t[i].date == targetDate) {
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
	optionListMoods(result, resultCount, ss)
}

// menampilkan detail mood berdasarkan ID
func optionShowMoods(moods Moods, moodCount int, moodSort SortStatus) {
	var idx int = getMoodIdxFromId(moods, moodCount, moodSort)
	if idx == -1 {
		show("ID Tidak Ditemukan")
		return
	}
	show("")
	show(fmt.Sprintf("[%d] %d/10 | %s", moods[idx].id, moods[idx].score, datestring(moods[idx].date, false)))
	show(moods[idx].description)
	show("")
}

// menginput dan menambahkan mood baru ke array
func optionAddMood(t *Moods, n *int, ss *SortStatus) {
	var mood Mood
	var highestIdMood = t[getHighestIdMoodIdx(*t, *n, *ss)]
	mood.id = highestIdMood.id + 1
	for mood.score < 1 || mood.score > 10 {
		showinline("Masukkan skor mood (1-10): ")
		fmt.Scanln(&mood.score)
	}
	for len(mood.description) == 0 {
		showinline("Masukkan deskripsi mood (contoh: sangat_senang): ")
		fmt.Scanln(&mood.description)
		mood.description = replace(mood.description, "_", " ")
	}
	for !isValidDate(mood.date) {
		showinline("Masukkan tanggal pencatatan mood (contoh: 2024 12 17): ")
		fmt.Scanf("%d %d %d", &mood.date.year, &mood.date.month, &mood.date.day)
	}
	t[*n] = mood
	*n++
	ss.sorted = false
	show("")
	show(fmt.Sprintf("Mood [%d] %d/10 %s %s berhasil ditambahkan", mood.id, mood.score, datestring(mood.date, true), truncate(mood.description, 10)))
}

// mengedit mood berdasarkan ID, user bisa pilih field mana yang mau diedit, dan bisa mengostongkan input kalau mau pakai data lama
func optionEditMood(t *Moods, n int, ss *SortStatus) {
	var m, oldm Mood
	var idx int = getMoodIdxFromId(*t, n, *ss)

	if idx == -1 {
		show("ID Tidak Ditemukan")
		return
	}

	oldm = t[idx]
	m.id = oldm.id

	show("Kosongkan untuk menggunakan data lama")
	for m.score < 1 || m.score > 10 {
		showinline(fmt.Sprintf("Masukkan skor mood (1-10): (%d) ", oldm.score))
		fmt.Scanln(&m.score)
		if m.score == 0 {
			m.score = oldm.score
		}
	}

	showinline(fmt.Sprintf("Masukkan deskripsi mood (contoh: sangat_senang): (%s) ", truncate(oldm.description, 12)))
	fmt.Scanln(&m.description)
	if len(m.description) == 0 {
		m.description = oldm.description
	}
	m.description = replace(m.description, "_", " ")

	for !isValidDate(m.date) {
		showinline(fmt.Sprintf("Masukkan tanggal pencatatan mood (contoh: 2029 12 17): (%s) ", datestring(oldm.date, false)))
		fmt.Scanf("%d", &m.date.year)
		if m.date.year == 0 {
			m.date = oldm.date
		} else {
			fmt.Scanf("%d %d", &m.date.month, &m.date.day)
		}
	}
	t[idx] = m
	ss.sorted = false
	show("")
	show(fmt.Sprintf("Mood [%d] berhasil diedit", m.id))
}

// menghapus mood dengan ID, lalu geser semua mood di kanannya 1 langkah ke kiri, lalu kurangi jumlah mood
func optionDeleteMood(t *Moods, n *int, ss *SortStatus) {
	var idx int = getMoodIdxFromId(*t, *n, *ss)
	if idx == -1 {
		show("ID Tidak Ditemukan")
		return
	}
	var id int = t[idx].id
	var i int
	for i = idx; i < *n-1; i++ {
		t[i] = t[i+1]
	}
	t[*n-1] = Mood{}
	*n = *n - 1
	ss.sorted = false
	show("")
	show(fmt.Sprintf("Mood [%d] berhasil dihapus", id))
}

// sorting mood berdasarkan field dan urutan dari pilihan user
func optionSortMood(t *Moods, n int, ss *SortStatus) {
	var urutan int = getOptions(Opsi{"Ascending (A-Z)", "Descending (Z-A)"}, 2)
	// kalo ubah orderby, ubah bandinginMood juga
	var pilihanOrder = Opsi{"ID", "Deskripsi", "Skor", "Tanggal"}
	var orderby int = getOptions(pilihanOrder, 4)
	// selection sort
	sortMood(t, n, orderby, urutan)
	ss.sorted = true
	ss.sortedby = pilihanOrder[orderby-1]
	ss.asc = urutan == 1
	show("")
	show(fmt.Sprintf("Berhasil diurutkan berdasarkan %s secara %s", ss.sortedby, ifstring(ss.asc, "Ascending", "Descending")))
}

// menampilkan statistik mood mingguan, user bisa pilih berapa minggu terakhir yang mau ditampilkan
func optionShowStatistikMood(t Moods, n int) {
	if n == 0 {
		show("Tidak ada data")
		return
	}
	var maxWeek int
	for maxWeek == 0 {
		showinline("Masukkan jumlah minggu untuk statistik (Minimal 1): ")
		fmt.Scanln(&maxWeek)
	}
	// isi temp, sort dulu biar enak
	sortMood(&t, n, 4, 2)
	var currentDate = t[0].date
	var endingweek Date
	var day, pointer int
	pointer = 0
	var week int
	week = 1
	var valid bool = true
	show("")
	for week <= maxWeek && valid {
		var thisweekscore = 0
		var thisweekdatacount = 0
		var startingweek = currentDate // cuma buat display
		day = 0
		for day < 7 && pointer < n {
			// loop dari 0 sampai datenya beda, tambahin ke score dan datacount
			for pointer < n && currentDate == t[pointer].date {
				thisweekscore += t[pointer].score
				thisweekdatacount++
				pointer++
			}
			// sebelum diubah, kita store dulu buat display aja
			if day == 6 || pointer >= n {
				endingweek = currentDate
			}
			currentDate = getYesterday(currentDate)
			day++
		}

		var avg float64
		if thisweekdatacount > 0 {
			avg = float64(thisweekscore) / float64(thisweekdatacount)
		}
		show(fmt.Sprintf("Minggu %d %s [%s s/d %s] | Rata-rata Mood: %.2f (%d data)",
			week,
			ifstring(week == 1, "(Minggu Ini)", fmt.Sprintf("(%d Minggu Lalu)", week-1)),
			datestring(startingweek, true),
			datestring(endingweek, true),
			avg,
			thisweekdatacount))

		if pointer >= n {
			valid = false
		}
		week++
	}
}

// ====================== MOOD (UTILITY)

// utilitas untuk mencari index mood berdsarkan ID, menggunakan binary search jika sorted by ID, jika tidak, menggunakan sequential search
func findMood(t Moods, n int, ss SortStatus, id int) int {
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

// utilitas untuk mendapatkan index mood dgn ID tertinggi, utk add mood, kalau udah sorted by ID, tinggal ambil first atua last element sesuai urutan sorting, jika tidak, cari dengan sequential search
func getHighestIdMoodIdx(t Moods, n int, ss SortStatus) int {
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

// reusable function, untuk input ID dan mendapatkan index mood
func getMoodIdxFromId(t Moods, n int, ss SortStatus) int {
	showinline("Masukkan ID Mood: ")
	var x int
	fmt.Scanln(&x)
	return findMood(t, n, ss, x)
}

// fungsi utility untuk membandingkan dua mood berdsarkan field dan urutan sorting
func bandinginMood(t1 Mood, t2 Mood, orderby int, biggerthan bool) bool {
	var result bool
	switch orderby {
	case 1:
		result = t1.id > t2.id
	case 2:
		result = t1.description > t2.description
	case 3:
		result = t1.score > t2.score
	case 4:
		result = isDateLarger(t1.date, t2.date)
	}
	if biggerthan {
		return result
	}
	return !result
}

// logika utama untuk sorting mood (selection sort)
func sortMood(t *Moods, n int, orderby int, urutan int) {
	var i, j, extreme int
	var temp Mood
	for i = 0; i < n; i++ {
		extreme = i
		for j = i + 1; j < n; j++ {
			// kalo descending, kita moods[j] hrs lbh gede dri moods[extreme]
			if bandinginMood(t[j], t[extreme], orderby, urutan == 2) {
				extreme = j
			}
		}
		temp = t[i]
		t[i] = t[extreme]
		t[extreme] = temp
	}
}
