package main

import "fmt"

type Mood struct {
	id          int
	description string
	score       int
	date        Date
}

const MAX_MOOD = 9999

var moods = [MAX_MOOD]Mood{
	{id: 1, description: "Hari yang menyenangkan!", score: 8, date: Date{year: 2024, month: 6, day: 1}},
	{id: 2, description: "Agak stres dengan pekerjaan. Sedih banget pokoknya diputusin terus remedial aduhhh sad bgt coy", score: 1, date: Date{year: 2024, month: 6, day: 2}},
	{id: 3, description: "Merasa sangat bahagia!", score: 9, date: Date{year: 2024, month: 6, day: 3}},
	{id: 4, description: "Biasa aja, flat banget seharian ga ngapa-ngapain.", score: 5, date: Date{year: 2024, month: 6, day: 4}},
	{id: 5, description: "Dapet sanksi telat masuk kantor, apes banget.", score: 3, date: Date{year: 2024, month: 6, day: 5}},
	{id: 6, description: "Project akhirnya kelar dan diapprove client, lega!", score: 8, date: Date{year: 2024, month: 6, day: 6}},
	{id: 7, description: "Mulai pelihara kucing baru, gemes banget bikin mood naik.", score: 9, date: Date{year: 2024, month: 6, day: 7}},
	{id: 8, description: "Badan agak meriang dan pusing, cuma rebahan doang.", score: 4, date: Date{year: 2024, month: 6, day: 8}},
	{id: 9, description: "Minggu produktif, berhasil beres-beres kamar yang berantakan.", score: 7, date: Date{year: 2024, month: 6, day: 9}},
	{id: 10, description: "Gaji masuk! Waktunya self reward makan enak.", score: 10, date: Date{year: 2024, month: 6, day: 10}},
	{id: 11, description: "Sakit gigi melanda, ga bisa tidur semalaman.", score: 2, date: Date{year: 2024, month: 6, day: 11}},
	{id: 12, description: "Kondisi membaik, dapet kopi gratis dari temen kantor.", score: 6, date: Date{year: 2024, month: 6, day: 12}},
	{id: 13, description: "Terjebak macet total pas hujan deras, capek banget di jalan.", score: 3, date: Date{year: 2024, month: 6, day: 13}},
	{id: 14, description: "Nonton konser band favorit, capek tapi seru parah!", score: 9, date: Date{year: 2024, month: 6, day: 14}},
	{id: 15, description: "Santai di rumah sambil marathon series bareng keluarga.", score: 8, date: Date{year: 2024, month: 6, day: 15}},
}

var moodCount int = 15

func showMoods() {
	var i int
	for i = 0; i < moodCount; i++ {
		fmt.Printf("[%d] Mood %d/10 | %s | %s\n", moods[i].id, moods[i].score, datestring(moods[i].date), truncate(moods[i].description, (App.width/3)*2))
	}
	show("")
}

func pencatatanMood(showItems bool) {
	title("Halaman Pencatatan Mood")
	if showItems {
		showMoods()
	}
	var option int = getOptions([]string{"Lihat", "Tambah", "Edit", "Hapus", "Beranda"})
	switch option {
	case 1:
		optionLihatMood()
	case 2:
		optionTambahMood()
	case 3:
		optionEditMood()
	case 4:
		optionHapusMood()
	case 5:
		return
	}
}

func optionEditMood() {
	var mood, oldMood Mood
	var idx int = -1
	for idx == -1 {
		idx = getMoodIdxFromId()
	}
	oldMood = moods[idx]
	mood.id = oldMood.id

	show("Kosongkan untuk menggunakan data lama")
	for mood.score < 1 || mood.score > 10 {
		showinline(fmt.Sprintf("Masukkan skor mood (1-10): (%d) ", oldMood.score))
		fmt.Scanln(&mood.score)
		if mood.score == 0 {
			mood.score = oldMood.score
		}
	}

	showinline(fmt.Sprintf("Masukkan deskripsi mood (contoh: sangat_senang): (%s) ", truncate(oldMood.description, 12)))
	fmt.Scanln(&mood.description)
	if len(mood.description) == 0 {
		mood.description = oldMood.description
	}
	mood.description = replace(mood.description, "_", " ")

	for !isValidDate(mood.date) {
		showinline(fmt.Sprintf("Masukkan tanggal pencatatan mood (contoh: 2024 12 17): (%s) ", datestring(oldMood.date)))
		fmt.Scanln(&mood.date.year)
		if mood.date.year == 0 {
			mood.date = oldMood.date
		} else {
			fmt.Scanln(&mood.date.month, &mood.date.day)
		}
	}

	updateMood(idx, mood)
	pencatatanMood(true)
}

func optionTambahMood() {
	var mood Mood
	var highestIdMood = moods[getHighestIdMoodIdx()]
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
		fmt.Scanln(&mood.date.year, &mood.date.month, &mood.date.day)
	}
	insertMood(mood)
	pencatatanMood(true)
}

func optionHapusMood() {
	var idx int = -1
	for idx == -1 {
		idx = getMoodIdxFromId()
	}
	deleteMood(idx)
	pencatatanMood(true)
}

func optionLihatMood() {
	var idx int = -1
	for idx == -1 {
		idx = getMoodIdxFromId()
	}
	show("")
	show(fmt.Sprintf("[%d] %d/10 | %s", moods[idx].id, moods[idx].score, datestring(moods[idx].date)))
	show(moods[idx].description)
	show("")
	pencatatanMood(false)
}

// sequential search, return indexnya
func findMood(id int) int {
	var i int
	for i = 0; i < moodCount; i++ {
		if moods[i].id == id {
			return i
		}
	}
	return -1
}

func getHighestIdMoodIdx() int {
	var max_idx = 0
	var i int
	for i = 0; i < moodCount; i++ {
		if moods[i].id > moods[max_idx].id {
			max_idx = i
		}
	}
	return max_idx
}

func getMoodIdxFromId() int {
	showinline("Masukkan ID Mood: ")
	var x int
	fmt.Scanln(&x)
	return findMood(x)
}

func deleteMood(idx int) {
	var newMood [MAX_MOOD]Mood
	var j, i int
	j = 0
	for i = 0; i < moodCount; i++ {
		if i == idx {
			continue
		}
		newMood[j] = moods[i]
		j++
	}
	moods = newMood
	moodCount = j
}

func insertMood(mood Mood) {
	moods[moodCount] = mood
	moodCount++
}

func updateMood(idx int, mood Mood) {
	moods[idx] = mood
}
