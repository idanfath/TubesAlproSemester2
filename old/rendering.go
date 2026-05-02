package main

import (
	"fmt"
)

type RenderData struct {
	multiline []string
	text      string
	tableName string
	options   Options
	emptyline bool
	inline    func() []string
}

type Page struct {
	name    string
	content []RenderData
	noBack  bool
}

var RenderQueue = []RenderData{}

var history []string
var currentPage string
var pages []Page
var temporaryRenderData RenderData

// initialize variabel page, kalau di init langsung berpotensi initialization cycle
func initPage() {
	tablesStorage = []TablesStorage{
		{
			name: "tempTable",
			table: Table{
				header: []string{"ID", "Name", "Age"},
				rows: [][]any{
					{1, "Alice", 30},
					{2, "Bob", 25},
					{3, "Charlie", 35},
				},
			},
		},
		{
			name: "todo",
			table: Table{
				header: []string{"ID", "Tugas", "Kesulitan [1-5]"},
				rows: [][]any{
					{1, "Daftar Lomba Anu", 5},
				},
			},
		},
	}

	pages = []Page{
		{
			name:   "Home",
			noBack: true,
			content: []RenderData{
				{text: "Home Page"},
				{multiline: []string{"Welcome to the Home Page!", "This is a simple CLI application."}},
				{emptyline: true},
				{inline: func() []string {
					return []string{"Motivasi Harian:", getDailyMotivation().quote, getDailyMotivation().source}
				}},
				{emptyline: true},
				{options: Options{
					{
						name: "Go to About Page",
						action: func() {
							toPage("About")
						},
					},
					{
						name: "Go to Data Table Page",
						action: func() {
							toPage("Data Table")
						},
					},
					{
						name: "To-do List",
						action: func() {
							toPage("Todo")
						},
					},
					{
						name: "Motivasi Lainnya",
						action: func() {
							toPage("Motivasi")
						},
					},
					{
						name: "Exit",
						action: func() {
							exitApp()
						},
					},
				}},
			},
		},
		{
			name: "Todo",
			content: []RenderData{
				{text: "Todo List ✨"},
				{tableName: "todo"},
				{options: Options{
					{
						name: "New Data",
						action: func() {
							var id, difficulty int
							var name string
							printAndListenInputNumber(InputNumber{
								Prompt: "ID Tugas: ",
								onSubmit: func(input float64) {
									id = int(input)
								},
							})
							printAndListenInput(Input{
								Prompt: "Nama Tugas (Gunakan _ untuk spasi): ",
								onSubmit: func(input string) {
									name = replace(input, "_", " ")
								},
							})
							printAndListenInputNumber(InputNumber{
								Prompt: "Tingkat Kesulitan (1-5): ",
								onSubmit: func(input float64) {
									difficulty = int(input)
									if difficulty > 5 {
										difficulty = 5
									} else if difficulty < 1 {
										difficulty = 1
									}
								},
							})
							appendToTable("todo", []any{id, name, difficulty})
						},
					},
					{
						name: "Delete Data",
						action: func() {
							var index int
							printAndListenInputNumber(InputNumber{
								Prompt: "Masukkan Indeks: ",
								onSubmit: func(input float64) {
									index = int(input)
								},
							})
							removeRow("tempTable", index)
						},
					},
				}},
			},
		},
		{
			name: "Data Table",
			content: []RenderData{
				{text: "Data Table Page"},
				{tableName: "tempTable"},
			},
		},
		{
			name: "Motivasi",
			content: []RenderData{
				{text: "Motivasi Acak ✨"},
				{inline: func() []string {
					return []string{getRandomMotivation().quote, getRandomMotivation().source}
				}},
				{options: Options{
					{
						name: "Dapatkan motivasi baru",
						action: func() {
							toPage("Motivasi")
						},
					},
				}},
			},
		},
		{
			name: "Exit",
			content: []RenderData{
				{text: "Thank you for using the app!"},
			},
			noBack: true,
		},
	}
	currentPage = "Home"
	history = append(history, currentPage)
	RenderQueue = getPage(currentPage).content
}

// navigation: mundur 1 page
func back() {
	if len(history) < 2 {
		toPage("Home")
		return
	}
	toPage(history[len(history)-2])
}

// navigation: ganti page
func toPage(name string) {
	if name == currentPage {
		return
	}
	history = buildNewHistory(history, name)
	currentPage = name
	RenderQueue = getPage(currentPage).content
}

// kalau page baru blm ada di history, tambahin, kalau udah ada, return history sampai page baru aja
func buildNewHistory(history []string, newPage string) []string {
	var historyLength = len(history)
	var newHistory []string
	var i int
	for i = 0; i < historyLength; i++ {
		if history[i] == newPage {
			break
		}
		newHistory = append(newHistory, history[i])
	}
	newHistory = append(newHistory, newPage)
	return newHistory
}

// buat get page berdasarkan nama (string)
func getPage(name string) Page {
	for i := range pages {
		if pages[i].name == name {
			return pages[i]
		}
	}
	return Page{
		name:   "Not Found",
		noBack: false,
		content: []RenderData{
			{text: "Page not found"},
			{options: Options{
				{
					name: "Back to Home",
					action: func() {
						toPage("Home")
					},
				},
			}},
		},
	}
}

// build options, biar dinamis kalau ada historynya otomatis ada option back
func buildOptions(history []string, options Options) Options {
	if len(history) > 0 && !getPage(currentPage).noBack {
		return append(Options{
			{
				name: "Back",
				action: func() {
					back()
				},
			},
		}, options...)
	}
	return options
}

func setTemporaryRenderData(data RenderData) {
	temporaryRenderData = data
}
func printMultiline(multiline []string) {
	var j int
	if len(multiline) > 0 {
		for j = 0; j < len(multiline); j++ {
			printText(multiline[j])
		}
	}
}
func printText(text string) {
	fmt.Println(center(text))
}
func dynamicPrint(data RenderData) {
	if data.emptyline {
		fmt.Println("")
	}
	if data.inline != nil {
		printMultiline(data.inline())
	}
	if data.text != "" {
		printText(data.text)
	}
	if len(data.multiline) > 0 {
		printMultiline(data.multiline)
	}
	if data.tableName != "" {
		printTable(getTableByName(data.tableName))
	}
}

func printTemporaryRenderData() {
	dynamicPrint(temporaryRenderData)
	temporaryRenderData = RenderData{}
}

// fungsi render utama, clear layar dan render ulang konten
func render() {
	var i int
	clearScreen()
	printTopMargin()
	var foundOptions = false
	var currentPage = getPage(currentPage)
	var currentPageData = currentPage.content
	for i = 0; i < len(currentPageData); i++ {
		dynamicPrint(currentPageData[i])
		if len(currentPageData[i].options) > 0 {
			foundOptions = true
			printAndListenOptions(buildOptions(history, currentPageData[i].options))
		}
	}
	printTemporaryRenderData()
	if !foundOptions && currentPage.name != "Exit" {
		if len(history) > 0 {
			printAndListenOptions(buildOptions(history, Options{}))
		} else {
			exitApp()
		}
	}
}

// rendering utility functions
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func printTopMargin() {
	fmt.Println(repeat("\n", App.TopMargin))
}
