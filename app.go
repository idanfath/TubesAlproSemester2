package main

type AppConfig struct {
	title        string
	width        int
	topMargin    int
	bottomMargin int
	history      []string
	currentPage  string
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
				{breakline: true},
				{options: Options{
					{
						name: "Manajemen Catatan Mood",
						action: func() {
							toPage("MenuMood")
						},
					},
					{
						name: "Manajemen Daftar Tugas",
						action: func() {
							toPage("MenuTask")
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
					return []string{getRandomMotivation().quote, "- " + getRandomMotivation().author}
				}},
				{breakline: true},
			},
		},
		{
			name: "MenuMood",
			content: []RenderData{
				{text: "MENU CATATAN MOOD"},
				{breakline: true},
				{options: Options{
					{name: "Tambah Catatan Mood", action: func() { tambahMood() }},
					{name: "Tampil Semua Mood", action: func() { tampilMood() }},
					{name: "Ubah Catatan Mood", action: func() { ubahMood() }},
					{name: "Hapus Catatan Mood", action: func() { hapusMood() }},
				}},
			},
		},
		{
			name: "MenuTask",
			content: []RenderData{
				{text: "MENU DAFTAR TUGAS"},
				{breakline: true},
				{options: Options{
					{name: "Tambah Tugas", action: func() { tambahtask() }},
					{name: "Tampil Semua Tugas", action: func() { tampilTask() }},
					{name: "Ubah Tugas", action: func() { ubahtask() }},
					{name: "Hapus Tugas", action: func() { hapustask() }},
				}},
			},
		},
		{
			name:   "Exit",
			noBack: true,
			content: []RenderData{
				{text: "BABABABABABAYYY!!!! see you kapan kapan!"},
			},
		},
	}
	tables = []Table{}
}
