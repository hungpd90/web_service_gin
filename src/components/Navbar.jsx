import React from "react";
import styled from "styled-components";
import MenuIcon from "@mui/icons-material/Menu";
import StoreIcon from "@mui/icons-material/Store";
import SearchIcon from "@mui/icons-material/Search";

const Container = styled.div`
  height: 90px;
  margin: 5px;
`;
const Logo = styled.div``;
const CategoryContainer = styled.div`
  display: flex;
  border: 0.5px solid gray;
  justify-content: center;
  align-items: center;
  padding: 5px;
`;
const Category = styled.div``;
const Wrapper = styled.div`
  display: flex;
`;
const SearchContainer = styled.div`
  border: 0.5px solid gray;
  width: 270px;
`;
const Search = styled.div``;
const Icon = styled.div`
  float: right;
`;
const Navbar = () => {
  return (
    <Container>
      <Wrapper>
        <Logo>
          <StoreIcon></StoreIcon>
        </Logo>

        <CategoryContainer>
          <MenuIcon></MenuIcon>
          <Category>Danh mục sản phẩm</Category>
        </CategoryContainer>
        <SearchContainer>
          <Search></Search>
          <Icon>
            <SearchIcon></SearchIcon>
          </Icon>
        </SearchContainer>
      </Wrapper>
    </Container>
  );
};

export default Navbar;
