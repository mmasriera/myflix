
require( './style.less' );

const React    = require( 'react' ),
      ReactDOM = require( 'react-dom' ),
      Info     = require( './components/Info/Info.jsx' ),
      List     = require( './components/List/List.jsx' );

const App = React.createClass({

    render() {

        return(
            <div id="container">
                {<Info />}
            	<List />
        	</div>
        );
    }
});

ReactDOM.render( <App />, document.getElementById( 'app' ) );
