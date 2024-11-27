package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Digit struct {
	value string
	line  int
	index int
}

type Symbol struct {
	value string
	line  int
	index int
}

const (
	filename = "puzzle"
	digits   = "0123456789"
)

func main() {
	fmt.Println("advent of code 2023 day 3")

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	r := bufio.NewReader(file)

	// partNumbers are now the product of two numbers in a gear
	var digits []*Digit
	var gears, symbols []*Symbol
	var partNumbers []int
	i := 0
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}

		linesymbols, linedigits := getTokens(string(line), i)
		symbols = append(symbols, linesymbols...)
		digits = append(digits, linedigits...)
		i++
	}

	gears = findGears(symbols, digits)

	partNumbers = eval(gears, digits)

	fmt.Println("sum: ", sum(partNumbers))
}

func sum(digits []int) (total int) {
	for _, d := range digits {
		total += d
	}
	return
}

// gets all the tokens (symbol or number) in the line
func getTokens(line string, lineNum int) (symbols []*Symbol, digits []*Digit) {
	number := ""
	i := 0
	for ; i < len(line); i++ {
		if line[i] == '.' {
			if number != "" {
				digits = append(digits, &Digit{
					value: number,
					line:  lineNum,
					index: i - 1,
				})
				number = ""
			}
			continue
		}

		// assumes that c is a symbol
		if line[i] == '*' {
			if number != "" {
				digits = append(digits, &Digit{
					value: number,
					line:  lineNum,
					index: i - 1,
				})
				number = ""
			}
			symbols = append(symbols, &Symbol{
				value: string(line[i]),
				line:  lineNum,
				index: i,
			})
			continue
		}

		// if digit
		number += string(line[i])
	}
	if number != "" {
		digits = append(digits, &Digit{
			value: number,
			line:  lineNum,
			index: i - 1,
		})
	}

	return
}

func isDigit(c rune) bool {
	for _, d := range digits {
		if c == d {
			return true
		}
	}
	return false
}

func eval(symbols []*Symbol, digits []*Digit) (partNumbers []int) {
	for _, s := range symbols {
		var product int
		// get adjacent digits
		for _, d := range digits {
			if isAdjacent(s, d) {
				if product == 0 {
					product = 1
				}

				value, err := strconv.Atoi(d.value)
				if err != nil {
					log.Println("error converting value to int:", err)
					return []int{}
				}

				product *= value
			}
		}
		partNumbers = append(partNumbers, product)
	}

	return partNumbers
}

func findGears(symbols []*Symbol, digits []*Digit) (gears []*Symbol) {
	for _, s := range symbols {
		dQty := 0
		// get adjacent digits
		for _, d := range digits {
			if isAdjacent(s, d) {
				dQty++
			}
		}

		if dQty == 2 {
			gears = append(gears, s)
		}
	}
	return
}

func isAdjacent(s *Symbol, d *Digit) bool {
	a := d.index - len(d.value)
	b := d.index + 1
	if (s.line == d.line || (s.line-1 == d.line) || (s.line+1 == d.line)) &&
		(s.index >= a && s.index <= b) {
		return true
	}
	return false
}

func isGear() bool {
	return false
}
