package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/tomkimsour/Advent-of-code/convert"
)

type position struct {
	value   int
	visited bool
}

func initMap(lines []string) [100][100]position {
	var depthMap [100][100]position
	for row, lineStr := range lines {
		line := convert.ArrStrToArrInt(strings.Split(lineStr, ""))
		for col, val := range line {
			depthMap[row][col].value = val
			depthMap[row][col].visited = false
		}
	}
	return depthMap
}

func findAllMin(depthMap [100][100]position) list.List {
	var minList list.List

	for row, line := range depthMap {
		for col, pos := range line {
			var isMinimum bool = true
			if row-1 >= 0 && isMinimum {
				if pos.value >= depthMap[row-1][col].value {
					isMinimum = false
				}
			}
			if row+1 < 100 && isMinimum {
				if pos.value >= depthMap[row+1][col].value {
					isMinimum = false
				}
			}
			if col-1 >= 0 && isMinimum {
				if pos.value >= depthMap[row][col-1].value {
					isMinimum = false
				}
			}
			if col+1 < 100 && isMinimum {
				if pos.value >= depthMap[row][col+1].value {
					isMinimum = false
				}
			}

			if isMinimum {
				minList.PushBack(pos)
			}
		}
	}

	return minList
}

func getRisk(minList list.List) int {
	sum := 0
	for e := minList.Front(); e != nil; e = e.Next() {
		sum += e.Value.(position).value + 1
	}
	return sum
}

func pb1(lines []string) [100][100]position {
	depthMap := initMap(lines)
	minList := findAllMin(depthMap)
	risk := getRisk(minList)
	fmt.Println(risk)
	return depthMap
}

func getNbBassin(depthMap [100][100]position, row, col, previousValue int) int {
	depthMap[row][col].visited = true
	nb := 1
	currentValue := depthMap[row][col]
	if currentValue.value == 9 {
		return 0
	}

	if previousValue <= currentValue.value {
		if row-1 >= 0 {
			if !depthMap[row-1][col].visited {
				nb += getNbBassin(depthMap, row-1, col, currentValue.value)
			}
		}
		if row+1 < 100 {
			if !depthMap[row+1][col].visited {
				nb += getNbBassin(depthMap, row+1, col, currentValue.value)
			}
		}
		if col-1 >= 0 {
			if !depthMap[row][col-1].visited {
				nb += getNbBassin(depthMap, row, col-1, currentValue.value)
			}
		}
		if col+1 < 100 {
			if !depthMap[row][col+1].visited {
				nb += getNbBassin(depthMap, row, col+1, currentValue.value)
			}
		}
	}

	return nb
}

func findBasinsSize(depthMap [100][100]position) []int {
	var basinList []int

	for row, line := range depthMap {
		for col, pos := range line {
			var isMinimum bool = true
			if row-1 >= 0 && isMinimum {
				if pos.value >= depthMap[row-1][col].value {
					isMinimum = false
				}
			}
			if row+1 < 100 && isMinimum {
				if pos.value >= depthMap[row+1][col].value {
					isMinimum = false
				}
			}
			if col-1 >= 0 && isMinimum {
				if pos.value >= depthMap[row][col-1].value {
					isMinimum = false
				}
			}
			if col+1 < 100 && isMinimum {
				if pos.value >= depthMap[row][col+1].value {
					isMinimum = false
				}
			}

			if isMinimum {
				nbBasins := getNbBassin(depthMap, row, col, -1)
				fmt.Println(nbBasins)
				basinList = append(basinList, nbBasins)
			}
		}
	}

	return basinList
}

func getMax(basinList []int) (int, int, int) {
	sort.Ints(basinList)
	// size := len(basinList)
	fmt.Println(basinList)
	// m1 := basinList[size-1]
	// m2 := basinList[size-2]
	// m3 := basinList[size-3]
	// fmt.Println(m1)
	// fmt.Println(m2)
	// fmt.Println(m3)
	return 1, 2, 3
	// return m1, m2, m3
}

func pb2(lines []string) {
	depthMap := initMap(lines)
	basinList := findBasinsSize(depthMap)
	m1, m2, m3 := getMax(basinList)
	fmt.Println(m1 * m2 * m3)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// pb1(lines)
	pb2(lines)
}
