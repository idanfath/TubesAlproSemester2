package main

import "fmt"

type Option struct {
	name   string
	action func()
}
type Options []Option
type Input struct {
	Prompt   string
	onSubmit func(string)
}
type InputNumber struct {
	Prompt   string
	onSubmit func(float64)
}

func printAndListenInput(input Input) {
	var s string
	fmt.Print(center(input.Prompt))
	fmt.Scanln(&s)
	input.onSubmit(s)
}

func printAndListenInputNumber(input InputNumber) {
	var f float64
	fmt.Print(center(input.Prompt))
	fmt.Scanln(&f)
	input.onSubmit(f)
}

func printAndListenOptions(options Options) {
	var longestLen = getOptionsLongestLength(options)
	var optionsAmount int = len(options)
	var i int
	var s string
	for i = 0; i < optionsAmount; i++ {
		s = fmt.Sprintf("[%d] %-*s", i, longestLen, options[i].name)
		fmt.Println(center(s))
	}
	readOption(options)
}
func getOptionsLongestLength(options Options) int {
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
func selectOption(options Options, index int) {
	if index < 0 || index >= len(options) {
		return
	}
	options[index].action()
}
func readOption(options Options) {
	var index int
	fmt.Print(center("Select an option: "))
	fmt.Scan(&index)
	selectOption(options, index)
}
