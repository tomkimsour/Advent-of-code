package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tomkimsour/Advent-of-code/convert"
)

// 0 - taille : 6
// 1 - taille : 2
// 2 - taille : 5
// 3 - taille : 5
// 4 - taille : 4
// 5 - taille : 5
// 6 - taille : 6
// 7 - taille : 3
// 8 - taille : 7
// 9 - taille : 6
func pb1(lines []string) {
	var nb int = 0
	for _, line := range lines {
		splitLine := strings.Split(line, " | ")
		res := strings.Split(splitLine[1], " ")
		for _, singleDigit := range res {
			size := len(singleDigit)
			if size == 2 || size == 4 || size == 3 || size == 7 {
				nb++
			}
		}
	}
	fmt.Println(nb)
}

type sample struct {
	zero  []string
	one   []string
	two   []string
	three []string
	four  []string
	five  []string
	six   []string
	seven []string
	eight []string
	nine  []string
}

func initSample(allDigits []string) sample {
	var sp sample
	sp.zero = nil
	sp.one = nil
	sp.two = nil
	sp.three = nil
	sp.four = nil
	sp.five = nil
	sp.six = nil
	sp.seven = nil
	sp.eight = nil
	sp.nine = nil

	for _, singleDigit := range allDigits {
		size := len(singleDigit)
		switch size {
		case 2:
			sp.one = strings.Split(singleDigit, "")
		case 4:
			sp.four = strings.Split(singleDigit, "")
		case 3:
			sp.seven = strings.Split(singleDigit, "")
		case 7:
			sp.eight = strings.Split(singleDigit, "")
		}
	}
	return sp
}

func getLenFive(allDigits []string) [3][]string {
	var res [3][]string

	var i int = 0
	for _, singleDigit := range allDigits {
		size := len(singleDigit)
		switch size {
		case 5:
			res[i] = strings.Split(singleDigit, "")
			i++
		}
	}

	return res
}

func getLenSix(allDigits []string) [3][]string {
	var res [3][]string

	var i int = 0
	for _, singleDigit := range allDigits {
		size := len(singleDigit)
		switch size {
		case 6:
			res[i] = strings.Split(singleDigit, "")
			i++
		}
	}

	return res
}

func initSevenDigit(sp sample, allDigits []string) sample {

	// assign value 2
	fv := getLenFive(allDigits)
	for _, val := range fv {
		var iter int = 0
		for _, v5 := range val {
			for _, v4 := range sp.four {
				if v4 == v5 {
					iter++
				}
			}
		}
		if iter != 3 {
			sp.two = val
		}
	}

	// assign 3 and 5
	for _, val := range fv {
		var iter int = 0
		for _, v5 := range val {
			for _, v2 := range sp.two {
				if v2 == v5 {
					iter++
				}
			}
		}
		if iter == 4 {
			sp.three = val
		} else if iter == 3 {
			sp.five = val
		}
	}

	// 0,6,9 are left
	sx := getLenSix(allDigits)
	var lasts [2][]string
	var lastsIndex int = 0
	// assign 6
	for _, val := range sx {
		var iter int = 0
		for _, v6 := range val {
			for _, v1 := range sp.one {
				if v1 == v6 {
					iter++
				}
			}
		}
		if iter == 1 {
			sp.six = val
		} else {
			lasts[lastsIndex] = val
			lastsIndex++
		}
	}

	//assign 0 and 9
	for _, val := range lasts {
		var iter int = 0
		for _, v6 := range val {
			for _, v4 := range sp.four {
				if v4 == v6 {
					iter++
				}
			}
		}
		if iter == 4 {
			sp.nine = val
		} else {
			sp.zero = val
		}
	}
	return sp
}

func isNumber(digit string, spDigit []string) bool {
	splitDigit := strings.Split(digit, "")
	if len(splitDigit) != len(spDigit) {
		return false
	}
	for _, rd := range splitDigit {
		var isEqual bool = false
		for _, spd := range spDigit {
			if rd == spd {
				isEqual = true
				break
			}
		}
		if !isEqual {
			return false
		}
	}
	return true
}

func pb2(lines []string) {
	// var nb int = 0
	finalSum := 0
	var number string

	for _, line := range lines {
		splitLine := strings.Split(line, " | ")

		// parse the 10 digits
		allDigits := strings.Split(splitLine[0], " ")
		sp := initSample(allDigits)
		sp = initSevenDigit(sp, allDigits)

		// parse the result
		res := strings.Split(splitLine[1], " ")
		number = ""
		for _, singleDigit := range res {
			if isNumber(singleDigit, sp.zero) {
				number += "0"
			} else if isNumber(singleDigit, sp.one) {
				number += "1"
			} else if isNumber(singleDigit, sp.two) {
				number += "2"
			} else if isNumber(singleDigit, sp.three) {
				number += "3"
			} else if isNumber(singleDigit, sp.four) {
				number += "4"
			} else if isNumber(singleDigit, sp.five) {
				number += "5"
			} else if isNumber(singleDigit, sp.six) {
				number += "6"
			} else if isNumber(singleDigit, sp.seven) {
				number += "7"
			} else if isNumber(singleDigit, sp.eight) {
				number += "8"
			} else if isNumber(singleDigit, sp.nine) {
				number += "9"
			} else {
				log.Fatal("Matches no number")
			}
		}
		fmt.Println(res)
		fmt.Println(number)
		finalSum += convert.StringToInt(number)
	}
	fmt.Println(finalSum)
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
