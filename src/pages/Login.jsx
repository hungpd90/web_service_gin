import React, { useState } from "react";
import { useDispatch } from "react-redux";
import { NavLink, useNavigate } from "react-router-dom";
import Header from "../components/Header";
import "../css/LoginStyle.css";
import { loginUser } from "../redux/apiRequest";

const Login = () => {
  const [userName, setUserName] = useState("");
  const [passWord, setPassWord] = useState("");
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const handleLogin = (e) => {
    e.preventDefault();
    const newUser = {
      userName: userName,
      passWord: passWord,
    };
    loginUser(newUser, dispatch, navigate);
  };
  return (
    <>
      <div className="headers">
        <Header />
      </div>
      <div className="Auth-form-container">
        <form className="Auth-form" onSubmit={handleLogin}>
          <div className="Auth-form-content">
            <h3 className="Auth-form-title">Đăng Nhập</h3>
            <div className="form-group mt-3">
              <label>Email</label>
              <input
                type="email"
                className="form-control mt-1"
                placeholder="Nhập Email"
                onChange={(e) => setUserName(e.target.value)}
              />
            </div>
            <div className="form-group mt-3">
              <label>Mật khẩu</label>
              <input
                type="password"
                className="form-control mt-1"
                placeholder="Nhập mật khẩu"
                onChange={(e) => setPassWord(e.target.value)}
              />
            </div>
            <div className="d-grid gap-2 mt-3">
              <button type="submit" className="btn btn-primary">
                Đăng Nhập
              </button>
              <p class="signup-text">Hoặc</p>
              <NavLink to="/SignUp" className="signin-nav">
                Đăng Ký
              </NavLink>
            </div>
          </div>
        </form>
      </div>
    </>
  );
};
export default Login;
