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
	title:        "My CLI App",
	width:        80,
	topMargin:    10,
	bottomMargin: 5,
	history:      []string{},
	currentPage:  "",
}

var KILLSIG bool

func exit() {
	KILLSIG = true
}

func initializeApp() {
	App.currentPage = "Home"
	App.history = []string{App.currentPage}
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
