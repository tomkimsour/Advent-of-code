package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func pb1(lines []string) {
	var gamma string = ""
	var epsilon string = ""

	bitSize := len(lines[0])
	bitResultArray := make([]int, bitSize)
	for i := 0; i < bitSize; i++ {
		bitResultArray[i] = 0
	}

	for _, line := range lines {
		for i, bit := range line {
			value, _ := strconv.Atoi(string(bit))
			if value == 1 {
				bitResultArray[i]++
			} else {
				bitResultArray[i]--
			}
		}
	}

	for i := 0; i < bitSize; i++ {
		if bitResultArray[i] > 0 {
			gamma = gamma + "1"
			epsilon = epsilon + "0"
		} else {
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		}
	}
	gam, _ := strconv.ParseInt(gamma, 2, 64)
	eps, _ := strconv.ParseInt(epsilon, 2, 64)

	fmt.Println(gam * eps)
}

func getOxygen(lines []string) int64 {
	bitSize := len(lines[0])

	for i := 0; i < bitSize && len(lines) > 1; i++ {
		var bitpos []string
		var bitneg []string

		for _, line := range lines {
			value, _ := strconv.Atoi(string(line[i]))
			if value == 1 {
				bitpos = append(bitpos, line)
			} else {
				bitneg = append(bitneg, line)
			}
		}

		if len(bitpos) >= len(bitneg) {
			lines = bitpos
		} else {
			lines = bitneg
		}
	}
	var oxygen int64
	oxygen, _ = strconv.ParseInt(lines[0], 2, 64)
	return oxygen
}

func getCo2(lines []string) int64 {
	bitSize := len(lines[0])

	for i := 0; i < bitSize && len(lines) > 1; i++ {
		var bitpos []string
		var bitneg []string

		for _, line := range lines {
			value, _ := strconv.Atoi(string(line[i]))
			if value == 1 {
				bitpos = append(bitpos, line)
			} else {
				bitneg = append(bitneg, line)
			}
		}

		if len(bitpos) < len(bitneg) {
			lines = bitpos
		} else {
			lines = bitneg
		}
	}

	var co2 int64
	co2, _ = strconv.ParseInt(lines[0], 2, 64)
	return co2
}

func pb2(lines []string) {

	oxygen := getOxygen(lines)
	co2 := getCo2(lines)
	fmt.Println(oxygen * co2)
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
