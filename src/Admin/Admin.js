import React from "react";
import Bar from "./Bar";
import "./AdminStyle.css";
const Admin = () => {
  return (
    <div className="admin-container">
      <div className="admin-sidebar">
        <Bar></Bar>
      </div>
      <div className="admin-content ">content</div>
    </div>
  );
};

export default Admin;
