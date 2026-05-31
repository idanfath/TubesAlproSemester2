package main

func validateDate(year, month, day int) bool {
	if year < 0 || month < 1 || month > 12 {
		return false
	}

	// pake 13, biar tinggal day <= daysInMonth[month], gaperlu daysInMonth[month-1] karena indexnya beda 1
	var daysInMonth = [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	// handle tahun kabisat
	if month == 2 && isLeapYear(year) {
		daysInMonth[2] = 29
	}

	if day < 1 || day > daysInMonth[month] {
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

func minuteToTime(minutes int) string {
	var hours, mins int
	var s string
	hours = minutes / 60
	mins = minutes % 60
	if hours > 0 {
		s += toString(hours) + " Jam "
	}
	if mins > 0 {
		s += toString(mins) + " Menit"
	}
	if s == "" {
		s = "0 Menit"
	}
	return s
}
