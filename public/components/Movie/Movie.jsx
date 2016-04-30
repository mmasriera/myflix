
const React = require( 'react' ),
      LazyLoad = require( 'react-lazy-load' );

module.exports = ( props ) =>

    <div className="movie" onClick={ props.select.bind( null, props.movie ) }>
        <LazyLoad height={ 170.2 } offsetVertical={ 170 } >
            <img src={ `./posters/${ props.movie.Title }.jpg` } />
        </LazyLoad>
        {/*<p className="title">{ props.movie.Title.substr( 0, 30 ) }</p>*/}
    </div>;
