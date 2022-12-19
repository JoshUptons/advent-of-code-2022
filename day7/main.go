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
	DIR     string = "dir"
	FILE           = "FILE"
	COMMAND        = "$"
	LS             = "ls"
	CD             = "cd"
	UPDIR          = ".."
	ROOT           = "/"
	MAX     int    = 100000
)

type FileTree struct {
	Root        *Directory
	Current     *Directory
	Directories []*Directory
}

func newDirectory(name string, size int, parent *Directory) Directory {
	return Directory{
		Name:     name,
		Size:     size,
		Children: DirectoryChildren{},
		Parent:   parent,
	}
}

func newFile(name string, size int, parent *Directory) File {
	return File{
		Name:   name,
		Size:   size,
		Parent: parent,
	}
}

type Directory struct {
	Name     string
	Size     int
	Children DirectoryChildren
	Parent   *Directory
}

func (d *Directory) addDir(name string) *Directory {
	newDir := newDirectory(name, 0, d)
	d.Children.Directories = append(d.Children.Directories, &newDir)
	return &newDir
}

func (d *Directory) addFile(name string, size int) {
	newFile := newFile(name, size, d)
	d.Children.Files = append(d.Children.Files, &newFile)
	calcSize(newFile.Size, d)
}

func (d *Directory) childExists(name, childType string) bool {
	if childType == "DIR" {
		for _, dir := range d.Children.Directories {
			if dir.Name == name {
				return true
			}
		}
	} else if childType == "FILE" {
		for _, file := range d.Children.Files {
			if file.Name == name {
				return true
			}
		}
	} else {
		log.Fatalf("Error: %s not valid child type", childType)
		os.Exit(1)
	}
	return false
}

func (d *Directory) print() {
	fmt.Println("dir - ", d.Size, d.Name)
	for _, dir := range d.Children.Directories {
		fmt.Println("\tdir - ", dir.Size, dir.Name)
	}
	for _, file := range d.Children.Files {
		fmt.Println("\tfile - ", file.Size, file.Name)
	}
}

func (d *Directory) isBelowAmount(amount int) bool {
	if d.Size <= amount {
		return true
	}
	return false
}

type DirectoryChildren struct {
	Files       []*File
	Directories []*Directory
}

type File struct {
	Name   string
	Size   int
	Parent *Directory
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func calcSize(size int, dir *Directory) {
	dir.Size += size
	if dir.Parent == nil {
		return
	} else {
		calcSize(size, dir.Parent)
	}
}

func main() {
	file, err := os.Open("input.txt")
	handleErr(err)

	defer file.Close()

	lineScanner := bufio.NewScanner(file)
	ft := FileTree{}

	for lineScanner.Scan() {
		text := lineScanner.Text()

		tokens := strings.Split(text, " ")

		switch tokens[0] {
		case COMMAND:

			if tokens[1] == CD {

				if tokens[2] == ROOT {

					// root dir
					ft.Root = &Directory{ROOT, 0, DirectoryChildren{}, nil}
					ft.Current = ft.Root
					ft.Directories = append(ft.Directories, ft.Root)

				} else if tokens[2] == UPDIR {

					ft.Current = ft.Current.Parent

				} else {

					for _, dir := range ft.Current.Children.Directories {
						if dir.Name == tokens[2] {
							ft.Current = dir
						}
					}

				}
			}
			break
		case DIR:
			if exists := ft.Current.childExists(tokens[1], "DIR"); !exists {
				newDir := ft.Current.addDir(tokens[1])
				ft.Directories = append(ft.Directories, newDir)
			}
			break
		default:
			// file
			if exists := ft.Current.childExists(tokens[1], "FILE"); !exists {
				size, err := strconv.Atoi(tokens[0])
				handleErr(err)
				ft.Current.addFile(tokens[1], size)
			}
		}
	}

	totalBytes := 0
	for _, dir := range ft.Directories {
		if dir.isBelowAmount(MAX) {
			fmt.Printf("%s was below the amount with %d bytes\n", dir.Name, dir.Size)
			totalBytes += dir.Size
		}
	}
	fmt.Printf("%d bytes ready to be deleted\n", totalBytes)

}
