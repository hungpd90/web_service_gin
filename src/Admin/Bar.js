import { ProSidebarProvider } from "react-pro-sidebar";
import { Link } from "react-router-dom";
import React from "react";
import { Sidebar, Menu, MenuItem, SubMenu } from "react-pro-sidebar";
const Bar = () => {
  return (
    <>
      <ProSidebarProvider>
        <Sidebar>
          <Menu>
            <SubMenu label="Account">
              <MenuItem> Staff</MenuItem>
              <MenuItem routerLink={<Link to="/Customer" />}>
                {" "}
                Customer
              </MenuItem>
            </SubMenu>
            <MenuItem> Product</MenuItem>
            <MenuItem> Order </MenuItem>
          </Menu>
        </Sidebar>
      </ProSidebarProvider>
    </>
  );
};

export default Bar;
