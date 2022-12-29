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
	EXECUTING bool = false
)

var (
	commLengths = map[string]int{
		"addx": 2,
		"noop": 1,
	}
)

type Input struct {
	command string
	val     int
	ticks   int
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func parseLine(line string) Input {
	arr := strings.Split(line, " ")
	if len(arr) == 1 {
		return Input{
			command: arr[0],
			ticks:   1,
		}
	}

	n, err := strconv.Atoi(arr[1])
	handleErr(err)

	return Input{
		command: arr[0],
		val:     n,
		ticks:   2,
	}
}

func contains(target int, slice [3]int) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

func isSprite(val, reg int) bool {
	diff := reg - val
	if math.Abs(float64(diff)) <= 1 {
		return true
	}
	return false
}

func main() {

	// file, err := os.Open("sample.txt")
	file, err := os.Open("input.txt")

	handleErr(err)

	defer file.Close()

	ls := bufio.NewScanner(file)

	cycle := 0
	reg := 1

	for ls.Scan() {

		text := ls.Text()

		input := parseLine(text)

		for input.ticks > 0 {

			if cycle%40 == 0 && cycle > 0 {
				fmt.Printf("\n")
				reg += 40
			}

			if isSprite(cycle, reg) {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}

			cycle++
			input.ticks--
		}

		reg += input.val

	}

	fmt.Printf("\n")

}
