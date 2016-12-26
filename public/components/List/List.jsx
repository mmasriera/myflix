
import React from 'react';
import MovieEntry from '../MovieEntry/MovieEntry.jsx';
import movies from './movies.json';

export default class List extends React.Component {

    constructor() {
        super();

        this.state = { movieList : movies };

        this.select = ( title ) => {
            console.log( title );
        };
    }

    render() {
        let movies = this.state.movieList.map( ( movie, idx ) => { 
            movie.select = this.select;
            return <MovieEntry key={ `mv-${ idx }` } { ...movie } />;
        });

        return(
            <div id="list">
                <ul>{ movies }</ul>
            </div>
        );
    }
};
