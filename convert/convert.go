package convert

import "strconv"

func StringToInt(s string) int {
	intConv, _ := strconv.Atoi(s)
	return intConv
}
