package main

import "time"

func pow(base int, power int) int {
	if power <= 0 {
		return 1
	}
	return base * pow(base, power-1)
}

// bubble sort, gapapa yang penting ada, kalau boleh pakai import, ganti dengan lib
func sortArrInt(arr []int) []int {
	var i, j int
	for i = 0; i < len(arr); i++ {
		for j = 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				// swap
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// rands
func randInt(min, max int) int {
	return min + time.Now().Nanosecond()%(max-min+1)
}
func randDaily(min, max int) int {
	var seed = time.Now().Day()*31 + int(time.Now().Month())*7 + time.Now().Year()
	return min + seed%(max-min+1)
}
