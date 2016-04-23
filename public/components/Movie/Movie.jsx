const React = require( 'react' );

module.exports = ( props ) => {

    return(

            <div className="movie">
                <img src={ props.Poster }/>
                <p className="title">{ props.Title }</p>
            </div>

    );
}
