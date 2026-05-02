package main

import "fmt"

// options
func printAndListenOptions(options Options) {
	var longestLen = getOptionsLongestLength(options)
	var optionsAmount int = len(options)
	var i int
	var s string
	for i = 0; i < optionsAmount; i++ {
		s = fmt.Sprintf("[%d] %-*s", i, longestLen, options[i].Name)
		fmt.Println(center(s))
	}
	readOption(options)
}
func getOptionsLongestLength(options Options) int {
	var max int = 0
	var i int
	var optionsAmount int = len(options)
	for i = 0; i < optionsAmount; i++ {
		if len(options[i].Name) > max {
			max = len(options[i].Name)
		}
	}
	return max
}
func selectOption(options Options, index int) {
	if index < 0 || index >= len(options) {
		fmt.Println("Invalid option")
		return
	}
	options[index].Action()
}
func readOption(options Options) {
	var index int
	fmt.Print(center("Select an option: "))
	fmt.Scan(&index)
	selectOption(options, index)
}
