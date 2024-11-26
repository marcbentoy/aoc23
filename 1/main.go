package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	puzzleFile = "puzzle.txt"
	numbers    = "0123456789"
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
	i := 0
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
		i++
	}

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

	for _, s := range line {
		if isNum(s) {
			return string(s)
		}
	}

	return string(0)
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
