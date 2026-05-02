package mindflow

import "fmt"

// options
func printOptions(options Options) {
	var longestLen = getOptionsLongestLength(options)
	var optionsAmount int = len(options)
	var i int
	var s string
	for i = 0; i < optionsAmount; i++ {
		s = fmt.Sprintf("[%d] %-*s", i, longestLen, options[i].Name)
		fmt.Println(center(s))
	}
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
