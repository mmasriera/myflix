
const React = require( 'react' ),
      movieList = require( '../../../lists/list.js' ),
      Movie = require( '../Movie/Movie' );

module.exports =  React.createClass({

    render() {

        var movies = movieList.map( ( movie, idx ) => <Movie data={ movie } key={ `mov-${idx}` } /> );

        return(
        	<div>{ movies }</div>
    	);
    }
});
