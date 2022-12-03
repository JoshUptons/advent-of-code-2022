package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Round struct {
	Moves          []string
	DesiredOutcome int
}

func (r *Round) calcValue() int {
	return outcomes[r.Moves[1]] + mappings[r.Moves[0]][r.DesiredOutcome]
}

type Move struct {
	Text string
	Win  string
	Draw string
	Lose string
}

const (
	win  int = 6
	draw     = 3
	loss     = 0
)

var (
	mappings = map[string]map[int]int{
		"A": {
			win:  values["Y"],
			draw: values["X"],
			loss: values["Z"],
		},
		"B": {
			win:  values["Z"],
			draw: values["Y"],
			loss: values["X"],
		},
		"C": {
			win:  values["X"],
			draw: values["Z"],
			loss: values["Y"],
		},
	}
	values = map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	outcomes = map[string]int{
		"X": loss,
		"Y": draw,
		"Z": win,
	}
)

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	totalPoints := 0
	for fileScanner.Scan() {
		r := Round{}
		r.Moves = strings.Split(fileScanner.Text(), " ")
		r.DesiredOutcome = outcomes[r.Moves[1]]
		totalPoints += r.calcValue()
	}

	fmt.Println(totalPoints)

}
