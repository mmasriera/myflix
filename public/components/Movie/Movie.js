const React = require( 'react' );

module.exports = ( props ) => {

    return(
        <div>
            <h3>{ props.Title }</h3>
            <p>{ props.Released }</p>
            <img src={ props.Poster } />
        </div>
    );
}
