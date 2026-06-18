package main

// menambahkan spasi di depan string agar string terlihat di tengah layar (asumsi layar 75 karakter)
func cen(s string) string {
	return repeat(" ", (75-len(s))/2) + s
}

// mengembalikan string yang merupakan hasil pengulangan string s sebanyak count kali
func repeat(s string, count int) string {
	var result string
	var i int
	for i = 0; i < count; i++ {
		result += s
	}
	return result
}

// mengembalikan string dimana semua karakter old diganti dengan karakter new
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

// mengembalikan string dalam format lowecase
func lower(s string) string {
	var result string
	var temp byte
	var i int
	// rebuild string, tapi kalo ketemu huruf kapital ubah ke huruf kecil
	for i = 0; i < len(s); i++ {
		temp = s[i]
		if temp >= 'A' && temp <= 'Z' {
			result += string(temp + 32)
		} else {
			result += string(temp)
		}
	}
	return result
}

// mengembalikan string dalam format uppercase
func upper(s string) string {
	var result string
	var temp byte
	var i int
	// rebuild string, tapi kalo ketemu huruf kecil ubah ke huruf kapital
	for i = 0; i < len(s); i++ {
		temp = s[i]
		if temp >= 'a' && temp <= 'z' {
			result += string(temp - 32)
		} else {
			result += string(temp)
		}
	}
	return result
}

// simple switch, meniru inline if else (: ?) di bahasa lain
func ifstring(condition bool, trueVal string, falseVal string) string {
	if condition {
		return trueVal
	}
	return falseVal
}

// mengecek apakah string mengandung substring tertentu, case-sensitive
func contains(str string, substr string) bool {
	var lenStr, lenSub, i, j int
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
		var match bool = true
		j = 0
		for j < lenSub && match {
			if str[i+j] != substr[j] {
				match = false
			} else {
				j++
			}
		}
		if match {
			return true
		}
	}
	return false
}

// memotong string dengan panjang maksimal n
func truncate(s string, n int) string {
	var newstring string
	var i int
	if n > len(s) {
		return s
	}
	for i = 0; i < n; i++ {
		newstring = newstring + string(s[i])
	}
	return newstring + "..."
}
