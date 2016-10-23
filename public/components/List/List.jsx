
const React = require( 'react' );

module.exports = React.createClass({

    getInitialState() {

        let moviesArray = Array.from( new Array(200) );

        return {
            movieList : moviesArray
        };
    },

    render() {

        let movies = this.state.movieList.map( (_, idx) => <li key={ `mv${idx}` }>{ `Movie #${idx}` }</li> );

        return(
        	<div id="list">
                <ul>{ movies }</ul>
            </div>
    	);
    }
});
