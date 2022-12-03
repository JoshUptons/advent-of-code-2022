package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	win  int = 6
	draw     = 3
	loss     = 0
)

var (
	values = map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	outcomes = map[string]int{
		"AX": draw,
		"AY": win,
		"AZ": loss,
		"BX": loss,
		"BY": draw,
		"BZ": win,
		"CX": win,
		"CY": loss,
		"CZ": draw,
	}
)

func calcOutcome(x, y string) int {
	key := strings.Join([]string{x, y}, "")
	outcome := outcomes[key]
	return outcome
}

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	points := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		moves := strings.Split(line, " ")
		points += calcOutcome(moves[0], moves[1])
		points += values[moves[1]]
	}

	fmt.Println(points)
}
