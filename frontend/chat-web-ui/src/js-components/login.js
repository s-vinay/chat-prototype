import React from 'react';
import Input from './input.js';
import '../styles/style.css';

class Login extends React.Component {

    constructor(props) {
        super(props);
        this.validateInputs = this.validateInputs.bind(this);
    }

    validateInputs(event) {
        if (document.getElementById("username").value &&
            document.getElementById("pick-user").value) {

            this.props.history.replace({
                pathname: "/chat",
                postLogin: {
                    userid: document.getElementById("username").value,
                    recid: document.getElementById("pick-user").value
                }
            });
        }
    }

    render() {

        const keyupHandler = (event) => {
            if(event.key === "Enter") {
                event.preventDefault();
                document.getElementById("login").click();
            }
        };

        return (
            <div id="login-container-outer">
                <div id="login-container-title">
                    <h1>User Login</h1>
                </div>
                <div id="login-container-inner">
                    <div className="login_row">
                        <label htmlFor="username">Username</label>
                        <Input type="text" id="username" placeholder="Enter your username" 
                            autofocus={true} keyup={keyupHandler} />
                    </div>

                    <div className="login_row">
                        <label htmlFor="pick-user">Initiate Conversation with</label>
                        <Input type="text" id="pick-user" placeholder="Pick a user from the list" 
                            autofocus={false} keyup={keyupHandler} />
                    </div>

                    <div className="login_row">
                        <button type="submit" id="login" onClick={this.validateInputs}>Login</button>
                    </div>
                </div>
            </div>
        );
    }
}

export default Login;