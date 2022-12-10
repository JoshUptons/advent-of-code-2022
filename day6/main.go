package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func uniqueBytes(stack []byte) bool {

	m := make(map[byte]bool)

	for _, b := range stack {
		if _, ok := m[b]; !ok {
			m[b] = true
		} else {
			return false
		}
	}

	return true
}

type stack []byte

func main() {

	file, err := os.Open("input.txt")
	handleErr(err)

	defer file.Close()

	b, err := ioutil.ReadAll(file)
	handleErr(err)

	for i := 0; i < len(b)-13; i++ {
		start := i
		end := i + 14
		if uniqueBytes(b[start:end]) {
			fmt.Printf("found unique sequence after %d characters read\n", i+14)
			break
		}
	}
}
