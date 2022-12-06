import React from "react";
import "react-responsive-carousel/lib/styles/carousel.min.css"; // requires a loader
import { Carousel } from "react-responsive-carousel";
import img1 from "../assets/slider1.jpg";
import img2 from "../assets/slider2.jpg";
import img3 from "../assets/slider3.jpg";
import "../css/SliderStyle.css";
const Slider = () => {
  return (
    <div>
      <Carousel autoPlay="true" infiniteLoop="true">
        <div>
          <img src={img1} height="600px" width="300px" />
        </div>
        <div>
          <img src={img2} height="600px" width="300px" />
        </div>
        <div>
          <img src={img3} height="600px" width="300px" />
        </div>
      </Carousel>
    </div>
  );
};

export default Slider;
