import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import Login from './js-components/login.js';
import Chat from './js-components/chat.js';
import './styles/style.css';

class ChatUI extends React.Component {
    render() {
        return (
            <Router>
                <Switch>
                    <Route path="/" exact component={Login} />
                </Switch>

                <Switch>
                    <Route path='/chat' exact component={Chat} />
                </Switch>

            </Router>
        );
    }
}

ReactDOM.render(
    <ChatUI />,
    document.getElementById('root')
);