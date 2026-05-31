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
}

func MoodTable() Table {
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

// --

func appOptionMoodView() {
	var id int
	inputNumber(InputNumber{
		Prompt: "Masukkan ID catatan mood yang ingin dilihat: ",
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
		Prompt: "Deskripsi mood: ",
		onSubmit: func(input string) {
			mood.description = input
		},
	})
	if !v(mood.description != "", "Deskripsi mood tidak boleh kosong!") {
		return
	}

	inputNumber(InputNumber{
		Prompt: "Skor mood (1-10): ",
		onSubmit: func(input float64) {
			mood.score = int(input)
		},
	})
	if !v(mood.score >= 1 && mood.score <= 10, "Skor harus antara 1 dan 10!") {
		return
	}

	input(Input{
		Prompt: "Tanggal (YYYY-MM-DD): ",
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
			Prompt: "ID catatan mood: ",
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
		Prompt: "Deskripsi mood: (sama) ",
		onSubmit: func(input string) {
			mood.description = input
		},
	})

	if mood.description == "" {
		mood.description = moods[index].description
	}

	inputNumber(InputNumber{
		Prompt: "Skor mood (1-10): (" + toString(moods[index].score) + ") ",
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
		Prompt: "Tanggal (YYYY-MM-DD): (" + toString(moods[index].date) + ") ",
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
			Prompt: "ID catatan mood: ",
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
