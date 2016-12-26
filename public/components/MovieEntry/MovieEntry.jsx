
import React from 'react';

export default class MovieEntry extends React.Component {

    constructor( props ) {
        super( props );

        this.handleClick = ( event ) => {
            event.preventDefault();
            props.select( props.title );
        };
    }

    render() {
        return(
            <li onClick={ this.handleClick }>
                <p>{ this.props.Title }</p>
                <p>{ `${ this.props.Director } (${ this.props.Year })` }</p>
            </li>
        );
    }
};
