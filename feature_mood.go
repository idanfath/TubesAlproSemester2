package main

import "fmt"

type Mood struct {
	id          int
	description string
	score       int
	date        string
}

var moods = []Mood{
	{id: 1, description: "Hari yang menyenangkan!", score: 8, date: "2024-06-01"},
	{id: 2, description: "Agak stres dengan pekerjaan. Sedih banget pokoknya diputusin terus remedial aduhhh sad bgt coy", score: 1, date: "2024-06-02"},
	{id: 3, description: "Merasa sangat bahagia!", score: 9, date: "2024-06-03"},
	{id: 4, description: "Biasa aja, flat banget seharian ga ngapa-ngapain.", score: 5, date: "2024-06-04"},
	{id: 5, description: "Dapet sanksi telat masuk kantor, apes banget.", score: 3, date: "2024-06-05"},
	{id: 6, description: "Project akhirnya kelar dan diapprove client, lega!", score: 8, date: "2024-06-06"},
	{id: 7, description: "Mulai pelihara kucing baru, gemes banget bikin mood naik.", score: 9, date: "2024-06-07"},
	{id: 8, description: "Badan agak meriang dan pusing, cuma rebahan doang.", score: 4, date: "2024-06-08"},
	{id: 9, description: "Minggu produktif, berhasil beres-beres kamar yang berantakan.", score: 7, date: "2024-06-09"},
	{id: 10, description: "Gaji masuk! Waktunya self reward makan enak.", score: 10, date: "2024-06-10"},
	{id: 11, description: "Sakit gigi melanda, ga bisa tidur semalaman.", score: 2, date: "2024-06-11"},
	{id: 12, description: "Kondisi membaik, dapet kopi gratis dari temen kantor.", score: 6, date: "2024-06-12"},
	{id: 13, description: "Terjebak macet total pas hujan deras, capek banget di jalan.", score: 3, date: "2024-06-13"},
	{id: 14, description: "Nonton konser band favorit, capek tapi seru parah!", score: 9, date: "2024-06-14"},
	{id: 15, description: "Santai di rumah sambil marathon series bareng keluarga.", score: 8, date: "2024-06-15"},
}
var searchResultsMood []Mood

func MoodTable() Table {
	return buildMoodTable(moods)
}

func buildMoodTable(moods []Mood) Table {
	var table Table
	table.header = []string{"ID", "Deskripsi", "Skor", "Tanggal"}
	table.maxLengths = []int{5, 40, 5, 12}
	var i int
	for i = 0; i < len(moods); i++ {
		table.rows = append(table.rows, []string{
			toString(moods[i].id),
			moods[i].description,
			toString(moods[i].score),
			moods[i].date,
		})
	}
	return table
}

func insertMood(mood Mood) {
	mood.id = getMaxMoodID() + 1
	moods = append(moods, mood)
}

func deleteMood(id int) {
	var i int
	var newMoods []Mood
	for i = 0; i < len(moods); i++ {
		if moods[i].id == id {
			continue
		}
		newMoods = append(newMoods, moods[i])
	}
	moods = newMoods
}

func updateMood(id int, newMood Mood) {
	var index = findMoodIndex(id)
	if index == -1 {
		return
	}
	moods[index] = newMood
}

// pencarian sequential deskripsi dan tanggal
func seqSearchMood(keyword string) []Mood {
	var result []Mood
	var i int
	for i = 0; i < len(moods); i++ {
		if contains(lower(moods[i].date), lower(keyword)) || contains(lower(moods[i].description), lower(keyword)) {
			result = append(result, moods[i])
		}
	}
	return result
}

func isMoodScoreSorted(x []Mood, asc bool) bool {
	var i int
	var last = x[0].score
	for i = 0; i < len(x); i++ {
		if asc {
			if x[i].score < last {
				return false
			}
		} else {
			if x[i].score > last {
				return false
			}
		}
	}
	return true
}

