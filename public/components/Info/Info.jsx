
const React  = require( 'react' ),
      ghLogo = require( './gh-logo.png' );

module.exports = React.createClass({

    getInitialState() {

        return {
            movieInfo : 'click a movie'
        };
    },

    render() {

        return(
            <div id="info">

                <div id="page-info">
                    <h2>movies I've seen</h2>
                    <a href="https://github.com/mmasriera/myflix">
                        <img src={ ghLogo } />
                    </a>
                </div>

                <div id="movie-info">
                    { this.state.movieInfo }
                </div>

            </div>
        );
    }
});
