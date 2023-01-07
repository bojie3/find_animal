import React from "react";
import axios from "axios";
import { useState, useEffect } from "react";
import { Navigate, useLocation } from "react-router-dom";

const Private = (props) => {
    const auth = sessionStorage.getItem("Auth")
    const refresh = sessionStorage.getItem("Refresh")
    const [valid, setValid] = useState(true)

    useEffect(() => {
        axios
            .post("http://localhost:8080/auth", {
                "auth": auth,
                "refresh": refresh,
            })
            .then(resp => {
                console.log(resp)
                if (resp.data.Auth) {
                    sessionStorage.setItem("Auth", resp.data.Auth)
                }
                if (resp.data.Refresh) {
                    sessionStorage.setItem("Refresh", resp.data.Refresh)
                }
            })
            .catch(err => {
                console.log(err)
                setValid(false)
            })
    })

    if (!valid) {
        return <Navigate to={"/login"} replace={true} />
    }

    return props.children
}

export default Private