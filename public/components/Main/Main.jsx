
import React from 'react';
import List from '../List/List.jsx';
import Details from '../MovieInfo/MovieInfo.jsx';

export default class Main extends React.Component {

    constructor() {

        super();
        this.state = {};
    }

    updateMovie( newMovie ) {}

    handleChange( event ) {}

    render() {

        return(
            <div id="mainContainer">
                <div id="left">
                    <div id="myflixInfo">
                        <h3>myflixInfo</h3>
                        <p>myflix info...</p>
                    </div>
                    {<List />}
                </div>
                <Details />
            </div>
        );
    }
};