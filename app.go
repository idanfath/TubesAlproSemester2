// disini simpan konfigurasi, data awal aplikasi, dll

package main

import "fmt"

type AppConfig struct {
	title        string
	width        int
	topMargin    int
	bottomMargin int
	history      []string
	pageParam    []any
	currentPage  string
}

type TempData struct {
	lastMotivationID int
	view_id          int
}

var Temp = TempData{
	lastMotivationID: -1,
	view_id:          -1,
}

var App = AppConfig{
	title:        "Mindflow",
	width:        150,
	topMargin:    10,
	bottomMargin: 5,
	history:      []string{},
	currentPage:  "",
}

var KILLSIG bool

func exit() {
	KILLSIG = true
	toPage("Exit")
}

func initializeApp() {
	App.currentPage = "Home"
	App.history = []string{App.currentPage}
	fallback_notfound_page = Page{
		name: "Not Found",
		content: []RenderData{
			{text: "Page Not Found"},
			{options: Options{
				{
					name: "Kembali ke Menu Utama",
					action: func() {
						toPage("Home")
					},
				},
			}},
		},
	}
	pages = []Page{
		{
			name:   "Home",
			noBack: true,
			content: []RenderData{
				{text: "Mindflow - Tugas Besar Alpro 2"},
				{text: "Selamat datang di Mindflow!"},
				{breakline: true},
				{text: "Daily Motivation"},
				{dynamic: func() []string {
					return []string{getDailyMotivation().quote, "- " + getDailyMotivation().author}
				}},
				{options: Options{
					{
						name: "Manajemen Catatan Mood",
						action: func() {
							toPage("Mood")
						},
					},
					{
						name: "Motivasi Lainnya",
						action: func() {
							toPage("Motivation")
						},
					},
					{
						name: "Exit " + App.title,
						action: func() {
							exit()
						},
					},
				}},
			},
		},
		{
			name: "Motivation",
			content: []RenderData{
				{text: "Motivasi Acak ✨"},
				{text: "Dapatkan motivasi baru dengan memilih opsi di bawah!"},
				{breakline: true},
				{dynamic: func() []string {
					var motivation Motivation = getRandomMotivation()
					for motivation.id == Temp.lastMotivationID {
						motivation = getRandomMotivation()
					}
					Temp.lastMotivationID = motivation.id
					return []string{motivation.quote, "- " + motivation.author}
				}},
				{options: Options{
					{
						name: "Dapatkan Motivasi Baru",
						action: func() {
							toPage("Motivation")
						},
					},
				}},
			},
		},
		{
			name: "Mood",
			content: []RenderData{
				{text: "Menu Mood"},
				{text: "Kelola catatan mood harianmu dengan mudah!"},
				{breakline: true},
				{table: "Mood"},
				{options: Options{
					{
						name: "Lihat Catatan Mood",
						action: func() {
							inputNumber(InputNumber{
								Prompt: "Masukkan ID catatan mood yang ingin dilihat: ",
								onSubmit: func(input float64) {
									Temp.view_id = int(input)
								},
							})
							toPage("MoodView")
						},
					},
					{
						name: "Tambah Catatan Mood",
						action: func() {
							var newMood []any
							inputNumber(InputNumber{
								Prompt: "ID catatan mood (angka unik): ",
								onSubmit: func(input float64) {
									newMood = append(newMood, int(input))
								},
							})
							input(Input{
								Prompt: "Deskripsi mood hari ini: ",
								onSubmit: func(input string) {
									newMood = append(newMood, input)
								},
							})
							inputNumber(InputNumber{
								Prompt: "Skor mood hari ini (1-10): ",
								onSubmit: func(input float64) {
									if input < 1 {
										newMood = append(newMood, 0)
										return
									}
									if input > 10 {
										newMood = append(newMood, 10)
										return
									}
									newMood = append(newMood, int(input))
								},
							})
							input(Input{
								Prompt: "Tanggal catatan (YYYY-MM-DD): ",
								onSubmit: func(input string) {
									newMood = append(newMood, input)
								},
							})
							TempRenderQ = []RenderData{{text: "Catatan mood berhasil ditambahkan!"}}
							insertT("Mood", newMood)
						},
					},
					{
						name: "Edit Catatan Mood",
						action: func() {
							var i, id int
							inputNumber(InputNumber{
								Prompt: "Masukkan ID catatan mood yang ingin diedit: ",
								onSubmit: func(input float64) {
									id = int(input)
								},
							})
							var colValues = colValuesT("Mood", 0)
							// kalo ada id yang cocok, edit datanya
							var rowIndex = -1
							for i = 0; i < len(colValues); i++ {
								if colValues[i] == id {
									rowIndex = i
									break
								}
							}
							if rowIndex == -1 {
								TempRenderQ = []RenderData{{text: "ID tidak ditemukan!"}}
								return
							}
							var newMood []any
							newMood = append(newMood, id) // idnya tetep sama
							input(Input{
								Prompt: "Deskripsi mood hari ini: ",
								onSubmit: func(input string) {
									newMood = append(newMood, input)
								},
							})
							inputNumber(InputNumber{
								Prompt: "Skor mood hari ini (1-10): ",
								onSubmit: func(input float64) {
									fmt.Println(input)
									// karna skor cuma dari 1-10, kalo 0/lebih kita anggap user keep skor lama (soalnya kalo kosong, defaultnya 0)
									if input < 1 || input > 10 {
										newMood = append(newMood, "")
										return
									}
									newMood = append(newMood, int(input))
								},
							})
							input(Input{
								Prompt: "Tanggal catatan (YYYY-MM-DD): ",
								onSubmit: func(input string) {
									// kalau kosong, keep tanggal lama
									newMood = append(newMood, input)
								},
							})
							updateRowT("Mood", rowIndex, newMood)
						},
					},
					{
						name: "Hapus Catatan Mood",
						action: func() {
							var id int
							inputNumber(InputNumber{
								Prompt: "Masukkan ID catatan mood yang ingin dihapus: ",
								onSubmit: func(input float64) {
									id = int(input)
								},
							})
							var rowIndex = findFirstInTableCol("Mood", id, 0)
							if rowIndex == -1 {
								TempRenderQ = []RenderData{{text: "ID tidak ditemukan!"}}
								return
							}
							removeT("Mood", rowIndex)
							TempRenderQ = []RenderData{{text: "Catatan mood berhasil dihapus!"}}
						},
					},
				}},
				{options: Options{}},
			},
		},
		{
			name: "MoodView",
			content: []RenderData{
				{text: "Detail Catatan Mood"},
				{breakline: true},
				{dynamic: func() []string {
					// cek dulu id yang mau dilihat ada ga
					var colValues = colValuesT("Mood", 0)
					var rowIndex = -1
					var i int
					for i = 0; i < len(colValues); i++ {
						if colValues[i] == Temp.view_id {
							rowIndex = i
							break
						}
					}
					if rowIndex == -1 {
						return []string{"ID tidak ditemukan!"}
					}
					// kalo ada, tampilkin detailnya
					var table = getT("Mood").rows[rowIndex]
					return []string{
						fmt.Sprintf("-- %v/10 | %v --", table[2], table[3]),
						fmt.Sprintf("%v", table[1]),
					}
				}},
			},
		},
		{
			name:   "Exit",
			noBack: true,
			content: []RenderData{
				{text: "Exitting app.."},
			},
		},
	}
	tables = []Table{
		{
			name:   "Mood",
			header: []string{"ID", "Deskripsi", "Skor", "Tanggal"},
			rows: [][]any{
				{1, "sangat senang karna saya ga remed alpro", 8, "2024-06-01"},
				{2, "saya remed alpro :(", 4, "2024-06-02"},
				{3, "saya diputusin :((((", 2, "2024-06-03"},
				{4, "saya remed alpro lagi :(, udah diputusin, nilai anjlok, ancur dah", 1, "2024-06-04"},
			},
			col_max_length:   []int{5, 40, 5, 12},
			unique_col_index: []int{0},
		},
	}
}
