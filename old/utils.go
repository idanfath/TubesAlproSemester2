package main

import (
	"fmt"
	"time"
)

func centeredString(s string, w int) string {
	return repeat(" ", (w-len(s))/2) + s
}
func center(s string) string {
	return centeredString(s, App.Width)
}

func repeat(s string, count int) string {
	var result string
	var i int
	for i = 0; i < count; i++ {
		result += s
	}
	return result
}
func toString(s any) string {
	return fmt.Sprint(s)
}
func replace(s string, old string, new string) string {
	var length = len(s)
	var i int
	var result string
	var currentChar string
	for i = 0; i < length; i++ {
		currentChar = string(s[i])
		if currentChar == old {
			result += new
		} else {
			result += currentChar
		}
	}
	return result
}

// random dari 0-31 buat daily motivation
func getRandomDate() int {
	var x int = (time.Now().Day() + int(time.Now().Month()) + time.Now().Year() - 2) % 31
	return x
}
func getRandomInt(max int) int {
	var x int = time.Now().Nanosecond() % max
	return x
}

func power(base int, pow int) int {
	if pow < 0 {
		return 1
	}
	return base * power(base, pow-1)
}

func sortIntArray(array []int) {}

// TODO
