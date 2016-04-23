
const React  = require( 'react' );

module.exports = React.createClass({


    render() {

        var movie = this.props;
        console.log( 'info : ', this.props );

        return(

            <div id="movie-info">
                <h3>{ movie.Title || `click a movie` }</h3>
            </div>
        );
    }
});
