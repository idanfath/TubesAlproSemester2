package main

import (
	"fmt"
)

// --
func centerString(s string, w int) string {
	return repeat(" ", (w-len(s))/2) + s
}
func cen(s string) string {
	return centerString(s, App.width)
}

// --
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
	var result, currentChar string
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
func lower(s string) string {
	var result string
	var temp byte
	var i int
	// rebuild string, tapi kalo ketemu huruf kapital ubah ke huruf kecil
	for i = 0; i < len(s); i++ {
		temp = s[i]
		if temp >= 'A' && temp <= 'Z' {
			result += string(temp + 32)
			continue
		}
		result += string(temp)
	}
	return result
}
func truncate(s string, maxLength int) string {
	// ("tes", 0) -> ""
	if maxLength <= 0 {
		return ""
	}
	// ("tes", 5) -> "tes"
	if len(s) <= maxLength {
		return s
	}
	// ("testing", 3) -> "tes"
	// kalau bisa jgn masukin kurang dari 3 dah maxLengthnya
	if maxLength <= 3 {
		return s[:maxLength]
	}
	// ("testing", 5) -> "te..."
	return s[:maxLength-3] + "..."
}
