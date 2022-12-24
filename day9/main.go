package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Knot struct {
	x int
	y int
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func parseLine(line string) (string, int) {

	arr := strings.Split(line, " ")

	val, err := strconv.Atoi(arr[1])
	handleErr(err)

	return arr[0], val

}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

func createKey(tail Knot) string {
	return fmt.Sprint(tail.x, "", tail.y)
}

func keyCheck(tail Knot, grid map[string]bool) {
	key := createKey(tail)
	if _, ok := grid[key]; !ok {
		grid[key] = true
	}

}

func main() {
	file, err := os.Open("input.txt")
	// file, err := os.Open("sample.txt")
	handleErr(err)

	defer file.Close()

	lineScanner := bufio.NewScanner(file)

	grid := map[string]bool{}
	head := Knot{100000, 100000}
	tail := Knot{100000, 100000}
	grid["100000100000"] = true

	for lineScanner.Scan() {
		text := lineScanner.Text()

		direction, distance := parseLine(text)

		for i := 0; i < distance; i++ {
			switch direction {
			case "U":
				fmt.Printf("moving head UP to [%d, %d]\n", head.x, head.y+1)
				head.y++
				fmt.Printf("checking the distance between head and tail\n")

				if head.y-tail.y > 1 {
					fmt.Printf("distance between tail[%d, %d] and head[%d, %d], was greater than 1\nmoving tail\n", tail.x, tail.y, head.x, head.y)
					tail.y++
					tail.x = head.x
					grid[createKey(tail)] = true
					fmt.Println("unique points visited by tail:", len(grid))

				} else {
					fmt.Println("the distance was not greater than 1")

				}
				break

			case "R":
				fmt.Printf("moving head RIGHT to [%d, %d]\n", head.x+1, head.y)
				head.x++

				if head.x-tail.x > 1 {
					fmt.Printf("distance between tail[%d, %d] and head[%d, %d], was greater than 1\nmoving tail\n", tail.x, tail.y, head.x, head.y)
					tail.x++
					tail.y = head.y
					grid[createKey(tail)] = true
					fmt.Println("unique points visited by tail:", len(grid))

				} else {
					fmt.Println("the distance was not greater than 1")

				}
				break

			case "D":
				fmt.Printf("moving head DOWN to [%d, %d]\n", head.x, head.y-1)
				head.y--

				if tail.y-head.y > 1 {
					fmt.Printf("distance between tail[%d, %d] and head[%d, %d], was greater than 1\nmoving tail\n", tail.x, tail.y, head.x, head.y)
					tail.y--
					tail.x = head.x
					grid[createKey(tail)] = true
					fmt.Println("unique points visited by tail:", len(grid))

				} else {
					fmt.Println("the distance was not greater than 1")

				}
				break

			case "L":
				fmt.Printf("moving head LEFT to [%d, %d]\n", head.x-1, head.y)
				head.x--

				if tail.x-head.x > 1 {
					fmt.Printf("distance between tail[%d, %d] and head[%d, %d], was greater than 1\nmoving tail\n", tail.x, tail.y, head.x, head.y)
					tail.x--
					tail.y = head.y
					grid[createKey(tail)] = true
					fmt.Println("unique points visited by tail:", len(grid))

				} else {
					fmt.Println("the distance was not greater than 1")

				}
				break

			}
			fmt.Println()
		}

	}

	fmt.Printf("total points tail has visited: %d\n", len(grid))

}
