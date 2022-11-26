package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/tomkimsour/Advent-of-code/convert"
)

type Dumbo struct {
	intensite   int
	highlighted bool
}

func initOctopus(lines []string) [][]Dumbo {
	tab := make([][]Dumbo, len(lines))
	for i := 0; i < len(lines); i++ {
		tab[i] = make([]Dumbo, len(lines[0]))
		for j := 0; j < len(lines[0]); j++ {
			var d Dumbo
			d.highlighted = false
			d.intensite = convert.StringToInt(string(lines[i][j]))
			tab[i][j] = d
		}
	}
	return tab
}

func updateNeighbours(dumbo [][]Dumbo, row, col, width, length int) {
	if row-1 >= 0 {
		dumbo[row-1][col].intensite++
		if dumbo[row-1][col].intensite == 9 && !dumbo[row-1][col].highlighted {
			dumbo[row-1][col].highlighted = true
			updateNeighbours(dumbo, row-1, col, width, length)
		}
	}
	if row+1 < length {
		dumbo[row+1][col].intensite++
		if dumbo[row+1][col].intensite == 9 && !dumbo[row+1][col].highlighted {
			dumbo[row+1][col].highlighted = true
			updateNeighbours(dumbo, row+1, col, width, length)
		}
	}
	if col-1 >= 0 {
		dumbo[row][col-1].intensite++
		if dumbo[row][col-1].intensite == 9 && !dumbo[row][col-1].highlighted {
			dumbo[row][col-1].highlighted = true
			updateNeighbours(dumbo, row, col-1, width, length)
		}
	}
	if col+1 < width {
		dumbo[row][col+1].intensite++
		if dumbo[row][col+1].intensite == 9 && !dumbo[row][col+1].highlighted {
			dumbo[row][col+1].highlighted = true
			updateNeighbours(dumbo, row, col+1, width, length)
		}
	}

	if col-1 >= 0 && row-1 >= 0 {
		dumbo[row-1][col-1].intensite++
		if dumbo[row-1][col-1].intensite == 9 && !dumbo[row-1][col-1].highlighted {
			dumbo[row-1][col-1].highlighted = true
			updateNeighbours(dumbo, row-1, col-1, width, length)
		}
	}
	if col+1 < width && row-1 >= 0 {
		dumbo[row-1][col+1].intensite++
		if dumbo[row-1][col+1].intensite == 9 && !dumbo[row-1][col+1].highlighted {
			dumbo[row-1][col+1].highlighted = true
			updateNeighbours(dumbo, row-1, col+1, width, length)
		}
	}
	if col-1 >= 0 && row+1 < length {
		dumbo[row+1][col-1].intensite++
		if dumbo[row+1][col-1].intensite == 9 && !dumbo[row+1][col-1].highlighted {
			dumbo[row+1][col-1].highlighted = true
			updateNeighbours(dumbo, row+1, col-1, width, length)
		}
	}
	if col+1 < width && row+1 < length {
		dumbo[row+1][col+1].intensite++
		if dumbo[row+1][col+1].intensite >= 9 && !dumbo[row+1][col+1].highlighted {
			dumbo[row+1][col+1].highlighted = true
			updateNeighbours(dumbo, row+1, col+1, width, length)
		}
	}
}

func updateOctopus(dumbo [][]Dumbo, width, length int) {
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			dumbo[i][j].intensite++
			if dumbo[i][j].intensite >= 9 && !dumbo[i][j].highlighted {
				dumbo[i][j].highlighted = true
				updateNeighbours(dumbo, i, j, width, length)
			}
		}
	}
}

func resetState(dumbo [][]Dumbo, width, length int) int {
	counter := 0
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			if dumbo[i][j].intensite >= 9 {
				dumbo[i][j].intensite = 0
				dumbo[i][j].highlighted = false
				counter++
			}
		}
	}
	return counter
}

func pb1(lines []string) {
	matrix := initOctopus(lines)
	length := len(matrix)
	width := len(matrix[0])
	counter := 0
	fmt.Println(matrix[0][3])
	for i := 0; i <= 100; i++ {
		updateOctopus(matrix, width, length)
		fmt.Println(matrix[0][3])
		counter += resetState(matrix, width, length)
		fmt.Println(matrix[0][3])
		fmt.Println("-------------", i)
	}
	fmt.Println(counter)

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
