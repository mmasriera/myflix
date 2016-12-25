package main

import (
	"fmt"
	"os"
	"strings"

	"./listutils"
)

const titlesFileName string = "./files/titles.txt"
const jsonFileName string = "./files/movies.json"
const usage string = `USAGE:
  - go run listUtils.go add movie_title
  - go run listutils.go jsonify`

func printUsage() {
	fmt.Println(usage)
	os.Exit(0)
}

func main() {
	// read args
	if len(os.Args) < 2 {
		printUsage()
	}
	titlesFile, errOpen := os.Open(titlesFileName)
	defer titlesFile.Close()
	defer fmt.Println("- - - done! - - -")
	if errOpen != nil {
		panic(errOpen)
	}
	switch os.Args[1] {
	case "add":
		newTitle := strings.Join(os.Args[2:], " ")
		listutils.AddMovie(newTitle, titlesFile)
	case "jsonify":
		listutils.Titles2json(titlesFile, jsonFileName)
	default:
		printUsage()
	}
}
