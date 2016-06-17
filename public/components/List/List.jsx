
const React     = require( 'react' ),
      movieList = require( '../../../lists/list.js' ),
      Movie     = require( '../Movie/Movie.jsx' );

module.exports =  React.createClass({

    selectMovie( movieData ) {

        this.props.updateMovie( movieData );
    },

    render() {

        var searchText = this.props.searchText,
            movies = movieList.filter( ( mv ) => {

                if ( !searchText ) return true; //empty string

                return mv.Title.toLowerCase().includes( searchText ) ||
                    mv.Year.toLowerCase().includes( searchText ) ||
                    mv.Director.toLowerCase().includes( searchText ) ||
                    mv.Actors.toLowerCase().includes( searchText );
            })
            .map( ( movie, idx ) => <Movie movie={ movie } key={ `mov-${idx}` } select={ this.selectMovie } /> );

        return(
        	<div id="movies">{ movies }</div>
    	);
    }
});
