package convert

import "strconv"

func stringToInt(s string) int {
	intConv, _ := strconv.Atoi(s)
	return intConv
}
