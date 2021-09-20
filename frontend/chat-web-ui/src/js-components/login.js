import React from 'react';
import { Link } from 'react-router-dom';
import Input from './input.js';
import '../styles/style.css';

class Login extends React.Component {
    render() {
        return (
            <div id="login-container-outer">
                <div id="login-container-title">
                    <h1>User Login</h1>
                </div>
                <div id="login-container-inner">
                    <div className="login_row">
                        <label for="username">Username</label>
                        <Input type="text" id="username" placeholder="Enter your username" />
                    </div>

                    <div className="login_row">
                        <label for="pick-user">Initiate Conversation with</label>
                        <Input type="text" id="pick-user" placeholder="Pick a user from the list" />
                    </div>

                    <div className="login_row">
                        <Link to="/chat">
                            <button>Login</button>
                        </Link>
                    </div>
                </div>
            </div>
        );
    }
}

export default Login;