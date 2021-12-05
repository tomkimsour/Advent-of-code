package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tomkimsour/Advent-of-code/convert"
)

type Point struct {
	x int
	y int
}

func initMap() [][]int {
	ventMap := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		ventMap[i] = make([]int, 1000)
		for j := 0; j < 1000; j++ {
			ventMap[i][j] = 0
		}
	}
	return ventMap
}

func southEast(ventMap [][]int, deb, fin Point) {
	taille := fin.x - deb.x
	for i := 0; i < taille+1; i++ {
		ventMap[deb.x+i][deb.y+i]++
	}
}

func northEast(ventMap [][]int, deb, fin Point) {
	taille := fin.x - deb.x
	for i := 0; i < taille+1; i++ {
		ventMap[deb.x+i][deb.y-i]++
	}
}
func northWest(ventMap [][]int, deb, fin Point) {
	taille := deb.x - fin.x
	for i := 0; i < taille+1; i++ {
		ventMap[deb.x-i][deb.y-i]++
	}
}
func southWest(ventMap [][]int, deb, fin Point) {
	taille := deb.x - fin.x
	for i := 0; i < taille+1; i++ {
		ventMap[deb.x-i][deb.y+i]++
	}
}

func addLineToMap(ventMap [][]int, deb, fin Point) {
	if deb.x == fin.x {
		if deb.y < fin.y {
			for i := deb.y; i < fin.y+1; i++ {
				ventMap[deb.x][i]++
			}
		} else {
			for i := fin.y; i < deb.y+1; i++ {
				ventMap[deb.x][i]++
			}
		}
	} else if deb.y == fin.y {
		if deb.x < fin.x {
			for i := deb.x; i < fin.x+1; i++ {
				ventMap[i][deb.y]++
			}
		} else {
			for i := fin.x; i < deb.x+1; i++ {
				ventMap[i][deb.y]++
			}
		}
	} else {
		if deb.x < fin.x && deb.y < fin.y {
			southEast(ventMap, deb, fin)
		} else if deb.x < fin.x && deb.y > fin.y {
			northEast(ventMap, deb, fin)
		} else if deb.x > fin.x && deb.y > fin.y {
			northWest(ventMap, deb, fin)
		} else if deb.x > fin.x && deb.y < fin.y {
			southWest(ventMap, deb, fin)
		}
	}
}

func getNumberOfoverlap(ventMap [][]int) int {
	var count int = 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if ventMap[i][j] > 1 {
				count++
			}
		}
	}
	return count
}

func watchGrid(ventMap [][]int) {
	fmt.Println("printing map :")
	deb := 800
	for i := 450; i < 900; i++ {
		for j := deb; j < deb+50; j++ {
			fmt.Print(ventMap[i][j])
			fmt.Print(" ")

		}
		fmt.Println()
	}
}

func pb1(lines []string) {
	ventMap := initMap()
	var deb Point
	var fin Point
	for _, line := range lines {
		coordinate := strings.Split(line, " -> ")
		deb.x = convert.StringToInt(strings.Split(coordinate[0], ",")[0])
		deb.y = convert.StringToInt(strings.Split(coordinate[0], ",")[1])
		fin.x = convert.StringToInt(strings.Split(coordinate[1], ",")[0])
		fin.y = convert.StringToInt(strings.Split(coordinate[1], ",")[1])
		addLineToMap(ventMap, deb, fin)
	}
	// watchGrid(ventMap)
	res := getNumberOfoverlap(ventMap)
	fmt.Println(res)
}

func pb2(lines []string) {

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

	pb1(lines)
	pb2(lines)
}
