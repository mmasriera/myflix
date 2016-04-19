
const React = require( 'react' ),
      movies = require( '../../lists/list.js' );

module.exports =  React.createClass({

    render() {

        return(
        	<div>{ JSON.stringify( movies ) }</div>
    	);
    }
});
