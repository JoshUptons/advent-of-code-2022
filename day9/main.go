package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	UP    int = 0
	RIGHT     = 1
	DOWN      = 2
	LEFT      = 3
)

var (
	directions map[string]int = map[string]int{
		"U": 0,
		"R": 1,
		"D": 2,
		"L": 3,
	}
)

type Rope struct {
	Head   *Knot
	Points map[[2]int]bool
}

type Point struct {
	x float64
	y float64
}

func (r *Rope) addKnot(p Point) {
	if r.Head == nil {

		r.Head = &Knot{p, Point{0, 0}, nil}

	} else {

		head := r.Head
		for head.Next != nil {
			head = head.Next
		}
		head.Next = &Knot{p, Point{0, 0}, nil}

	}
}

type Knot struct {
	Position     Point
	PrevPosition Point
	Next         *Knot
}

func (k *Knot) moveHead(dir int) {
	k.PrevPosition = k.Position
	switch dir {
	case UP:
		k.Position.y++
		break
	case DOWN:
		k.Position.y--
		break
	case LEFT:
		k.Position.x--
		break
	case RIGHT:
		k.Position.x++
		break
	}
	// fmt.Println("head", k.Position)
}

func (k *Knot) move(prev *Knot) {

	offsetX := prev.Position.x - k.Position.x
	offsetY := prev.Position.y - k.Position.y

	if math.Abs(offsetX) > 1 && math.Abs(offsetY) > 1 {

		k.Position.x += offsetX / 2
		k.Position.y += offsetY / 2

	} else if math.Abs(offsetX) > 1 {

		k.Position.x += offsetX / 2
		k.Position.y += offsetY

	} else if math.Abs(offsetY) > 1 {

		k.Position.y += offsetY / 2
		k.Position.x += offsetX

	}

	// fmt.Println(k.Position)

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

func main() {
	file, err := os.Open("input.txt")
	// file, err := os.Open("sample.txt")

	handleErr(err)

	defer file.Close()

	lineScanner := bufio.NewScanner(file)

	rope := Rope{}
	for i := 0; i < 10; i++ {
		rope.addKnot(Point{0, 0})
	}

	rope.Points = map[[2]int]bool{}

	for lineScanner.Scan() {

		text := lineScanner.Text()

		dir, dis := parseLine(text)

		for i := 0; i < dis; i++ {

			curr := rope.Head
			curr.moveHead(directions[dir])
			prev := curr
			curr = curr.Next

			for curr.Next != nil {

				curr.move(prev)
				prev = curr
				curr = curr.Next

			}

			curr.move(prev)

			rope.Points[[2]int{int(curr.Position.x), int(curr.Position.y)}] = true

			// fmt.Println()

		}
	}

	fmt.Println(len(rope.Points), "points visited")

}
