package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tomkimsour/Advent-of-code/convert"
)

func arrToLinkedList(in []int) *list.List {
	l := list.New()
	for _, i := range in {
		l.PushBack(i)
	}
	return l
}

func pb1(in []int) {
	fishList := arrToLinkedList(in)
	fmt.Println(fishList.Len())

	for i := 0; i < 80; i++ {
		for e := fishList.Front(); e != nil; e = e.Next() {
			val := e.Value.(int)
			if val == 0 {
				e.Value = 6
				fishList.PushFront(8)
			} else {
				e.Value = val - 1
			}
		}
	}
	fmt.Println(fishList.Len())
}

func pb2(in []int) {
	var fishList [9]int
	for _, val := range in {
		fishList[val]++
	}

	for i := 0; i < 256; i++ {
		fertile := fishList[0]
		for j := 0; j < 8; j++ {
			fishList[j] = fishList[j+1]
		}
		fishList[8] = fertile
		fishList[6] = fishList[6] + fertile
	}
	sum := 0
	for _, val := range fishList {
		sum = sum + val
	}
	fmt.Println(sum)

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
	line := strings.Split(lines[0], ",")
	input := convert.ArrStrToArrInt(line)

	pb1(input)
	pb2(input)
}
