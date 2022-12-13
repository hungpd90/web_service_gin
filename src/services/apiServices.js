import axios from "axios";

const postLogin = (username, password) => {
  return axios.post("http://localhost:8080/auth/login", { username, password });
};
const postSignup = (firstName, lastName, password, email, phoneNumber) => {
  return axios.post("http://localhost:8080/auth/register", {
    firstName,
    lastName,
    password,
    email,
    phoneNumber,
  });
};
const postActive = (email, code) => {
  return axios.post("http://localhost:8080/auth/active-account", {
    email,
    code,
  });
};

export { postLogin, postSignup, postActive };
