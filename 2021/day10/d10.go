package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"

	stack "github.com/golang-collections/collections/stack"
)

type balises struct {
	nbParentheses int
	nbCrochets    int
	nbAccolades   int
	nbChevron     int
}

func pb1(lines []string) {
	var b balises
	b.nbParentheses = 0
	b.nbCrochets = 0
	b.nbAccolades = 0
	b.nbChevron = 0

	for _, line := range lines {
		leftStr := stack.New()
		rightStr := stack.New()
		isFinished := false
		for _, char := range line {
			switch char {
			case '(':
				leftStr.Push('(')
				rightStr.Push(')')
			case ')':
				if rightStr.Peek() != ')' {
					b.nbParentheses++
					isFinished = true
				} else {
					rightStr.Pop()
				}
			case '{':
				leftStr.Push('{')
				rightStr.Push('}')
			case '}':
				if rightStr.Peek() != '}' {
					b.nbAccolades++
					isFinished = true
				} else {
					rightStr.Pop()
				}
			case '[':
				leftStr.Push('[')
				rightStr.Push(']')
			case ']':
				if rightStr.Peek() != ']' {
					b.nbCrochets++
					isFinished = true
				} else {
					rightStr.Pop()
				}
			case '<':
				leftStr.Push('<')
				rightStr.Push('>')
			case '>':
				if rightStr.Peek() != '>' {
					b.nbChevron++
					isFinished = true
				} else {
					rightStr.Pop()
				}
			}
			if isFinished {
				break
			}
		}
	}
	fmt.Println(b.nbAccolades*1197 + b.nbChevron*25137 + b.nbCrochets*57 + b.nbParentheses*3)
}

func pb2(lines []string) {
	var scores []int

	for _, line := range lines {
		leftStr := stack.New()
		rightStr := stack.New()
		isFinished := false

		for _, char := range line {
			switch char {
			case '(':
				leftStr.Push('(')
				rightStr.Push(')')
			case ')':
				if rightStr.Peek() != ')' {
					isFinished = true
				} else {
					rightStr.Pop()
				}
			case '{':
				leftStr.Push('{')
				rightStr.Push('}')
			case '}':
				if rightStr.Peek() != '}' {
					isFinished = true
				} else {
					rightStr.Pop()
				}
			case '[':
				leftStr.Push('[')
				rightStr.Push(']')
			case ']':
				if rightStr.Peek() != ']' {
					isFinished = true
				} else {
					rightStr.Pop()
				}
			case '<':
				leftStr.Push('<')
				rightStr.Push('>')
			case '>':
				if rightStr.Peek() != '>' {
					isFinished = true
				} else {
					rightStr.Pop()
				}
			}
			if isFinished {
				break
			}
		}
		if !isFinished {
			score := 0
			for i := rightStr.Peek(); rightStr.Len() > 0; i = rightStr.Pop() {
				switch i {
				case ')':
					score = score*5 + 1
				case ']':
					score = score*5 + 2
				case '}':
					score = score*5 + 3
				case '>':
					score = score*5 + 4
				}
			}
			scores = append(scores, score)
		}
	}
	size := len(scores)
	sort.Ints(scores)
	for i, e := range scores {
		fmt.Println(i, ": ", e)
	}
	fmt.Println(scores[(size/2)-1])
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
