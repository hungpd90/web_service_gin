import React from "react";
import Bar from "../Bar";
import TableUser from "./TableUser";

const ManageCustomer = () => {
  return (
    <>
      <div className="table-content">
        <Bar></Bar>
        <div className="table">
          <h2 style={{ margin: "10px 10px" }}>Customer</h2>
          <div className="table-user">
            <TableUser></TableUser>
          </div>
        </div>
      </div>
    </>
  );
};

export default ManageCustomer;
