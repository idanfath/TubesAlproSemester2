package main

import "fmt"

type Date struct {
	year  int
	month int
	day   int
}

// mengecek apakah sebuah date valid
func isValidDate(date Date) bool {
	if date.year < 0 || date.month < 1 || date.month > 12 || date.day < 1 || date.day > 31 {
		return false
	}

	// pake 13, biar tinggal day <= daysInMonth[month], gaperlu daysInMonth[month-1] karena indexnya beda 1
	var daysInMonth = [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	// handle tahun kabisat
	if date.month == 2 && isLeapYear(date.year) {
		daysInMonth[2] = 29
	}

	if date.day < 1 || date.day > daysInMonth[date.month] {
		return false
	}

	return true
}

// mengecek apakah tahun merupakan tahun kabisat
func isLeapYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		return true
	}
	return false
}

// mengembalikan string dari date dalam format YYYY-MM-DD atau YYYY MM DD
func datestring(date Date, hyphen bool) string {
	// %02d artinya apa bang mesi, %xyd, x = karakter padding, y = ukuran padding, d = integer (tipe data)
	if hyphen {
		return fmt.Sprintf("%04d-%02d-%02d", date.year, date.month, date.day)
	}
	return fmt.Sprintf("%04d %02d %02d", date.year, date.month, date.day)
}

// membandingkan dua date, mengembalikan true jika date di d1 setelah date di d2
func isDateLarger(d1 Date, d2 Date) bool {
	if d1.year != d2.year {
		return d1.year > d2.year
	}
	if d1.month != d2.month {
		return d1.month > d2.month
	}
	return d1.day > d2.day
}

// mengembalikan valid date 1 hari sebelum dari date yang diberikan
func getYesterday(date Date) Date {
	var prevDate = date
	prevDate.day--
	if prevDate.day < 1 {
		prevDate.month--
		if prevDate.month < 1 {
			prevDate.year--
			prevDate.month = 12
		}
		prevDate.day = 31
	}
	for !isValidDate(prevDate) {
		prevDate.day--
	}
	return prevDate
}
