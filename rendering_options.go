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

func input(input Input) {
	var s string
	fmt.Print(cen(input.Prompt))
	fmt.Scanln(&s)
	input.onSubmit(s)
}

func inputNumber(input InputNumber) {
	var f float64
	fmt.Print(cen(input.Prompt))
	fmt.Scanln(&f)
	input.onSubmit(f)
}

func showOptions(options Options) {
	var strLen = longestOptionName(options)
	var i int
	var s string
	for i = 0; i < len(options); i++ {
		s = fmt.Sprintf("[%d] %-*s", i, strLen, options[i].name)
		fmt.Println(cen(s))
	}
	readOption(options)
}
