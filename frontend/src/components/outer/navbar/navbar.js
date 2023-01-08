import React from "react";
import { Link } from "react-router-dom";
import "./navbar.css"
import HomeIcon from "../../../images/HomeIcon.png"

const Navbar = () => {
    return (
        <nav className="navbar navbar-expand-md navbar-custom">
            <div className="container-fluid">
                <Link to="/" className="navbar-brand">
                    <img className="image" src={HomeIcon} alt="Icon" />
                </Link>
                <div>
                    <ul className="navbar-nav me-auto mb-2 mb-md-0">
                        <li className="nav-item ">
                            <Link to="/login" className="nav-link btn custom-button" >Login</Link>
                        </li>
                        <li className="nav-item ">
                            <Link to="/register" className="nav-link btn custom-button">Register</Link>
                        </li>
                    </ul>
                </div>
            </div>
        </nav >
    )
}

export default Navbar;