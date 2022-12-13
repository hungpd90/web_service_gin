import axios from "axios";

const postLogin = (username, password) => {
  return axios.post("http://localhost:8080/auth/login", { username, password });
};
export { postLogin };
