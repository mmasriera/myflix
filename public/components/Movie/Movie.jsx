const React = require( 'react' );

module.exports = ( props ) => {

    return(

            <div className="movie">
                <img src={ props.Poster }/>
                <p className="title">{ props.Title.substr( 0, 30 ) }</p>
            </div>

    );
}
