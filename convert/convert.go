package convert

import (
	"log"
	"strconv"
)

func StringToInt(s string) int {
	intConv, ok := strconv.Atoi(s)
	if ok != nil {
		log.Fatal("bad string conversion")
	}
	return intConv
}

func ArrStrToArrInt(s []string) []int {
	var res []int
	for _, i := range s {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		res = append(res, j)
	}
	return res
}
