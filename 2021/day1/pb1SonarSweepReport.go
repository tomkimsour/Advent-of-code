package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func maxIncreasingInARow(scanner bufio.Scanner) int {
	var previous int = 0
	var increase int = 0
	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if depth > previous {
			increase++
		}
		previous = depth
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return increase - 1
}

func batchIncreament(file1, file2, file3 *os.File) int {
	scanner1 := bufio.NewScanner(file1)

	scanner2 := bufio.NewScanner(file2)
	scanner2.Scan()
	if err := scanner2.Err(); err != nil {
		log.Fatal(err)
	}

	scanner3 := bufio.NewScanner(file3)
	scanner3.Scan()
	if err := scanner3.Err(); err != nil {
		log.Fatal(err)
	}
	scanner3.Scan()
	if err := scanner3.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(scanner1.Text())
	fmt.Println(scanner2.Text())
	fmt.Println(scanner3.Text())
	var increase int = 0
	var previous int = 0

	for scanner3.Scan() && scanner1.Scan() && scanner2.Scan() {

		depth1, err := strconv.Atoi(scanner1.Text())
		if err != nil {
			log.Fatal(err)
		}
		depth2, err := strconv.Atoi(scanner2.Text())
		if err != nil {
			log.Fatal(err)
		}
		depth3, err := strconv.Atoi(scanner3.Text())
		if err != nil {
			log.Fatal(err)
		}
		sum := depth1 + depth2 + depth3
		if (sum) > previous {
			increase++
		}
		previous = sum
	}
	return increase - 1
}

func main() {
	file, err := os.Open("input/ss_report.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	max := maxIncreasingInARow(*scanner)
	fmt.Println(max)

	file1, _ := os.Open("input/ss_report.txt")
	defer file1.Close()
	file2, _ := os.Open("input/ss_report.txt")
	defer file2.Close()
	file3, _ := os.Open("input/ss_report.txt")
	defer file3.Close()
	max = batchIncreament(file1, file2, file3)
	fmt.Println(max)
}