func insertionSortMood(arr []Mood, comparator func(key, prev Mood) bool) []Mood {
	var i, j int
	var key Mood
	for i = 0; i < len(arr); i++ {
		key = arr[i]
		j = i - 1
		for j >= 0 && comparator(key, arr[j]) {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}

// pencarian binary khusus untuk score mood
// todo: pecah jadi 2 fungsi, satu buat search 1satu buat orchestration
func binarySearchMood(target int) []Mood {
	var arr []Mood //biar ga override urutan data mood, kita buat copy
	var i int
	for i = 0; i < len(moods); i++ {
		arr = append(arr, moods[i])
	}
	if !isMoodScoreSorted(arr, true) {
		arr = insertionSortMood(arr, func(key, prev Mood) bool {
			return key.score < prev.score // asc
		})
	}

	var left, right, mid int
	var found bool
	left = 0
	right = len(arr) - 1

	// cari satu data yang memenuhi
	for left <= right && !found {
		mid = left + (right-left)/2
		if arr[mid].score == target {
			found = true
		} else if arr[mid].score < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if !found {
		return []Mood{}
	}

	// 1 contoh udah ketemu, cari lower and higher bound
	// pake ulg var left right
	left = mid
	for left > 0 && arr[left-1].score == arr[mid].score {
		left = left - 1
	}
	right = mid
	for right < len(arr)-1 && arr[right+1].score == arr[mid].score {
		right = right + 1
	}
	// arr[0:5] tuh include 0 dan exclude 5, jadi harus +1 (4+1 = 5)
	return arr[left : right+1]
}

// --

func appOptionMoodSearch() {
	var results []Mood
	var option int
	var searchOptionsTable = Table{
		header:     []string{"1", "2"},
		rows:       [][]string{{"Deskripsi & Tanggal", "Skor Mood"}},
		maxLengths: []int{30, 30},
	}
	printTable(searchOptionsTable, true)
	inputNumber(InputNumber{
		prompt: "Masukkan tipe pencarian (1-2): ",
		onSubmit: func(input float64) {
			option = int(input)
		},
	})
	switch option {
	case 1:
		var keyword string
		input(Input{
			prompt: "Masukkan pencarian: ",
			onSubmit: func(input string) {
				keyword = input
			},
		})
		results = seqSearchMood(keyword)
	case 2:
		var keyscore int
		inputNumber(InputNumber{
			prompt: "Masukkan pencarian: ",
			onSubmit: func(input float64) {
				keyscore = int(input)
			},
		})
		results = binarySearchMood(keyscore)
	default:
		alert("Opsi pencarian tidak dikenali!")
		return
	}
	if len(results) == 0 {
		alert("Tidak ada catatan mood yang cocok.")
		return
	}
	searchResultsMood = results
	toPage("MoodSearchResult")
}

func appOptionMoodView() {
	var id int
	inputNumber(InputNumber{
		prompt: "Masukkan ID catatan mood yang ingin dilihat: ",
		onSubmit: func(input float64) {
			id = int(input)
		},
	})
	if !v(findMoodIndex(id) != -1, "ID tidak ditemukan!") {
		return
	}
	Temp.id = id
	toPage("MoodView")
}

func appOptionMoodInsert() {
	var mood Mood

	input(Input{
		prompt: "Deskripsi mood: ",
		onSubmit: func(input string) {
			mood.description = input
		},
	})
	if !v(mood.description != "", "Deskripsi mood tidak boleh kosong!") {
		return
	}

	inputNumber(InputNumber{
		prompt: "Skor mood (1-10): ",
		onSubmit: func(input float64) {
			mood.score = int(input)
		},
	})
	if !v(mood.score >= 1 && mood.score <= 10, "Skor harus antara 1 dan 10!") {
		return
	}

	input(Input{
		prompt: "Tanggal (YYYY-MM-DD): ",
		onSubmit: func(input string) {
			mood.date = input
		},
	})
	if !v(validateMoodDate(mood.date), "Tanggal tidak valid! Pastikan formatnya YYYY-MM-DD dan tanggalnya benar.") {
		return
	}
	alert("Catatan mood berhasil ditambahkan!")
	insertMood(mood)
}

func appOptionMoodUpdate() {
	var mood Mood
	if App.currentPage == "MoodView" && Temp.id != -1 {
		mood.id = Temp.id
	} else {
		inputNumber(InputNumber{
			prompt: "ID catatan mood: ",
			onSubmit: func(input float64) {
				mood.id = int(input)
			},
		})
	}

	var index = findMoodIndex(mood.id)
	if !v(index != -1, "ID tidak ditemukan!") {
		return
	}

	input(Input{
		prompt: "Deskripsi mood: (sama) ",
		onSubmit: func(input string) {
			mood.description = input
		},
	})

	if mood.description == "" {
		mood.description = moods[index].description
	}

	inputNumber(InputNumber{
		prompt: "Skor mood (1-10): (" + toString(moods[index].score) + ") ",
		onSubmit: func(input float64) {
			mood.score = int(input)
		},
	})

	if mood.score == 0 {
		mood.score = moods[index].score
	}

	if !v(mood.score >= 1 && mood.score <= 10, "Skor harus antara 1 dan 10!") {
		return
	}

	input(Input{
		prompt: "Tanggal (YYYY-MM-DD): (" + toString(moods[index].date) + ") ",
		onSubmit: func(input string) {
			mood.date = input
		},
	})

	if mood.date == "" {
		mood.date = moods[index].date
	}

	if !v(validateMoodDate(mood.date), "Tanggal tidak valid! Pastikan formatnya YYYY-MM-DD dan tanggalnya benar.") {
		return
	}
	alert("Catatan mood berhasil diperbarui!")
	updateMood(mood.id, mood)
}

func appOptionMoodDelete() {
	var id int
	if App.currentPage == "MoodView" && Temp.id != -1 {
		id = Temp.id
	} else {
		inputNumber(InputNumber{
			prompt: "ID catatan mood: ",
			onSubmit: func(input float64) {
				id = int(input)
			},
		})
	}

	if !v(findMoodIndex(id) != -1, "ID tidak ditemukan!") {
		return
	}

	deleteMood(id)
	alert("Catatan mood berhasil dihapus!")
	if App.currentPage == "MoodView" {
		toPage("Mood")
	}
}

func appDynamicMoodView() []string {
	var index = findMoodIndex(Temp.id)
	if index == -1 {
		return []string{"ID tidak ditemukan!"}
	}

	var m = moods[index]
	return []string{
		fmt.Sprintf("%d | %d/10 | %s", m.id, m.score, m.date),
		m.description,
	}
}

// --

func findMoodIndex(id int) int {
	var i int
	for i = 0; i < len(moods); i++ {
		if moods[i].id == id {
			return i
		}
	}
	return -1
}

func getMaxMoodID() int {
	var maxID = 0
	var i int
	for i = 0; i < len(moods); i++ {
		if moods[i].id > maxID {
			maxID = moods[i].id
		}
	}
	return maxID
}

func validateMoodDate(input string) bool {
	if len(input) != 10 {
		return false
	}

	// validasi format YYYY-MM-DD
	if input[4] != '-' || input[7] != '-' {
		return false
	}

	// validasi YYYY, MM, DD adalah angka
	if !(isNumStr(input[0:4], false) && isNumStr(input[5:7], false) && isNumStr(input[8:10], false)) {
		return false
	}

	// validasi tanggal yang valid
	var year, month, day int
	year = toInt(input[0:4])
	month = toInt(input[5:7])
	day = toInt(input[8:10])
	if !validateDate(year, month, day) {
		return false
	}

	return true
}
