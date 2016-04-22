const React = require( 'react' );

module.exports = ( props ) => {

    return(
        <div className="movie-container">
            <div className="movie">
                <h3>{ props.Title }</h3>
                <p>{ props.Released }</p>
                <p>{ props.Plot }</p>
                {/*<img src={ props.Poster } />*/}
            </div>
        </div>
    );
}
