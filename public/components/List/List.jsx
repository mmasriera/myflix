
const React     = require( 'react' ),
      movieList = require( '../../../lists/list.js' ),
      Movie     = require( '../Movie/Movie.jsx' );

module.exports =  React.createClass({

    selectMovie( movieData ) {

        this.props.updateMovie( movieData );
    },

    render() {

        var movies = movieList.map( ( movie, idx ) => <Movie movie={ movie } key={ `mov-${idx}` } select={ this.selectMovie } /> );

        return(
        	<div id="movies">{ movies }</div>
    	);
    }
});
