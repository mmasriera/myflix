const React    = require( 'react' ),
      ReactDOM = require( 'react-dom' );

const App = React.createClass({

    render() {

        return(
            <div>it works</div>
        );
    }
});

ReactDOM.render( <App />, document.getElementById( 'app' ) );
