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

type Card struct {
	wNums   []int
	hNums   []int
	qty     int
	matches int
}

// algorithm:
// 1. get all cards
// 2. evalulate card points

func main() {
	fmt.Println("advent of code 2023 day 4p2")

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	r := bufio.NewReader(file)

	var points []int
	var cards []*Card
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}

		card := getCard(string(line))
		cards = append(cards, &card)
		evalPoints(&card)

		wNums, hNums := getNums(string(line))
		point := calcPoints(wNums, hNums)
		if point == 0 {
			continue
		}
		points = append(points, point)
	}

	evalQty(cards)

	fmt.Println("after evaluating qty:")

	fmt.Println("sum qty: ", sumQty(cards))
}

func sum(points []int) (total int) {
	for _, p := range points {
		total += p
	}
	return
}

func sumQty(cards []*Card) (totalQty int) {
	for _, c := range cards {
		totalQty += c.qty
	}
	return
}

func evalQty(cards []*Card) {
	i := 0
	for ; i < len(cards); i++ {
		j := cards[i].matches
		for ; j > 0; j-- {
			if i+j >= len(cards) {
				continue
			}
			cards[i+j].qty += cards[i].qty
		}
	}
}

func getCard(line string) (card Card) {
	mainSplit := strings.Split(line, ":")
	rawNums := strings.Split(mainSplit[1], "|")
	rawWNums := strings.TrimSpace(rawNums[0])
	rawHNums := strings.TrimSpace(rawNums[1])
	strWNums := strings.Split(rawWNums, " ")
	strHNums := strings.Split(rawHNums, " ")

	var wNums, hNums []int

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

	card.wNums = wNums
	card.hNums = hNums
	card.qty = 1
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

func evalPoints(card *Card) {
	for _, w := range card.wNums {
		if contains(w, card.hNums) {
			card.matches++
		}
	}
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
