import { BrowserRouter, Route, Routes } from "react-router-dom";

import './App.css';
import Home from "./components/inner/home/home";
import Homepage from "./components/outer/homepage/homepage";
import Login from "./components/outer/login/login";
import Register from "./components/outer/register/register";
import Logout from "./components/inner/home/logout";
import Public from "./routes/public";
import Private from "./routes/private";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Public children={<Homepage />} />} />
        <Route path="/login" element={<Public children={<Login />} />} />
        <Route path="/register" element={<Public children={<Register />} />} />
        <Route path="/logout" element={<Public children={<Logout />} />} />
        <Route path="/home" element={<Private children={<Home />} />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
