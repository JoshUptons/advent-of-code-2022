package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func overlap(a []int, b []int) bool {
	if (a[0] <= b[0] && a[1] >= b[1]) || (b[0] <= a[0] && b[1] >= a[1]) {
		fmt.Println(a, b)
		return true
	}
	return false
}

func handleErr(err error) {
	log.Fatal(err)
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

		arr1 := strings.Split(strings.Split(line, ",")[0], "-")
		arr2 := strings.Split(strings.Split(line, ",")[1], "-")

		i, err := strconv.Atoi(arr1[0])
		if err != nil {
			handleErr(err)
		}

		i2, err := strconv.Atoi(arr1[1])
		if err != nil {
			handleErr(err)
		}

		i3, err := strconv.Atoi(arr2[0])
		if err != nil {
			handleErr(err)
		}

		i4, err := strconv.Atoi(arr2[1])
		if err != nil {
			handleErr(err)
		}

		if overlap([]int{i, i2}, []int{i3, i4}) {
			sum++
		}

	}

	fmt.Println(sum)
}
