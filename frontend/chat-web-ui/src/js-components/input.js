import React from 'react';
import '../styles/style.css';

class Input extends React.Component {
    render() {
        return (
            <input type={this.props.type} id={this.props.id} placeholder={this.props.placeholder} 
                autoFocus={this.props.autofocus} onKeyUp={this.props.keyup} />
        );
    }
}

export default Input;