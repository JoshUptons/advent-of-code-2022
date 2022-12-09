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

	for i := 0; i < len(stack); i++ {
		for j := 0; j < len(stack); j++ {
			if i == j {
				continue
			}
			if stack[i] == stack[j] {
				return false
			}
		}
	}

	return true
}

func main() {

	file, err := os.Open("input.txt")
	handleErr(err)

	defer file.Close()

	b, err := ioutil.ReadAll(file)
	handleErr(err)

	stack := make([]byte, 4)

	for i := 0; i < len(b); i++ {
		stack[0] = b[i]
		stack[1] = b[i+1]
		stack[2] = b[i+2]
		stack[3] = b[i+3]
		if uniqueBytes(stack) {
			fmt.Println(stack)
			fmt.Printf("%d characters before finding a unique sequence", i+4)
			break
		}
	}

}
