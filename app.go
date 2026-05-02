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
			name:   "Exit",
			noBack: true,
			content: []RenderData{
				{text: "BABABABABABAYYY!!!! see you kapan kapan!"},
			},
		},
	}
	tables = []Table{
		{
			name:   "Food",
			header: []string{"ID", "Name", "Favorite Food"},
			rows: [][]any{
				{1, "Alice", "Pizza"},
				{2, "Bob", "Sushi"},
				{3, "Charlie", "Burger"},
			},
		},
	}
}
