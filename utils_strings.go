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
func upper(s string) string {
	var result string
	var temp byte
	var i int
	// rebuild string, tapi kalo ketemu huruf kecil ubah ke huruf kapital
	for i = 0; i < len(s); i++ {
		temp = s[i]
		if temp >= 'a' && temp <= 'z' {
			result += string(temp - 32)
			continue
		}
		result += string(temp)
	}
	return result
}
func truncate(s string, maxLength int) string {
	// ("tes", 0) -> ""
	if maxLength == 0 {
		return ""
	}
	// ("tes", -1) -> "tes"
	if maxLength < 0 {
		return s
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

func toInt(s string) int {
	var result int
	var i int
	if len(s) == 0 {
		return 0
	}

	// kalo ada tanda minus, tandai, dan hapus tandanya dulu
	var isNegative bool
	if s[0] == '-' {
		isNegative = true
		s = s[1:]
	}

	for i = 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		// dikurangi '0' karna 48, msial '3' - '0' = 51 - 48 = 3
		result = result*10 + int(s[i]-'0')
	}

	if isNegative {
		result = -result
	}
	return result
}

// iterasi tiap char string, kalo ada yang bukan angka, return false, kalo expectNegative true, boleh ada 1 minus di awal
func isNumStr(s string, expectNegative bool) bool {
	if len(s) == 0 {
		return false
	}
	var i int
	if expectNegative && s[0] == '-' {
		if len(s) == 1 {
			return false
		}
		s = s[1:]
	}
	for i = 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

// simple switch, basically ? : (ini golang knp gapunya fitur gini dah)
func stringswitch(condition bool, trueVal string, falseVal string) string {
	if condition {
		return trueVal
	}
	return falseVal
}
func intSwitch(condition bool, trueVal int, falseVal int) int {
	if condition {
		return trueVal
	}
	return falseVal
}

// untuk lowercase search, tugas pemanggil, bukan tugas fungsi ini.
func contains(str string, substr string) bool {
	var lenStr, lenSub, i int
	lenStr = len(str)
	lenSub = len(substr)

	// kalo substring lebih panjang otomatis gagal
	if lenSub > lenStr || lenSub == 0 {
		return false
	}

	if lenStr == lenSub {
		return str == substr
	}

	// cari dari index 0 sampai index yang masih muat substringnya
	for i = 0; i <= lenStr-lenSub; i++ {
		// misal str = "testing", substr = "STI", iterasi 1 bandingkan [tes]ting dengan sti, iterasi 2 bandingkan te[sti]ng dengan sti
		if str[i:i+lenSub] == substr {
			return true
		}
	}
	return false
}
