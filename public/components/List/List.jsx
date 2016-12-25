
import React from 'react';
import MovieEntry from '../MovieEntry/MovieEntry.jsx';
import movies from './movies.json';

export default class List extends React.Component {

    constructor() {
        super();

        console.log(movies);

        this.state = { movieList : [
            'upstream color',
            'inland empire',
            'the darjeeling limited',
            'synecdoche new york'
        ]};

        this.select = ( title ) => {
            console.log( title );
        };
    }

     render() {
        let movies = this.state.movieList.map( ( title, idx ) => { 
            const props = { title, select: this.select };
            return <MovieEntry key={ `mv-${ idx }` } { ...props } />;
        });

        return(
            <div id="list">
                <ul>{ movies }</ul>
            </div>
        );
    }
};
