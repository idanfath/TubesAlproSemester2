package main

import "fmt"

// return options paling panjang
func longestOptionName(options Options) int {
	var max int = 0
	var i int
	var optionsAmount int = len(options)
	for i = 0; i < optionsAmount; i++ {
		if len(options[i].name) > max {
			max = len(options[i].name)
		}
	}
	return max
}

func buildOptions(options Options) Options {
	var result Options
	var i int
	if len(App.history) > 0 && !getPage(App.currentPage).noBack {
		result = append(result, Option{
			name: "Back",
			action: func() {
				back()
			},
		})
	}
	for i = 0; i < len(options); i++ {
		result = append(result, options[i])
	}
	return result
}

func runOption(options Options, index int) {
	if index < 0 || index >= len(options) || options[index].action == nil {
		return
	}
	options[index].action()
}

func readOption(options Options) {
	var index int
	fmt.Print(cen("Select an option: "))
	fmt.Scan(&index)
	runOption(options, index)
}
