/*
'npm run list2json'
downloads info from every movie of titles.txt
result:
	 list.js : array, each element contains movie info
	 public/images/posters : 1 image / movie
*/

package main

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

// Movie data, fields of the json response from ombd
type Movie struct {
	Title    string
	Year     string
	Released string
	Runtime  string
	Genre    string
	Director string
	Writer   string
	Actors   string
	Plot     string
	Language string
	Country  string
	Awards   string
	Poster   string
	Error    string
	ImdbID   string
	//ImdbRating 	string
	//Metascore 	string
	//ImdbVotes 	string
	//Type 		    string
	//Response 		string
	//Rated 		string
}

// MoviesData lsit of movies and mutex to safely update it
type MoviesData struct {
	list []Movie
	sync.Mutex
}

// implementing sort interface
func (mvs MoviesData) Len() int           { return len(mvs.list) }
func (mvs MoviesData) Swap(a, b int)      { mvs.list[a], mvs.list[b] = mvs.list[b], mvs.list[a] }
func (mvs MoviesData) Less(a, b int) bool { return mvs.list[a].Title < mvs.list[b].Title }

// makes an http.get request, and returns the response
func getData(url string) []byte {

	resp, errGet := http.Get(url)
	if errGet != nil {

		panic(errGet)
	}

	bodyData, errBody := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if errBody != nil {

		panic(errBody)
	}

	return bodyData
}

// omdb format
func omdbFormatTitle(title string) string {

	result := strings.Replace(title, " ", "+", -1)
	result = strings.Replace(result, "&", "%26", -1)

	return result
}

// retreives the data of a movie, and adds it to the MoviesData list, an image of the movie is saved
func (mvs *MoviesData) getMovieData(title string) {

	fmt.Println("title:", title)
	formattedTitle := omdbFormatTitle(title)
	jsonData := getData("http://www.omdbapi.com/?t=" + formattedTitle + "&y=&plot=short&r=json") // get movie data

	var movie Movie
	json.Unmarshal(jsonData, &movie)
	if movie.Error == "Movie not found!" {

		fmt.Println("error:", title, movie.Error)
		return
	}

	if movie.Poster != "N/A" {

		image := getData(strings.Replace(movie.Poster, "SX300", "SX95", 1))             // width from 300 to 95px
		errImg := ioutil.WriteFile("../build/posters/"+movie.Title+".jpg", image, 0644) // save poster
		if errImg != nil {

			panic(errImg)
		}
	}

	mvs.Lock()
	mvs.list = append(mvs.list, movie) // add movie to the list
	mvs.Unlock()
}

func main() {

	file, err1 := os.Open("titles.txt")
	defer file.Close()
	if err1 != nil {

		panic(err1)
	}

	var wg sync.WaitGroup
	var movies MoviesData

	input := bufio.NewScanner(file)
	for input.Scan() { // line by line

		wg.Add(1)

		go func(title string) {

			defer wg.Done()
			movies.getMovieData(title)

		}(input.Text())
	}

	wg.Wait()

	sort.Sort(MoviesData(movies)) // sort movies

	jsonEncoded, errEncJSON := json.Marshal(movies.list) // to json
	if errEncJSON != nil {

		panic(errEncJSON)
	}

	content := "module.exports = " + string(jsonEncoded) + ";" // to string
	err2 := ioutil.WriteFile("list.js", []byte(content), 0644)
	if err2 != nil {

		panic(err2)
	}
}
