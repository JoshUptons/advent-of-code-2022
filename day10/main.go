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
	EXECUTING bool = false
)

var (
	commLengths = map[string]int{
		"addx": 2,
		"noop": 1,
	}
	interestingCycles = map[int]bool{
		20:  true,
		60:  true,
		100: true,
		140: true,
		180: true,
		220: true,
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

func main() {

	// file, err := os.Open("sample.txt")
	file, err := os.Open("input.txt")
	handleErr(err)

	defer file.Close()

	ls := bufio.NewScanner(file)

	inputs := []Input{}

	for ls.Scan() {

		text := ls.Text()

		inputs = append(inputs, parseLine(text))

	}

	cycle := 0
	str := 1
	totalSignalStr := 0

	for _, input := range inputs {

		for input.ticks > 0 {
			cycle++
			if interestingCycles[cycle] {
				fmt.Println(cycle, str, cycle*str)
				totalSignalStr += cycle * str
			}
			input.ticks--

		}
		str += input.val

	}

	fmt.Println(totalSignalStr)
}
