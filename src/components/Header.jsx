/* eslint-disable react/jsx-no-comment-textnodes */
import React from "react";
import { Routes, Route, NavLink } from "react-router-dom";
import "../css/HeaderStyle.css";
import "bootstrap/dist/css/bootstrap.min.css";
import MenuIcon from "@mui/icons-material/Menu";
import StoreIcon from "@mui/icons-material/Store";
import SearchIcon from "@mui/icons-material/Search";
import CircleNotificationsIcon from "@mui/icons-material/CircleNotifications";
import AccountCircleIcon from "@mui/icons-material/AccountCircle";
import ShoppingCartIcon from "@mui/icons-material/ShoppingCart";
import HeadphonesIcon from "@mui/icons-material/Headphones";
import NewspaperIcon from "@mui/icons-material/Newspaper";
import { Sell } from "@mui/icons-material";
import Login from "./../pages/Login";

const Header = () => {
  return (
    <div>
      <div className="header-container">
        <div className="header">
          <div className="header-items">
            <Sell></Sell>
            <div className="header-text">Khuyến Mại</div>
          </div>
          <div className="header-items">
            <HeadphonesIcon></HeadphonesIcon>
            <div className="header-text">CSKH</div>
          </div>
          <div className="header-items">
            <NewspaperIcon></NewspaperIcon>
            <div className="header-text">Tin tức</div>
          </div>
        </div>
        <div className="navbar">
          <div className="shop">
            <div className="shop-logo">
              <StoreIcon></StoreIcon>
            </div>
            <NavLink to="/Home" className="home-tocdo">
              TOCDO.VN
            </NavLink>
            {/* <div className="shop-text">TOCDO.VN</div> */}
          </div>

          <div class="dropdown">
            <button
              class="btn btn-secondary dropdown-toggle bg-transparent text-dark"
              type="button"
              id="dropdownMenu2"
              data-toggle="dropdown"
              aria-haspopup="true"
              aria-expanded="false"
            >
              Danh mục sản phẩm
            </button>
            <div class="dropdown-menu" aria-labelledby="dropdownMenu2">
              <button class="dropdown-item" type="button">
                Action
              </button>
              <button class="dropdown-item" type="button">
                Another action
              </button>
              <button class="dropdown-item" type="button">
                Something else here
              </button>
            </div>
          </div>

          <div className="search">
            <input
              type="text"
              className="text"
              placeholder="Nhập từ khóa cần tìm"
            ></input>
            <button className="search-icon">
              <SearchIcon></SearchIcon>
            </button>
          </div>
          <div className="user">
            <div className="user-icon">
              <AccountCircleIcon></AccountCircleIcon>
            </div>
            <NavLink to="/Logins" className="login-nav">
              Đăng Nhập
            </NavLink>
          </div>
          <div className="noti">
            <div className="noti-icon">
              <CircleNotificationsIcon></CircleNotificationsIcon>
            </div>
          </div>
          <div className="cart">
            <div className="cart-icon">
              <ShoppingCartIcon></ShoppingCartIcon>
            </div>
            <div className="cart-text">Giỏ hàng của bạn</div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Header;
