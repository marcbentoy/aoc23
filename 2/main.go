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
	filename = "puzzle.txt"
	reds     = 12
	greens   = 13
	blues    = 14
)

type Set struct {
	Green int
	Blue  int
	Red   int
}

type MaxBag struct {
	Id int
	Bag
}

type Bag struct {
	Red   int
	Green int
	Blue  int
}

func main() {
	// open file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	r := bufio.NewReader(file)

	var powers []int
	for {
		line, _, err := r.ReadLine()

		if err != nil {
			break
		}

		maxBag := *getMaxBag(string(line))
		power := maxBag.Blue * maxBag.Green * maxBag.Red
		powers = append(powers, power)
	}

	fmt.Println("advent of code day 2")
	fmt.Println("sum: ", sum(powers))
}

func sum(ints []int) (sum int) {
	for _, i := range ints {
		sum += i
	}
	return
}

func getMaxBag(line string) *MaxBag {
	// separate game segments
	idpart, setspart, found := strings.Cut(line, ":")
	if !found {
		log.Println("can't find the ':' char for cutting the game")
		return nil
	}

	// get id
	idchar := strings.TrimPrefix(idpart, "Game ")
	id, err := strconv.Atoi(idchar)
	if err != nil {
		log.Println("error converting game id to int:", err)
		return nil
	}

	// get gameSets
	gameSets := strings.Split(setspart, "; ")
	var parsedSets []*Set
	for _, s := range gameSets {
		// split cubes
		parsedSets = append(parsedSets, getCubes(s))
	}

	maxBag := &MaxBag{
		Id: id,
	}
	// get max number of cubes per set
	for _, s := range parsedSets {
		if s.Blue > maxBag.Blue {
			maxBag.Blue = s.Blue
		}
		if s.Green > maxBag.Green {
			maxBag.Green = s.Green
		}
		if s.Red > maxBag.Red {
			maxBag.Red = s.Red
		}
	}

	return maxBag
}

func getCubes(set string) *Set {
	set = strings.TrimSpace(set)
	cubes := strings.Split(set, ",")

	var gs, bs, rs int
	for _, c := range cubes {
		c = strings.TrimSpace(c)
		val := strings.Split(c, " ")

		qty, err := strconv.Atoi(val[0])
		if err != nil {
			log.Println("error converting qty to int:", err)
			return nil
		}

		switch val[1] {
		case "green":
			gs = qty
		case "blue":
			bs = qty
		case "red":
			rs = qty
		}
	}

	return &Set{
		Blue:  bs,
		Green: gs,
		Red:   rs,
	}
}

func isValid(bag MaxBag) bool {
	if bag.Blue > blues || bag.Red > reds || bag.Green > greens {
		return false
	}
	return true
}
