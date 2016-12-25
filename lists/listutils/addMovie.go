/*Package listutils adds a new movie to the list and sorts the list
TODO
- [x] make it work
- [] test
- [] benchmark
*/
package listutils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type MovieTitles []string

// impl. sort interface
func (mt MovieTitles) Len() int           { return len(mt) }
func (mt MovieTitles) Swap(a, b int)      { mt[a], mt[b] = mt[b], mt[a] }
func (mt MovieTitles) Less(a, b int) bool { return mt[a] < mt[b] }

func lowerCaseFirst(s string) string {
	return strings.ToLower(string(s[0])) + s[1:]
}

// AddMovie adds a movie to the list
func AddMovie(newTitle string, file *os.File) {
	movies := MovieTitles{newTitle}
	input := bufio.NewScanner(file)
	for input.Scan() { // line by line
		title := lowerCaseFirst(input.Text())
		if title != newTitle {
			movies = append(movies, title)
		}
	}
	sort.Sort(movies)
	newList := strings.Join(movies, "\n")
	fmt.Println(newList)
	errWrite := ioutil.WriteFile(file.Name(), []byte(newList), 0644)
	if errWrite != nil {
		panic(errWrite)
	}
	fmt.Println(newTitle + " added")
}
