const React = require( 'react' );

module.exports =  React.createClass({

    render() {

        return(
        	<div>{ JSON.stringify( this.props.data ) }</div>
    	);
    }
});
