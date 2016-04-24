
const React = require( 'react' ); // do not remove!

module.exports = ( props ) =>

    <div className="movie" onClick={ props.select.bind( null, props.movie ) }>
        <img src={ props.movie.Poster.replace( 'SX300', 'SX125' ) } />
        <p className="title">{ props.movie.Title.substr( 0, 30 ) }</p>
    </div>;
