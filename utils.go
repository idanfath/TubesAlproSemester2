package mindflow

import "fmt"

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
