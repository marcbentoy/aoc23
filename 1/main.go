package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	puzzleFile = "puzzle.txt"
	numbers    = "0123456789"
)

var (
	wordNums = []string{
		"one", "two", "three", "four", "five",
		"six", "seven", "eight", "nine",
	}
)

func main() {
	fmt.Println("advent of code 2023")

	var nums []int

	// open file
	file, err := os.Open(puzzleFile)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	r := bufio.NewReader(file)

	// read lines
	for {
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			first := findnum(string(line), false)
			second := findnum(string(line), true)
			num := first + second
			number, err := strconv.Atoi(string(num))
			if err != nil {
				log.Fatal("error converting string num to int: ", err)
				return
			}

			nums = append(nums, number)
		}
		if err != nil {
			break
		}
	}

	fmt.Printf("len(nums): %v\n", len(nums))

	// get sum
	fmt.Printf("sum: %d", sum(nums))
}

// finds and converst string num to int
// toLeft param is used to reverse the finding of the number
func findnum(line string, toLeft bool) string {
	if toLeft {
		// reverse the string
		line = Reverse(line)
	}

	var curr string
	for _, s := range line {
		if isNum(s) {
			return string(s)
		}

		curr += string(s)

		if toLeft {
			curr = Reverse(curr)
		}

		wordNum, yes := isWordNum(curr)
		if yes {
			return convertWordNumToDigit(wordNum)
		}

		if toLeft {
			curr = Reverse(curr)
		}
	}

	return string(0)
}

func FindNums(line string) (foundNums []string) {
	// find
	var curr string
	for _, c := range line {
		curr += string(c)
		// if c is digit
		if isNum(c) {
			foundNums = append(foundNums, string(c))
			curr = ""
			continue
		}

		wordNum, yes := isWordNum(curr)
		if yes {
			foundNums = append(foundNums, convertWordNumToDigit(wordNum))
			curr = ""
			continue
		}
	}
	return
}

func isWordNum(num string) (string, bool) {
	for _, wordNum := range wordNums {
		if strings.Contains(num, wordNum) {
			return wordNum, true
		}
	}
	return "", false
}

func convertWordNumToDigit(wordNum string) string {
	switch wordNum {
	case "zero":
		return "0"
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return ""
	}
}

func isNum(c rune) bool {
	for _, n := range numbers {
		if c == n {
			return true
		}
	}

	return false
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
