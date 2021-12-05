package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func initChartSet(lines []string, nbChart int) [][][]int {

	setCharts := make([][][]int, nbChart)
	for i := 0; i < nbChart; i++ {
		setCharts[i] = make([][]int, 5)
		for j := 0; j < 5; j++ {
			setCharts[i][j] = make([]int, 5)
		}
	}

	for i := 0; i < nbChart; i++ {
		for j := 0; j < 5; j++ {
			line := regexp.MustCompile(`[0-9]+`).FindAllString(lines[i*6+j], 5)
			for k := 0; k < 5; k++ {
				value, _ := strconv.Atoi(line[k])
				setCharts[i][j][k] = value
			}
		}
	}

	return setCharts
}

func initCheckBoard(nbCharts int) [][][]bool {
	checkBoard := make([][][]bool, nbCharts)
	for k := 0; k < nbCharts; k++ {
		checkBoard[k] = make([][]bool, 5)
		for i := 0; i < 5; i++ {
			checkBoard[k][i] = make([]bool, 5)
		}
	}

	for k := 0; k < nbCharts; k++ {
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				checkBoard[k][i][j] = false
			}
		}
	}

	return checkBoard
}

func isWin(board [][]bool) bool {
	// check lines
	for line := 0; line < 5; line++ {
		var win bool = true
		for j := 0; j < 5; j++ {
			win = board[line][j] && win
		}
		if win {
			return true
		}
	}

	// check col
	for col := 0; col < 5; col++ {
		var win bool = true
		for row := 0; row < 5; row++ {
			win = board[row][col] && win
		}
		if win {
			return true
		}
	}

	return false
}

func playMove(checkBoard [][]bool, chart [][]int, valueInt int) [][]bool {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if chart[i][j] == valueInt {
				checkBoard[i][j] = true
			}
		}
	}
	return checkBoard
}

func getMarkedUnmarked(winningBoard [][]int, winningCheckBoard [][]bool) (int, int) {
	var sumMarked int = 0
	var sumUnmarked int = 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if winningCheckBoard[i][j] {
				sumMarked = sumMarked + winningBoard[i][j]
			} else {
				sumUnmarked = sumUnmarked + winningBoard[i][j]
			}
		}
	}
	return sumMarked, sumUnmarked
}

func pb1(lines []string) {
	draw := strings.Split(lines[0], ",")
	fmt.Println(draw)
	lines = lines[2:]
	nbChart := len(lines)/6 + 1

	chartSet := initChartSet(lines, nbChart)
	checkBoard := initCheckBoard(nbChart)

	var isWon bool = false
	var winningBoard [][]int
	var winningCheckBoard [][]bool
	for value := 0; value < len(draw) && !isWon; value++ {
		valueInt, _ := strconv.Atoi(draw[value])
		for i := 0; i < nbChart; i++ {
			checkBoard[i] = playMove(checkBoard[i], chartSet[i], valueInt)
			if value < 5 {
				continue
			}
			isWon = isWin(checkBoard[i])
			if isWon {
				fmt.Printf("someone won at iter : %d\n", value)
				winningBoard = chartSet[i]
				winningCheckBoard = checkBoard[i]
				break
			}
		}
	}

	fmt.Println(winningBoard)

	marked, unmarked := getMarkedUnmarked(winningBoard, winningCheckBoard)
	fmt.Println(marked * unmarked)
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
