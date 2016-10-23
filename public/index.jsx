
require( './style.scss' );

const React     = require( 'react' ),
      ReactDOM  = require( 'react-dom' ),
      List      = require( './components/List/List.jsx' ),
      Details   = require( './components/Details/Details.jsx' );
      /*ghLogo    = require( './images/gh-logo.png' ),
      MovieInfo = require( './components/MovieInfo/MovieInfo.jsx' ),
      List      = require( './components/List/List.jsx' );*/

const App = React.createClass({

    getInitialState() {

        //return { movie : null, searchText : '' }
        return null;
    },

    updateMovie( newMovie ) {

        //this.setState({ movie : newMovie });
    },

    handleChange( event ) {

        //this.setState({ searchText : event.target.value });
    },

    render() {

        /*return(
            <div id="container">

                <div id="sidebar">
                    <div id="page-info">
                        <h2>{ "movies I've seen" }</h2>
                        <p>last update 11-09-2016</p>
                        <hr></hr>
                    </div>

                    <div id="search" onChange={ this.handleChange }>
                        <input type="text" placeholder="Title, Director, Actor or Year"/>
                    </div>

                    <MovieInfo { ...this.state.movie } />

                    <a id="src" href="https://github.com/mmasriera/myflix">
                        <img src={ ghLogo } />
                    </a>
                </div>

            	<List updateMovie={ this.updateMovie } searchText={ this.state.searchText }/>

        	</div>
        );*/
        return(
            <div id="mainContainer">
                <div id="left">
                    <div id="myflixInfo">
                        <h3>myflixInfo</h3>
                        <p>myflix info...</p>
                    </div>
                    <List />
                </div>
                <Details />
            </div>
        );
    }
});

ReactDOM.render( <App />, document.getElementById( 'app' ) );
