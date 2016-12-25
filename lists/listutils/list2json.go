package listutils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"
)

type Movie struct {
	Title    string
	Year     string
	Director string
}

type Movies struct {
	list []Movie
	sync.Mutex
}

func (mvs *Movies) add(m Movie) {
	mvs.Lock()
	mvs.list = append(mvs.list, m)
	mvs.Unlock()
}

// impl. sort interface
func (mvs Movies) Len() int           { return len(mvs.list) }
func (mvs Movies) Swap(a, b int)      { mvs.list[a], mvs.list[b] = mvs.list[b], mvs.list[a] }
func (mvs Movies) Less(a, b int) bool { return mvs.list[a].Title < mvs.list[b].Title }

// title string to OMDB queryparam
func omdbTitle(title string) string {
	result := strings.Replace(title, " ", "+", -1)
	result = strings.Replace(result, "&", "%26", -1)

	return result
}

func omdbRequest(title string) ([]byte, error) {
	formattedTitle := omdbTitle(title)
	resp, errGet := http.Get("http://www.omdbapi.com/?t=" + formattedTitle + "&y=&plot=short&r=json")
	if errGet != nil {
		return nil, errGet
	}
	bodyData, errBody := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if errBody != nil {
		return nil, errBody
	}

	return bodyData, nil
}

func getMovieData(title string) (Movie, error) {
	var movieData Movie
	jsonData, errReq := omdbRequest(title)
	if errReq != nil {
		return movieData, errReq
	}
	// check movieData ???
	json.Unmarshal(jsonData, &movieData)

	return movieData, nil
}

// Titles2json ...
func Titles2json(titlesFile *os.File, jsonFile string) {
	var movies Movies
	var wg sync.WaitGroup
	count := 0
	input := bufio.NewScanner(titlesFile)
	for input.Scan() { // line by line

		wg.Add(1)

		go func(title string) {
			defer wg.Done()
			fmt.Println("requesting " + title)
			movieData, errGet := getMovieData(title)
			if (errGet != nil) || (movieData.Title == "") {
				fmt.Println("error getting " + title)
				return
			}
			movies.add(movieData)
		}(input.Text())

		if count++; count%25 == 0 {
			wg.Wait()
		}
	}

	wg.Wait()

	sort.Sort(Movies(movies))
	jsonEncoded, errEncJSON := json.Marshal(movies.list) // to json
	if errEncJSON != nil {
		panic(errEncJSON)
	}
	errWrite := ioutil.WriteFile(jsonFile, []byte(jsonEncoded), 0644) // to file
	if errWrite != nil {
		panic(errWrite)
	}
}
