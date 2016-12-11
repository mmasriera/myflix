/*Package listutils adds a new movie to the list and sorts the list*/
package listutils

import (
	"bufio"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type Movies []string

// impl. sort interface
func (m Movies) Len() int           { return len(m) }
func (m Movies) Swap(a, b int)      { m[a], m[b] = m[b], m[a] }
func (m Movies) Less(a, b int) bool { return m[a] < m[b] }

func AddMovie(newTitle string, titlesFile string) {

	movies := Movies{newTitle}
	file, errOpen := os.Open(titlesFile)
	defer file.Close()
	if errOpen != nil {
		panic(errOpen)
	}
	input := bufio.NewScanner(file)
	for input.Scan() { // line by line
		title := input.Text()
		if title != newTitle {
			movies = append(movies, title)
		}
	}
	sort.Sort(movies)
	newList := strings.Join(movies, "\n")
	errWrite := ioutil.WriteFile(titlesFile, []byte(newList), 0644)
	if errWrite != nil {
		panic(errWrite)
	}
}
