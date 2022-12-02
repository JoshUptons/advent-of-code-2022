package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sortHeap(heap []int) {
	sort.Slice(heap, func(i, j int) bool {
		return heap[i] > heap[j]
	})
}

func main() {

	file, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	highestCals := make([]int, 3)

	currentElf := 0

	for _, line := range strings.Split(string(file), "\n") {
		if line == "" || line == "\n" {
			if currentElf > highestCals[len(highestCals)-1] {
				highestCals[2] = currentElf
				sortHeap(highestCals)
			}
			currentElf = 0
		}
		cals, _ := strconv.Atoi(line)
		currentElf += cals
	}

	totalCals := 0
	for _, cals := range highestCals {
		totalCals += cals
	}

	log.Println(totalCals)

}
