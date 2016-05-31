
const React = require( 'react' ),
      LazyLoad = require( 'react-lazy-load' ); // TODO: my lazy load

module.exports = ( props ) =>

    <div className="movie" onClick={ props.select.bind( null, props.movie ) }>
        <LazyLoad height={ 148 } offsetVertical={ 150 } >
            <img src={ `./posters/${ props.movie.Title }.jpg` } />
        </LazyLoad>
        <p className="recommended">recommended</p>
    </div>;
