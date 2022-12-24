package main

import (
	"bufio"
	"fmt"
	"log"
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

type Node struct {
	x    int
	y    int
	seen bool
	key  string
}

type Path struct {
	nodes       map[string]*Node
	head        *Node
	tail        *Node
	uniqueNodes int
}

func (p *Path) init() {
	p.nodes = make(map[string]*Node)
	p.head = p.add(0, 0)
	p.tail = p.head
	p.tail.seen = true
	p.uniqueNodes = 1
}

func (p *Path) print() {
	fmt.Println("head", p.head)
	fmt.Println("tail", p.tail)
	fmt.Println("current unique nodes seen: ", p.uniqueNodes)
}

func (p *Path) add(x, y int) *Node {
	key := createKey(x, y)
	newNode := &Node{x, y, false, key}
	p.nodes[key] = newNode
	return newNode
}

func (p *Path) moveTail() {
	p.tail = p.head
}

func (p *Path) keyExists(key string) bool {
	_, ok := p.nodes[key]
	return ok
}

func (p *Path) seeTail() {
	p.tail.seen = true
	p.uniqueNodes++
}

func (p *Path) step(dir int) {
	switch dir {
	case UP:
		if p.head.y+1 > p.tail.y+1 {
			p.moveTail()
		}
		key := createKey(p.head.x, p.head.y+1)
		if !p.keyExists(key) {
			p.head = p.add(p.head.x, p.head.y+1)
		} else {
			p.head = p.nodes[key]
		}
		break

	case DOWN:
		if p.head.y-1 < p.tail.y-1 {
			p.moveTail()
		}
		key := createKey(p.head.x, p.head.y-1)
		if !p.keyExists(key) {
			p.head = p.add(p.head.x, p.head.y-1)
		} else {
			p.head = p.nodes[key]
		}
		break

	case RIGHT:
		if p.head.x+1 > p.tail.x+1 {
			p.moveTail()
		}
		key := createKey(p.head.x+1, p.head.y)
		if !p.keyExists(key) {
			p.head = p.add(p.head.x+1, p.head.y)
		} else {
			p.head = p.nodes[key]
		}
		break

	case LEFT:
		if p.head.x-1 < p.tail.x-1 {
			p.moveTail()
		}
		key := createKey(p.head.x-1, p.head.y)
		if !p.keyExists(key) {
			p.head = p.add(p.head.x-1, p.head.y)
		} else {
			p.head = p.nodes[key]
		}
		break

	default:
		log.Fatal("invalid direction provided")
	}

	if !p.tail.seen {
		p.seeTail()
	}
}

func createKey(x, y int) string {
	return fmt.Sprint(x, "", y)
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	file, err := os.Open("input.txt")
	// file, err := os.Open("sample.txt")
	handleErr(err)

	defer file.Close()

	p := Path{}
	p.init()
	lineScanner := bufio.NewScanner(file)

	for lineScanner.Scan() {
		text := lineScanner.Text()

		arr := strings.Split(text, " ")

		direction := directions[arr[0]]
		distance, err := strconv.Atoi(arr[1])
		handleErr(err)

		for i := 0; i < distance; i++ {
			p.step(direction)
		}

	}

	fmt.Println(p.uniqueNodes)

}
