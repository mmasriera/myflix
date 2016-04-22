
require( './style.less' );

const React    = require( 'react' ),
      ReactDOM = require( 'react-dom' ),
      List     = require( './components/List/List.jsx' );

const App = React.createClass({

    render() {

        return(
            <div>
            	<List />
        	</div>
        );
    }
});

ReactDOM.render( <App />, document.getElementById( 'app' ) );
