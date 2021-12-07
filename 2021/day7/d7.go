package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"sync"

	"github.com/tomkimsour/Advent-of-code/convert"
)

func getMax(values []int) int {
	max := 0
	for _, val := range values {
		if val > max {
			max = val
		}
	}
	return max
}
func getMin(values []int) int {
	min := 999999999
	for _, val := range values {
		if val < min {
			min = val
		}
	}
	return min
}

func getGasCost(crabs []int, iter int) int {
	cost := 0
	for i := 0; i < len(crabs); i++ {
		cost += int(math.Abs(float64(crabs[i] - iter)))
	}
	return cost
}

func getRealGasCost(crabs []int, iter int) int {
	cost := 0
	for i := 0; i < len(crabs); i++ {
		dist := int(math.Abs(float64(crabs[i] - iter)))
		for j := 1; j < dist+1; j++ {
			cost += j
		}
	}
	return cost
}

func pb1(lines []string) {
	var wg sync.WaitGroup
	crabs := convert.ArrStrToArrInt(strings.Split(lines[0], ","))

	maxDistance := getMax(crabs)
	minDistance := getMin(crabs)
	nbPossibility := maxDistance - minDistance

	gasResults := make([]int, nbPossibility)
	for i := 0; i < nbPossibility; i++ {
		wg.Add(1)
		go func(i int) {
			gasResults[i] = getGasCost(crabs, i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(getMin(gasResults))

}

func pb2(lines []string) {
	var wg sync.WaitGroup
	crabs := convert.ArrStrToArrInt(strings.Split(lines[0], ","))

	maxDistance := getMax(crabs)
	minDistance := getMin(crabs)
	nbPossibility := maxDistance - minDistance

	gasResults := make([]int, nbPossibility)
	for i := 0; i < nbPossibility; i++ {
		wg.Add(1)
		go func(i int) {
			gasResults[i] = getRealGasCost(crabs, i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(getMin(gasResults))

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
