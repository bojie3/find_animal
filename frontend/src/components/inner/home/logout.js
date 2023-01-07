import React from "react";
import { Navigate } from "react-router-dom";

const Logout = () => {
    sessionStorage.removeItem("Auth")
    sessionStorage.removeItem("Refresh")
    return <Navigate to="/" replace={true} />
}

export default Logout