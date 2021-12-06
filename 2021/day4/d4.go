package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/tomkimsour/Advent-of-code/convert"
)

type board struct {
	grid      [][]int
	checkGrid [][]bool
	unmarked  int
	lastCall  int
}

func initBoards(lines []string, nbChart int) []board {
	boardList := make([]board, nbChart)

	bingoGrid := make([][]int, 5)
	for i := 0; i < 5; i++ {
		bingoGrid[i] = make([]int, 5)
	}

	checkGrid := make([][]bool, 5)
	for i := 0; i < 5; i++ {
		checkGrid[i] = make([]bool, 5)
	}

	for i := 0; i < nbChart; i++ {
		sum := 0
		for j := 0; j < 5; j++ {
			line := regexp.MustCompile(`[0-9]+`).FindAllString(lines[i*6+j], 5)
			for k := 0; k < 5; k++ {
				value, _ := strconv.Atoi(line[k])
				sum = sum + value
				bingoGrid[j][k] = value
				checkGrid[j][k] = false
			}
		}
		boardList[i].grid = bingoGrid
		boardList[i].checkGrid = checkGrid
		boardList[i].unmarked = sum
	}

	return boardList
}

func playMove(bingoBoard board, valueInt int) board {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if bingoBoard.grid[i][j] == valueInt {
				bingoBoard.checkGrid[i][j] = true
				bingoBoard.unmarked = bingoBoard.unmarked - valueInt
				bingoBoard.lastCall = valueInt
			}
		}
	}

	return bingoBoard
}

func isGameWon(bingoBoard board) bool {
	// check lines
	for row := 0; row < 5; row++ {
		var win bool = true
		for col := 0; col < 5; col++ {
			win = bingoBoard.checkGrid[row][col] && win
		}
		if win {
			return true
		}
	}

	// check col
	for col := 0; col < 5; col++ {
		var win bool = true
		for row := 0; row < 5; row++ {
			win = bingoBoard.checkGrid[row][col] && win
		}
		if win {
			return true
		}
	}
	return false
}

func pb1(lines []string) {
	draw := strings.Split(lines[0], ",")
	fmt.Println(draw)
	lines = lines[2:]
	nbChart := len(lines)/6 + 1

	boardList := initBoards(lines, nbChart)

	var isWon bool = false
	var winningBoard board

	for value := 0; value < len(draw) && !isWon; value++ {
		valueInt := convert.StringToInt(draw[value])
		for i := 0; i < nbChart; i++ {
			boardList[i] = playMove(boardList[i], valueInt)
			if value < 5 {
				continue
			}
			isWon = isGameWon(boardList[i])
			if isWon {
				fmt.Printf("someone won at iter : %d\n", value)
				winningBoard = boardList[i]
				break
			}
		}
	}

	fmt.Println(winningBoard.grid)
	fmt.Println(winningBoard.lastCall * winningBoard.unmarked)

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
