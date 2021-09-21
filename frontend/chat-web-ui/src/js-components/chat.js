import React from 'react';
import { withRouter } from 'react-router-dom';
class Chat extends React.Component {

    render(props) {
        document.title = "Chat Prototype | Chat UI"
        return (
            <div id="chat_outer_container">
                <div id = "chat_header_row">
                    <h2>Hello {this.props.location.postLogin.userid}</h2>
                </div>
            </div>
        );
    }
}

export default withRouter(Chat);