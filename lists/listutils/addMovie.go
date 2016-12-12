/*Package listutils adds a new movie to the list and sorts the list
TODO
- [x] working
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

type Movies []string

// impl. sort interface
func (m Movies) Len() int           { return len(m) }
func (m Movies) Swap(a, b int)      { m[a], m[b] = m[b], m[a] }
func (m Movies) Less(a, b int) bool { return m[a] < m[b] }

// AddMovie adds a movie to the list
func AddMovie(newTitle string, file *os.File) {
	movies := Movies{newTitle}
	input := bufio.NewScanner(file)
	for input.Scan() { // line by line
		title := input.Text()
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
}
