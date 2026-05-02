package main

type Page struct {
	name    string
	content []RenderData
	noBack  bool
}

var fallback_notfound_page Page
var pages []Page

func getPage(name string) Page {
	var i int
	for i = 0; i < len(pages); i++ {
		if pages[i].name == name {
			return pages[i]
		}
	}
	return fallback_notfound_page
}

func toPage(name string) {
	if name == App.currentPage {
		return
	}
	App.history = newHistory(name)
	App.currentPage = name
	RenderQ = getPage(App.currentPage).content
}

func back() {
	if len(App.history) < 2 {
		toPage("Home")
		return
	}
	toPage(App.history[len(App.history)-2])
}

// kalau page baru blm ada di history, tambahin, kalau udah ada, return history sampai page baru aja
func newHistory(newPage string) []string {
	var historyLength = len(App.history)
	var newHistory []string
	var i int
	for i = 0; i < historyLength; i++ {
		if App.history[i] == newPage {
			break
		}
		newHistory = append(newHistory, App.history[i])
	}
	newHistory = append(newHistory, newPage)
	return newHistory
}
