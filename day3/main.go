package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func getCommon(batch [][]byte) int {

	for i := 0; i < len(batch[0]); i++ {
		if foundIn(batch[1], batch[0][i]) && foundIn(batch[2], batch[0][i]) {
			return int(batch[0][i])
		}
	}
	return -1
}

func foundIn(data []byte, b byte) bool {
	for _, i := range data {
		if i == b {
			return true
		}
	}
	return false
}

func Value(a int) int {
	if a <= int('Z') {
		return a - int('A') + 27
	}
	return a - int('a') + 1
}

func main() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("error reading file, %v", err)
	}

	sum := 0
	lines := strings.Split(string(file), "\n")
	for i := 0; i < len(lines)-1; i += 3 {
		batch := [][]byte{[]byte(lines[i]), []byte(lines[i+1]), []byte(lines[i+2])}
		badge := getCommon(batch)
		if badge == -1 {
			continue
		}
		sum += Value(badge)
	}
	fmt.Println(sum)
}
