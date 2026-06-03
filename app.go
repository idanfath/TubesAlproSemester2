// disini simpan konfigurasi, data awal aplikasi, dll

package main

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
	lastMotivationID any
	id               int
}

var Temp = TempData{
	lastMotivationID: -1,
	id:               -1,
}

var App = AppConfig{
	title:        "Mindflow",
	width:        150,
	topMargin:    2,
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
				{dynamic: appDynamicGetRandomMotivation},
				{options: Options{
					{
						name: "Menu Pencatatan Mood",
						action: func() {
							toPage("Mood")
						},
					},
					{
						name: "Menu Daftar Tugas",
						action: func() {
							toPage("Task")
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
				{dynamic: appDynamicGetRandomMotivation},
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
				{table: MoodTable},
				{
					options: Options{
						{name: "Cari Catatan Mood", action: appOptionMoodSearch},
						{name: "Lihat Detail Catatan Mood", action: appOptionMoodView},
						{name: "Tambah Catatan Mood", action: appOptionMoodInsert},
						{name: "Ubah Catatan Mood", action: appOptionMoodUpdate},
						{name: "Hapus Catatan Mood", action: appOptionMoodDelete},
					},
				},
			},
		},
		{
			name: "MoodView",
			content: []RenderData{
				{text: "Detail Catatan Mood"},
				{breakline: true},
				{dynamic: appDynamicMoodView},
				{options: Options{
					{name: "Ubah Catatan Mood", action: appOptionMoodUpdate},
					{name: "Hapus Catatan Mood", action: appOptionMoodDelete},
				}},
			},
		},
		{
			name: "MoodSearchResult",
			content: []RenderData{
				{text: "Hasil Pencarian Catatan Mood"},
				{dynamic: func() []string {
					return []string{"Ditemukan " + toString(len(searchResultsMood)) + " catatan mood yang cocok:"}
				}},
				{breakline: true},
				{table: func() Table {
					return buildMoodTable(searchResultsMood)
				}},
			},
		},
		{
			name: "Task",
			content: []RenderData{
				{text: "Menu Daftar Tugas"},
				{breakline: true},
				{table: TaskTable},
				{options: Options{
					{name: "Cari Tugas", action: appOptionTaskSearch},
					{name: "Lihat Detail Tugas", action: appOptionTaskView},
					{name: "Tambah Tugas", action: appOptionTaskInsert},
					{name: "Ubah Tugas", action: appOptionTaskUpdate},
					{name: "Hapus Tugas", action: appOptionTaskDelete},
				}},
			},
		},
		{
			name: "TaskView",
			content: []RenderData{
				{text: "Detail Tugas"},
				{breakline: true},
				{dynamic: appDynamicTaskView},
				{options: Options{
					{name: "Tandai Selesai/Belum Selesai", action: appOptionToggleTaskStatus},
					{name: "Ubah Detail Tugas", action: appOptionTaskUpdate},
					{name: "Hapus Tugas", action: appOptionTaskDelete},
				}},
			},
		},
		{
			name: "TaskSearchResult",
			content: []RenderData{
				{text: "Hasil Pencarian Tugas"},
				{dynamic: func() []string {
					return []string{"Ditemukan " + toString(len(searchResultsTask)) + " tugas yang cocok:"}
				}},
				{breakline: true},
				{table: func() Table {
					return buildTaskTable(searchResultsTask)
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
}
