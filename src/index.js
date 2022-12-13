import React from "react";
import ReactDOM from "react-dom";
// import App from "./App";

import { BrowserRouter, Routes, Route } from "react-router-dom";
import Login from "./pages/Login";
import Home from "./pages/Home";
import SignUp from "./pages/SignUp";

ReactDOM.render(
  <BrowserRouter>
    <Routes>
      <Route exact path="/SignUp" element={<SignUp />}></Route>
      <Route exact path="/Home" element={<Home />}></Route>
      <Route exact path="/" element={<Home />}></Route>
      <Route exact path="/Logins" element={<Login />}></Route>
    </Routes>
  </BrowserRouter>,
  document.getElementById("root")
);
