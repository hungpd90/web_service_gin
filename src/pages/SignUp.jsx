import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

import Header from "../components/Header";
import { postSignup } from "../services/apiServices";
const SignUp = () => {
  const [firstName, setfName] = useState("");
  const [lastName, setlName] = useState("");
  const [password, setPassword] = useState("");
  const [email, setEmail] = useState("");
  const [phoneNumber, setPhoneNumber] = useState("");
  const navigate = useNavigate();
  const handleSignUp = async () => {
    let data = await postSignup(
      firstName,
      lastName,
      password,
      email,
      phoneNumber
    );
    console.log(data.data.header.status);

    if (data.data.header.status === 0) {
      alert(data.data.body.data.message);
      navigate("/Active");
    } else {
      navigate("/SignUp");
      // console.log(data.data.header);
      alert(data.data.header.errorDetail.errorMessage);

      // alert(data.data.message);
    }
  };

  return (
    <>
      <div>
        <Header></Header>
      </div>
      <div styled="margin-top:124px" class="justify-content-center">
        <section
          class="vh-100 justify-content-center"
          styled="background-color: #eee;"
        >
          <div class="container h-100">
            <div class="row d-flex justify-content-center align-items-center h-100">
              <div class="col-lg-12 col-xl-11">
                <div class="card text-black" styled="border-radius: 25px;">
                  <div class="card-body p-md-5">
                    <div class="row justify-content-center">
                      <div class="col-md-10 col-lg-6 col-xl-5 order-2 order-lg-1">
                        <p class="text-center h1 fw-bold mb-5 mx-1 mx-md-4 mt-4">
                          Đăng Ký
                        </p>

                        <form class="mx-1 mx-md-4">
                          <div class="d-flex flex-row align-items-center mb-4">
                            <i class="fas fa-user fa-lg me-3 fa-fw"></i>
                            <div class="form-outline flex-fill mb-0">
                              <label class="form-label" for="form3Example1c">
                                Họ
                              </label>
                              <input
                                name="firstName"
                                type="text"
                                id="form3Example1c"
                                class="form-control"
                                onChange={(event) =>
                                  setfName(event.target.value)
                                }
                              />
                            </div>
                          </div>
                          <div class="d-flex flex-row align-items-center mb-4">
                            <i class="fas fa-user fa-lg me-3 fa-fw"></i>
                            <div class="form-outline flex-fill mb-0">
                              <label class="form-label" for="form3Example1c">
                                Tên
                              </label>
                              <input
                                name="lastName"
                                type="text"
                                id="form3Example1c"
                                class="form-control"
                                onChange={(event) =>
                                  setlName(event.target.value)
                                }
                              />
                            </div>
                          </div>
                          <div class="d-flex flex-row align-items-center mb-4">
                            <i class="fas fa-envelope fa-lg me-3 fa-fw"></i>
                            <div class="form-outline flex-fill mb-0">
                              <label class="form-label" for="form3Example3c">
                                Email
                              </label>
                              <input
                                type="email"
                                id="form3Example3c"
                                class="form-control"
                                onChange={(event) =>
                                  setEmail(event.target.value)
                                }
                              />
                            </div>
                          </div>
                          <div class="d-flex flex-row align-items-center mb-4">
                            <i class="fas fa-envelope fa-lg me-3 fa-fw"></i>
                            <div class="form-outline flex-fill mb-0">
                              <label class="form-label" for="form3Example3c">
                                Số Điện Thoại
                              </label>
                              <input
                                type="email"
                                id="form3Example3c"
                                class="form-control"
                                onChange={(event) =>
                                  setPhoneNumber(event.target.value)
                                }
                              />
                            </div>
                          </div>
                          <div class="d-flex flex-row align-items-center mb-4">
                            <i class="fas fa-lock fa-lg me-3 fa-fw"></i>
                            <div class="form-outline flex-fill mb-0">
                              <label class="form-label" for="form3Example4c">
                                Mật khẩu
                              </label>
                              <input
                                type="password"
                                id="form3Example4c"
                                class="form-control"
                                onChange={(event) =>
                                  setPassword(event.target.value)
                                }
                              />
                            </div>
                          </div>

                          <div class="form-check d-flex justify-content-center mb-5">
                            <input
                              class="form-check-input me-2"
                              type="checkbox"
                              value=""
                              id="form2Example3c"
                            />
                            <label class="form-check-label" for="form2Example3">
                              Tôi đồng ý với tất cả các tuyên bố trong{" "}
                              <a href="#!">Điều khoản sử dụng</a>
                            </label>
                          </div>

                          <div class="d-flex justify-content-center mx-4 mb-3 mb-lg-4">
                            <button
                              type="button"
                              class="btn btn-primary btn-lg"
                              onClick={handleSignUp}
                            >
                              Đăng Ký
                            </button>
                          </div>
                        </form>
                      </div>
                      {/* <div class="col-md-10 col-lg-6 col-xl-7 d-flex align-items-center order-1 order-lg-2"> */}
                      {/* <img src={signup} alt="" /> */}
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          {/* </div> */}
        </section>
      </div>
    </>
  );
};

export default SignUp;
