import React from "react";
import logo from "../octocat.png";

const Header = () => (
    <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p><code>TODO app: </code> What's next?
        </p>
    </header>
);

export default Header;