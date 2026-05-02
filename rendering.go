package main

import "fmt"

type RenderData struct {
	multiline []string
	text      string
	table     string
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
		outputRenderData(page.content[i])
		if len(page.content[i].options) > 0 {
			foundOptions = true
			showOptions(page.content[i].options)
		}
	}
	// render temporary content
	outputTempRenderQueue()
	// handle no options
	if !foundOptions && page.name != "Exit" {
		if len(App.history) > 0 {
			showOptions(buildOptions(Options{}))
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

func outputRenderData(data RenderData) {
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
	if data.table != "" {
		printT(getT(data.table))
	}
}
func outputTempRenderQueue() {
	var i int
	for i = 0; i < len(TempRenderQ); i++ {
		outputRenderData(TempRenderQ[i])
	}
	TempRenderQ = []RenderData{}
}
