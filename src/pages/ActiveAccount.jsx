import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { toast } from "react-toastify";
import Header from "../components/Header";
import "../css/LoginStyle.css";
import { postActive } from "../services/apiServices";

const ActiveAccount = () => {
  const [email, setEmail] = useState("");
  const [code, setCode] = useState("");
  const navigate = useNavigate();
  const handleActive = async (e) => {
    //validate
    e.preventDefault();
    //
    let data = await postActive(email, code);
    console.log(data.data.body.data.status);
    if (data.data.body.data.status === 0) {
      toast.success("data.message");
      //   console.log(data.data.header.errorDetail.errorMessage);
      //   console.log(data);
      //   navigate("/Home");
      alert(data.data.body.data.message);
      navigate("/Home");
    } else {
      toast.error("");
      alert(data.data.body.data.message);

      //   alert(data.data.header.errorDetail.errorMessage);
    }
  };
  return (
    <>
      <div className="headers">
        <Header />
      </div>
      <div className="Auth-form-container">
        <form className="Auth-form">
          <div className="Auth-form-content">
            <h3 className="Auth-form-title">Xác thực tài khoản</h3>
            <div className="form-group mt-3">
              <label>Email</label>
              <input
                type="text"
                className="form-control mt-1"
                placeholder="Nhập Email"
                onChange={(e) => setEmail(e.target.value)}
              />
            </div>
            <div className="form-group mt-3">
              <label>Mã xác thực</label>
              <input
                type="text"
                className="form-control mt-1"
                placeholder="Nhập mã xác thực"
                onChange={(e) => setCode(e.target.value)}
              />
            </div>

            <div className="d-grid gap-2 mt-3">
              <button
                type="submit"
                className="btn btn-primary"
                onClick={handleActive}
              >
                Đăng Ký
              </button>
            </div>
          </div>
        </form>
      </div>
    </>
  );
};

export default ActiveAccount;
