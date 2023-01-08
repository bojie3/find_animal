import React from "react";
import { useState, useEffect } from "react";
import axios from "axios";
import { Navigate } from "react-router-dom";
import Navbar from "../navbar/navbar";
import "./login.css"

const Login = () => {
    const [name, setName] = useState("")
    const [password, setPassword] = useState("")
    const [err, setErr] = useState(false)
    const [errMsg, setErrMsg] = useState("")
    const [redirect, setRedirect] = useState(false)

    const submit = (e) => {
        e.preventDefault()
        console.log(name, password)

        axios
            .post("http://localhost:8080/users/login", {
                "username": name,
                "password": password,
            })
            .then(resp => {
                console.log(resp)
                setRedirect(true)
                sessionStorage.setItem("Auth", resp.data.auth)
                sessionStorage.setItem("Refresh", resp.data.refresh)
            })
            .catch(err => {
                setErr(true)
                setErrMsg(err.response.data.error)
                console.log(err, name, password, "there is error in login")
            })
    }

    if (redirect) {
        return <Navigate to="/home" replace={true} />
    }

    return (
        <div className="login" style={{ height: "100vh" }}>
            <Navbar />
            {err ? <div>{errMsg}</div> : <div></div>}
            <form onSubmit={e => submit(e)}>
                <div className="navbar-username">
                    <label htmlFor="exampleInputEmail1">Username</label>
                    <input type="text" className="form-control" required minLength={3} onChange={e => setName(e.target.value)} placeholder="Enter username" />
                </div>
                <div className="navbar-password">
                    <label htmlFor="exampleInputPassword1">Password</label>
                    <input type="password" className="form-control" required minLength={1} onChange={e => setPassword(e.target.value)} placeholder="Password" />
                </div>
                <div className="navbar-button">
                    <button type="submit" className="btn btn-primary">Login</button>
                </div>
            </form>
        </div >
    )
}

export default Login;
