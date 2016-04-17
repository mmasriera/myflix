const React    = require( 'react' ),
      ReactDOM = require( 'react-dom' ),
      List     = require( './components/List' );

const App = React.createClass({

    render() {

        return(
            <div>
            	<h3>it works</h3>
            	<List />
        	</div>
        );
    }
});

ReactDOM.render( <App />, document.getElementById( 'app' ) );
