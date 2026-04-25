import React from "react";
import logo from "../../assets/images/logo.png"
import "./Header.css";

function Header() {
    return (
        <header className="header">
            <img src={logo} alt="Logo" />
            <h1>Mindfull Education</h1>
        </header>
    )
}

export default Header;