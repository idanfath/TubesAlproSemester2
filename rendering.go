package main

import "fmt"

type RenderData struct {
	multiline []string
	text      string
	table     Table
	options   Options
}

var RenderQueue = []RenderData{}

type Page struct {
	Name    string
	Content []RenderData
}

var history []string
var currentPage string

var pages []Page

func init() {
	pages = []Page{
		{
			Name: "Home",
			Content: []RenderData{
				{text: "Home Page"},
				{multiline: []string{"Welcome to the Home Page!", "This is a simple CLI application."}},
				{table: Table{
					Header: []string{"ID", "Name", "Age"},
					Rows: [][]any{
						{1, "Alice", 30},
						{2, "Bob", 25},
						{3, "Charlie", 35},
					},
				}},
				{options: Options{
					{
						Name: "Go to About Page",
						Action: func() {
							toPage("About")
						},
					},
					{
						Name: "Exit",
						Action: func() {
							exitApp()
						},
					},
				}},
			},
		},
		{
			Name: "About",
			Content: []RenderData{
				{text: "About Page"},
			},
		},
		{
			Name: "Exit",
			Content: []RenderData{
				{multiline: []string{"Exiting...", "Goodbye!"}},
			},
		},
	}
	currentPage = "Home"
	RenderQueue = getPage(currentPage).Content
}

func back() {
	if len(history) == 0 {
		return
	}
	// page terakhir jadi page skrg
	currentPage = history[len(history)-1]
	// hapus page terakhir dari history
	history = history[0 : len(history)-1]
	RenderQueue = getPage(currentPage).Content
}

func toPage(name string) {
	if name == currentPage {
		return
	}
	history = append(history, currentPage)
	currentPage = name
	RenderQueue = getPage(currentPage).Content
}

func getPage(name string) Page {
	for i := range pages {
		if pages[i].Name == name {
			return pages[i]
		}
	}
	return Page{
		Name:    "Not Found",
		Content: []RenderData{{text: "Page not found"}},
	}
}

func buildOptions(history []string, options Options) Options {
	if len(history) > 0 {
		return append(Options{
			{
				Name: "Back",
				Action: func() {
					back()
				},
			},
		}, options...)
	}
	return options
}

func render() {
	var i, j int
	clearScreen()
	printTopMargin()
	var foundOptions = false
	var currentPage = getPage(currentPage)
	var currentPageData = currentPage.Content
	var currentPageName = currentPage.Name
	for i = 0; i < len(currentPageData); i++ {
		if len(currentPageData[i].multiline) > 0 {
			for j = 0; j < len(currentPageData[i].multiline); j++ {
				fmt.Println(center(currentPageData[i].multiline[j]))
			}
		}
		if currentPageData[i].text != "" {
			fmt.Println(center(currentPageData[i].text))
		}
		if len(currentPageData[i].table.Header) > 0 {
			printTable(currentPageData[i].table)
		}
		if len(currentPageData[i].options) > 0 {
			foundOptions = true
			printAndListenOptions(buildOptions(history, currentPageData[i].options))
		}
	}
	if !foundOptions && currentPageName != "Exit" {
		if len(history) > 0 {
			printAndListenOptions(buildOptions(history, Options{}))
		} else {
			RenderQueue = []RenderData{{text: "Exit reason: No options found"}}
			exitApp()
		}
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func printTopMargin() {
	fmt.Println(repeat("\n", App.TopMargin))
}
