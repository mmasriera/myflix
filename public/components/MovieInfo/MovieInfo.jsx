
const React = require( 'react' );

module.exports = React.createClass({

    render() {

        const movieInfo = this.props.Title ?
            <div>
                <h2>{ this.props.Title }</h2>
                <p>{ `${ this.props.Year } ${ this.props.Country.split( ',' )[ 0 ] }` }</p>
                <p>{ `${ this.props.Runtime } - ${ this.props.Genre }` }</p>
                <p>{ `directed by: ${ this.props.Director }` }</p>
                <p>{ `written by: ${ this.props.Writer }` }</p>
                <p>{ `cast: ${ this.props.Actors }` }</p>
                <p className="plot">{ `"${ this.props.Plot }"` }</p>
                <p>{ `Awards: ${ this.props.Awards }` }</p>
                <p><a href={ `http://www.imdb.com/title/${ this.props.ImdbID }` } target="_blank">imdb</a></p>
            </div>
            :
            <h2>Click a movie</h2>;

        return(
            <div id="movie-info">
                { movieInfo }
            </div>
        );
    }
});
