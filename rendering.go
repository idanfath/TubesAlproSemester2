package main

import "fmt"

type RenderData struct {
	multiline []string
	text      string
	table     func() Table // biar dinamis, tiap render build ulang tabel
	breakline bool
	dynamic   func() []string
	options   Options
}

var RenderQ = []RenderData{}
var TempRenderQ = []RenderData{}

func render() {
	var i int
	clearScreen()
	printTopMargin()
	var page = getPage(App.currentPage)
	var foundOptions = false
	// render page content
	for i = 0; i < len(page.content); i++ {
		handleRenderData(page.content[i])
		if len(page.content[i].options) > 0 {
			foundOptions = true
			outputTempRenderQueue()
			showOptions(buildOptions(page.content[i].options))
		}
	}
	// handle no options
	if !foundOptions && page.name != "Exit" {
		if len(App.history) > 0 && !getPage(App.currentPage).noBack {
			outputTempRenderQueue()
			showOptions(buildOptions(Options{}))
		} else {
			exit()
		}
	}
	printBottomMargin()
}

// --

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
func printTopMargin() {
	fmt.Println(repeat("\n", App.topMargin))
}
func printBottomMargin() {
	fmt.Println(repeat("\n", App.bottomMargin))
}
func printMultiline(multiline []string) {
	var i int
	if len(multiline) > 0 {
		for i = 0; i < len(multiline); i++ {
			fmt.Println(cen(multiline[i]))
		}
	}
}

func handleRenderData(data RenderData) {
	if data.breakline {
		fmt.Println("")
	}
	if data.dynamic != nil {
		printMultiline(data.dynamic())
	}
	if data.text != "" {
		fmt.Println(cen(data.text))
	}
	if len(data.multiline) > 0 {
		printMultiline(data.multiline)
	}
	// harus != nil, karna default value fungsi itu nil
	if data.table != nil {
		printTable(data.table())
	}
}

func outputTempRenderQueue() {
	var i int
	if len(TempRenderQ) > 0 {
		fmt.Println("")
		for i = 0; i < len(TempRenderQ); i++ {
			handleRenderData(TempRenderQ[i])
		}
		TempRenderQ = []RenderData{}
	}
}

// simple validation helper, result = false -> render message, result = true -> lanjut
func v(result bool, message string) bool {
	if !result {
		TempRenderQ = []RenderData{{text: message}}
	}
	return result
}

// reusable buat munculi text di next render cycle temporarily
func alert(s string) {
	TempRenderQ = []RenderData{{text: s}}
}
