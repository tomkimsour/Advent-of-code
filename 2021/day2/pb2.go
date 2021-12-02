package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func pb1(lines []string) {
	var depth int = 0
	var forward int = 0
	for _, line := range lines {
		fields := strings.Split(line, " ")
		switch fields[0] {
		case "forward":
			steps, _ := strconv.Atoi(fields[1])
			forward = forward + steps
		case "down":
			steps, _ := strconv.Atoi(fields[1])
			depth = depth + steps
		case "up":
			steps, _ := strconv.Atoi(fields[1])
			depth = depth - steps
		default:
			fmt.Println("the input file is badly formated")
		}
	}
	fmt.Println(depth * forward)
}

func pb2(lines []string) {
	var depth int = 0
	var aim int = 0
	var forward int = 0
	for _, line := range lines {
		fields := strings.Split(line, " ")
		switch fields[0] {
		case "forward":
			steps, _ := strconv.Atoi(fields[1])
			forward = forward + steps
			depth = depth + aim*steps
		case "down":
			steps, _ := strconv.Atoi(fields[1])
			aim = aim + steps
		case "up":
			steps, _ := strconv.Atoi(fields[1])
			aim = aim - steps
		default:
			fmt.Println("the input file is badly formated")
		}
	}
	fmt.Println(depth * forward)

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
