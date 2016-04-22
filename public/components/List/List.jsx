
const React     = require( 'react' ),
      movieList = require( '../../../lists/list.js' ),
      Movie     = require( '../Movie/Movie.jsx' );

module.exports =  React.createClass({

    render() {

        var movies = movieList.map( ( movie, idx ) => <Movie { ...movie } key={ `mov-${idx}` } /> );

        return(
        	<div className="movies">{ movies }</div>
    	);
    }
});
