package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	filename = "puzzle"
)

func main() {
	fmt.Println("advent of code 2023 day 4")

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	r := bufio.NewReader(file)

	var points []int
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}

		wNums, hNums := getNums(string(line))
		point := calcPoints(wNums, hNums)
		if point == 0 {
			continue
		}
		points = append(points, point)
	}
	for _, p := range points {
		fmt.Printf("p: %v\n", p)
	}

	fmt.Println("sum: ", sum(points))
}

func sum(points []int) (total int) {
	for _, p := range points {
		total += p
	}
	return
}

func getNums(line string) (wNums, hNums []int) {
	mainSplit := strings.Split(line, ":")
	rawNums := strings.Split(mainSplit[1], "|")
	rawWNums := strings.TrimSpace(rawNums[0])
	rawHNums := strings.TrimSpace(rawNums[1])
	strWNums := strings.Split(rawWNums, " ")
	strHNums := strings.Split(rawHNums, " ")

	// convert wnums to int
	for _, w := range strWNums {
		if w == "" {
			continue
		}

		value, err := strconv.Atoi(w)
		if err != nil {
			log.Println("error converting w to int:", err)
			return
		}
		wNums = append(wNums, value)
	}

	// convert hnums to int
	for _, h := range strHNums {
		if h == "" {
			continue
		}

		value, err := strconv.Atoi(h)
		if err != nil {
			log.Println("error converting h to int:", err)
			return
		}
		hNums = append(hNums, value)
	}
	return
}

func calcPoints(wNums, hNums []int) (points int) {
	for _, w := range wNums {
		if contains(w, hNums) {
			points++
		}
	}
	num := math.Exp2(float64(points) - 1)
	return int(num)
}

func contains(w int, hNums []int) bool {
	for _, h := range hNums {
		if h == w {
			return true
		}
	}
	return false
}
