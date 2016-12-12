package main

import (
	"fmt"
	"os"
	"strings"

	"./listutils"
)

const titlesFile string = "./test.txt"
const jsonFile string = "./movies.json"
const usage string = `USAGE:
  - go run listUtils.go add movie_title`

func printUsage() {
	fmt.Println(usage)
	os.Exit(0)
}

func main() {
	// read args
	if len(os.Args) < 2 {
		printUsage()
	}
	file, errOpen := os.Open(titlesFile)
	defer file.Close()
	if errOpen != nil {
		panic(errOpen)
	}
	switch os.Args[1] {
	case "add":
		newTitle := strings.Join(os.Args[2:], " ")
		listutils.AddMovie(newTitle, file)
	case "jsonify":
		fmt.Println("giahson")
	default:
		printUsage()
	}
}
