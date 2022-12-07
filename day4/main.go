package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	lineScanner := bufio.NewScanner(file)
	sum := 0

	for lineScanner.Scan() {
		line := lineScanner.Text()

		elves := [][]int{}

		for _, elf := range strings.Split(line, ",") {
			sections := []int{}
			for _, s := range strings.Split(elf, "-") {
				n, err := strconv.Atoi(s)
				handleErr(err)
				sections = append(sections, n)
			}
			elves = append(elves, sections)
		}

		overlap := false
		for i, n := range elves[0] {
			if n >= elves[1][0] && n <= elves[1][1] {
				overlap = true
			}
			if elves[1][i] >= elves[0][0] && elves[1][i] <= elves[0][i] {
				overlap = true
			}
		}
		if overlap {
			sum++
		}

	}
	fmt.Println(sum, "overlaps")
}
