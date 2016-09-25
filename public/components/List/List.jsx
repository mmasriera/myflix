
const React = require( 'react' );

module.exports = React.createClass({

    getInitialState() {

        return {
            movieList : [ 1, 2, 3, 4, 5, 6, 7, 8 ]
        };
    },

    render() {

        let movies = this.state.movieList.map( (mv, idx) => <li key={ `mv${idx}` }>{mv}</li> );

        return(
        	<div id="list">
                <ul>{ movies }</ul>
            </div>
    	);
    }
});
