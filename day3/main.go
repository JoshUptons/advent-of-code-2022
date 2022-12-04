package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func contains(a byte, b []byte) bool {
	for _, n := range b {
		if n == a {
			return true
		}
	}
	return false
}

func Value(a int) int {
	if a <= int(rune('Z')) {
		return a - int('A') + 27
	}
	return a - int('a') + 1
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	sum := 0

	for fileScanner.Scan() {
		line := fileScanner.Bytes()
		left := line[0 : len(line)/2]
		right := line[len(line)/2:]
		for i := 0; i < len(left); i++ {
			if contains(left[i], right) {
				sum += Value(int(left[i]))
				break
			}
		}
	}

	fmt.Println(sum)
}
