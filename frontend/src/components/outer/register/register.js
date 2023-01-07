import React from "react";
import { useState, useEffect } from "react";
import axios from "axios";
import Navbar from "../navbar/navbar";
import "./register.css"
import { Navigate } from "react-router-dom";

const Register = () => {
    const [name, setName] = useState("")
    const [password, setPassword] = useState("")
    const [err, setErr] = useState(false)
    const [errMsg, setErrMsg] = useState("")
    const [redirect, setRedirect] = useState(false)

    const submit = (e) => {
        e.preventDefault()

        axios
            .post("http://localhost:8080/users/register", {
                "username": name,
                "password": password,
            })
            .then(resp => {
                console.log(resp)
                setRedirect(true)
            })
            .catch(err => {
                setErr(true)
                setErrMsg(err.response.data.error)
                console.log(err, name, password, "there is error in registration")
            })
    }

    if (redirect) {
        return <Navigate to="/login" replace={true} />
    }

    return (
        <div className="register" style={{ height: "100vh" }}>
            <Navbar />
            {err ? <div className="alert alert-danger" role="alert">{errMsg}</div> : <div></div>}
            <form onSubmit={e => submit(e)}>
                <h1 className="reg-welcome">Welcome aboard!</h1>
                <div className="reg-username">
                    <label htmlFor="exampleInputEmail1">Username</label>
                    <input type="text" className="form-control" required minLength={3} onChange={e => setName(e.target.value)} placeholder="Username" />
                </div>
                <div className="reg-password">
                    <label htmlFor="floatingPassword">Password</label>
                    <input type="password" className="form-control" required minLength={5} onChange={e => setPassword(e.target.value)} placeholder="Password" />
                </div>
                <div className="reg-button">
                    <button type="submit" className="btn btn-primary">Register</button>
                </div>
            </form>
        </div>
    )
}

export default Register