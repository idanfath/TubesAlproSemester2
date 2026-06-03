package main

import "time"

func pow(base int, power int) int {
	if power <= 0 {
		return 1
	}
	return base * pow(base, power-1)
}

// rands
func randInt(min, max int) int {
	return min + time.Now().Nanosecond()%(max-min+1)
}
func randDaily(min, max int) int {
	var seed = time.Now().Day()*31 + int(time.Now().Month())*7 + time.Now().Year()
	return min + seed%(max-min+1)
}

func sum(arr []int) int {
	var total int
	var i int
	for i = 0; i < len(arr); i++ {
		total += arr[i]
	}
	return total
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
