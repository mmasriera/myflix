
import './Main.scss';

import React from 'react';
import {render} from 'react-dom';
import List from '../List/List.jsx';
//import Details from '../MovieInfo/MovieInfo.jsx';

class Main extends React.Component {

    constructor() {
        super();
        this.state = {};
    }

    render() {
        return(
            <div id="mainContainer">
                <div id="header">
                    <h3>myflix</h3>
                </div>
                <List />
            </div>
        );
    }
};

render( <Main />, document.getElementById('app') );
