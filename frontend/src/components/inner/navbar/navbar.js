import React from "react";
import { Link } from "react-router-dom";

const InnerNavbar = () => {
    return (
        <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4">
            <div className="container-fluid">
                <Link to="/home" className="navbar-brand">Penguins</Link>
                <div>
                    <ul className="navbar-nav me-auto mb-2 mb-md-0">
                        <li className="nav-item">
                            <Link to="/logout" className="nav-link">logout</Link>
                        </li>
                    </ul>
                </div>
            </div>
        </nav>
    )
}

export default InnerNavbar;