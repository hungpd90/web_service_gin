import axios from "axios";
import { loginStart, loginSuccess, loginFailed } from "./authSlice";
export const loginUser = async (user, dispatch, navigate) => {
  dispatch(loginStart());
  try {
    const res = await axios.post("/auth/login", user);
    dispatch(loginSuccess(res.data));
    navigate("/Home");
  } catch (err) {
    dispatch(loginFailed());
  }
};
