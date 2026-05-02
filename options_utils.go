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

// tambahkan opsi back jika memungkinkan, pakai "..."
// kalau gaboleh pake "..." bisa dihardcode di function showOptions dan runOptions
func buildOptions(options Options) Options {
	if len(App.history) > 0 && !getPage(App.currentPage).noBack {
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

func runOption(options Options, index int) {
	if index < 0 || index >= len(options) || options[index].action == nil {
		return
	}
	options[index].action()
}

func readOption(options Options) {
	var index int
	var prompt string = fmt.Sprintf("Select an option [0-%d]: ", len(options)-1)
	fmt.Print(cen(prompt))
	fmt.Scan(&index)
	runOption(options, index)
}
