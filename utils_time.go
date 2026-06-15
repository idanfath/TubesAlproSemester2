package main

import "fmt"

type Date struct {
	year  int
	month int
	day   int
}

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

func isLeapYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		return true
	}
	return false
}

func datestring(date Date) string {
	// %02d artinya apa bang mesi, %xyd, x = karakter padding, y = ukuran padding, d = integer (tipe data)
	return fmt.Sprintf("%04d-%02d-%02d", date.year, date.month, date.day)
}

// // cari tanggal selanjutnya yg msh valid
// func getNextValidDate(date Date) Date {
// 	var nextDate Date = date
// 	for !isValidDate(nextDate) && nextDate != date {
// 		nextDate.day++
// 		if nextDate.day > 31 {
// 			nextDate.month++
// 			nextDate.day = 1
// 		}
// 		if nextDate.month > 12 {
// 			nextDate.year++
// 		}
// 	}
// 	return nextDate
// }
