package main

import(
	"fmt"
	"bufio"
	"os"
	"net/http"
	"io/ioutil"
	"strings"
	"encoding/json"
)

type Movie struct {
    Title 		string
	Year 		string
	Rated 		string
	Released 	string
	Runtime 	string
	Genre 		string
	Director 	string
	Writer 		string
	Actors 		string
	Plot 		string
	Language 	string
	Country 	string
	Awards 		string
	Poster  	string
	Metascore 	string
	imdbRating 	string
	imdbVotes 	string
	imdbID 		string
	Type 		string
	Response 	string
}

// given a movie title, returns am omdb-paramerter string
func omdbFormatTitle( title string ) string {

	words := strings.Split( title, " " )
	return strings.Join( words, "+" )
}

// given the title of a film, returns a json string with its info
func title2jsonString( title string ) string {

	var formattedTitle string = omdbFormatTitle( title )

	resp, err1 := http.Get( "http://www.omdbapi.com/?t=" + formattedTitle + "&y=&plot=short&r=json" )
	defer resp.Body.Close()

	if err1 != nil {

		panic( err1 )
	}

    jsonData, err2 := ioutil.ReadAll( resp.Body ) // read json http response

	if err2 != nil {

    	panic( err2 )
    }

	var movie Movie
	json.Unmarshal( []byte( jsonData ), &movie )
	// do sth with movie
	jsonEncoded, err3 := json.Marshal( movie )

	if err3 != nil {

    	panic( err3 )
    }

	return string( jsonEncoded )
}

func main() {

	var movies []string

	file, err1 := os.Open( "titles.txt" )
	defer file.Close()

	if err1 != nil {

		fmt.Println( err1 )
	}

	input := bufio.NewScanner( file )

	for input.Scan() {

		var title string = input.Text()
		movies = append( movies, title2jsonString( title ) )
	}

	moviesData := []byte( "module.exports = [" + strings.Join( movies, "," ) + "];" )
	err2 := ioutil.WriteFile( "list.js", moviesData, 0644 )

	if err2 != nil {

		fmt.Println( err2 )
	}
}
