/*
'npm run list2json'
downloads info from every movie of titles.txt
result:
	 list.js : array, each element contains movie info
	 public/images/posters : 1 image / movie
*/

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
	ImdbRating 	string
	Metascore 	string
	ImdbVotes 	string
	ImdbID 		string
	Type 		string
	Response 	string
}

// makes an http.get request, and returns the response
func getData( url string ) []byte {

	resp, errGet := http.Get( url )

	if errGet != nil {

		panic( errGet )
	}

    bodyData, errBody := ioutil.ReadAll( resp.Body )
	defer resp.Body.Close()

	if errBody != nil {

    	panic( errBody )
    }

	return bodyData
}

// given the title of a film, gets its info (omdb) and saves its poster
func title2jsonString( title string ) string {

	// get movie data
	var formattedTitle string = strings.Replace( title, " ", "+", -1 )//omdbFormatTitle( title )
	var jsonData []byte = getData( "http://www.omdbapi.com/?t=" + formattedTitle + "&y=&plot=short&r=json" )
	var movie Movie
	json.Unmarshal( jsonData, &movie ) // struct
	jsonEncoded, err3 := json.Marshal( movie ) // to string

	if err3 != nil {

    	panic( err3 )
    }

	// save poster
	var image []byte = getData( strings.Replace( movie.Poster, "SX300", "SX110", 1 ) ) // width 125px (original: 300px)
	errImg := ioutil.WriteFile( "../public/images/posters/" + movie.Title + ".jpg" , image, 0644 )

	if errImg != nil {

		panic( errImg )
	}

	return string( jsonEncoded )
}

func main() {

	file, err1 := os.Open( "titles.txt" )
	defer file.Close()

	if err1 != nil {

		panic( err1 )
	}

	var movies []string
	input := bufio.NewScanner( file )
	for input.Scan() { // line by line

		var title string = input.Text()
		fmt.Println( title )
		movies = append( movies, title2jsonString( title ) )
	}

	// overwrite file
	moviesData := []byte( "module.exports = [" + strings.Join( movies, "," ) + "];" )
	err2 := ioutil.WriteFile( "list.js", moviesData, 0644 )

	if err2 != nil {

		panic( err2 )
	}
}
