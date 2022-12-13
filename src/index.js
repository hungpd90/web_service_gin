import React from "react";
import ReactDOM from "react-dom";
// import App from "./App";

import { BrowserRouter, Routes, Route } from "react-router-dom";
import Login from "./pages/Login";
import Home from "./pages/Home";
import SignUp from "./pages/SignUp";
import ActiveAccount from "./pages/ActiveAccount";
import Admin from "./Admin/Admin";
import ManageUser from "./Admin/Content/ManageCustomer";
ReactDOM.render(
  <BrowserRouter>
    <Routes>
      <Route exact path="/SignUp" element={<SignUp />}></Route>
      <Route exact path="/Home" element={<Home />}></Route>
      <Route exact path="/" element={<Home />}></Route>
      <Route exact path="/Logins" element={<Login />}></Route>
      <Route exact path="/Active" element={<ActiveAccount />}></Route>
      <Route exact path="/Admin" element={<Admin />}></Route>
      <Route exact path="/Customer" element={<ManageUser />}></Route>
    </Routes>
  </BrowserRouter>,
  document.getElementById("root")
);
