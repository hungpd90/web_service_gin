import React from "react";
import "../css/BannerStyle.css";
import banner1 from "../assets/banner1.jpg";
import banner2 from "../assets/banner2.jpg";
import banner3 from "../assets/banner3.jpg";
import banner4 from "../assets/banner4.jpg";
import laptopblock from "../assets/laptop-cate.jpg";
const Banner = () => {
  return (
    <div>
      <div class="banner">
        <div className="banner-item">
          <a className="banner-link " href="#">
            <div className="banner-img">
              <img src={banner1}></img>
            </div>{" "}
          </a>
        </div>
        <div className="banner-item">
          <a className="banner-link " href="#">
            <div className="banner-img">
              <img src={banner2}></img>
            </div>{" "}
          </a>
        </div>
        <div className="banner-item">
          <a className="banner-link " href="#">
            <div className="banner-img">
              <img src={banner3}></img>
            </div>{" "}
          </a>
        </div>
        <div className="banner-item">
          <a className="banner-link " href="#">
            <div className="banner-img">
              <img src={banner4}></img>
            </div>
          </a>
        </div>
      </div>
      <div className="laptop">
        <div className="latop-block">
          <img src={laptopblock} className="laptop-img"></img>
        </div>
      </div>
    </div>
  );
};

export default Banner;
