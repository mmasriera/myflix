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
	imdbRating 	string
	Metascore 	string
	ImdbVotes 	string
	ImdbID 		string
	Type 		string
	Response 	string
}

/*func getData( url string ) []byte {

	var result

	return
}*/

// given the title of a film, returns a json string with its info
func title2jsonString( title string ) string {

	fmt.Println( title )

	var formattedTitle string = strings.Replace( title, " ", "+", -1 )//omdbFormatTitle( title )
	resp, err1 := http.Get( "http://www.omdbapi.com/?t=" + formattedTitle + "&y=&plot=short&r=json" )

	if err1 != nil {

		panic( err1 )
	}

    jsonData, err2 := ioutil.ReadAll( resp.Body ) // read json http response
	defer resp.Body.Close()

	if err2 != nil {

    	panic( err2 )
    }

	var movie Movie
	json.Unmarshal( jsonData, &movie ) // struct

	// do sth with movie ...
	// save images

	/*var posterURL string = strings.Replace( movie.Poster, "SX300", "SX125", 1 ) // width 125px (original: 300px)
	fmt.Println( posterURL )

	image, errImage := http.Get( posterURL )

	if errImage != nil {

		panic( errImage )
	}*/

	///////

	jsonEncoded, err3 := json.Marshal( movie ) // to string

	if err3 != nil {

    	panic( err3 )
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
		movies = append( movies, title2jsonString( title ) )
	}

	// overwrite file
	moviesData := []byte( "module.exports = [" + strings.Join( movies, "," ) + "];" )
	err2 := ioutil.WriteFile( "list.js", moviesData, 0644 )

	if err2 != nil {

		panic( err2 )
	}
}
