
require( './style.less' );

const React     = require( 'react' ),
      ReactDOM  = require( 'react-dom' ),
      ghLogo    = require( './images/gh-logo.png' ),
      MovieInfo = require( './components/MovieInfo/MovieInfo.jsx' ),
      List      = require( './components/List/List.jsx' );

const App = React.createClass({

    getInitialState() {

        return { movie : null }
    },

    updateMovie( newMovie ) {

        this.setState({ movie : newMovie });
    },

    render() {

        return(
            <div id="container">

                <div id="sidebar">
                    <div id="page-info">
                        <h2>{ "movies I've seen" }</h2>
                        <p>last update 12-06-2016</p>
                        <p>miquel masriera</p>
                        <hr></hr>
                    </div>

                    <MovieInfo { ...this.state.movie } />

                    <a id="src" href="https://github.com/mmasriera/myflix">
                        <img src={ ghLogo } />
                    </a>
                </div>

            	<List updateMovie={ this.updateMovie } />

        	</div>
        );
    }
});

ReactDOM.render( <App />, document.getElementById( 'app' ) );
