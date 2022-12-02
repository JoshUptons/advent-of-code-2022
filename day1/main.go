package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	highCalories := 0

	currentElf := 0

	for _, line := range strings.Split(string(file), "\n") {
		if line == "" || line == "\n" {
			if currentElf > highCalories {
				highCalories = currentElf
			}
			currentElf = 0
		}
		cals, _ := strconv.Atoi(line)
		currentElf += cals
	}

	log.Println(highCalories)

}
