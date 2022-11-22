import { Sell } from "@mui/icons-material";
import React from "react";
import styled from "styled-components";
import HeadphonesIcon from "@mui/icons-material/Headphones";
import NewspaperIcon from "@mui/icons-material/Newspaper";
const Wrapper = styled.div`
  display: flex;
  justify-content: space-between;
  color: white;
  align-items: center;
  height: 100%;
`;
const Left = styled.div`
  flex: 1;
  text-align: center;
`;
const Center = styled.div`
  flex: 1;
`;
const Right = styled.div`
  flex: 1;
`;
const Container = styled.div`
  height: 50px;
  background-color: blue;
`;
const PromotionContainer = styled.div`
  display: flex;
  align-items: center;
  text-align: center;
  justify-content: center;
`;
const Promotion = styled.div`
  cursor: pointer;
  text-align: center;
  margin-left: 8px;
`;
const CareContainer = styled.div`
  display: flex;
  align-items: center;
  text-align: center;
  justify-content: center;
`;
const Care = styled.div`
  cursor: pointer;
  text-align: center;
  margin-left: 8px;
`;
const News = styled.div`
  cursor: pointer;
  text-align: center;
  margin-left: 8px;
`;
const NewsContainer = styled.div`
  display: flex;
  align-items: center;
  text-align: center;
  justify-content: center;
`;
const Header = () => {
  return (
    <Container>
      <Wrapper>
        <Left>
          <PromotionContainer>
            <Sell></Sell>
            <Promotion>Khuyến Mại</Promotion>
          </PromotionContainer>
        </Left>

        <Center>
          <CareContainer>
            <HeadphonesIcon></HeadphonesIcon>
            <Care>CSKH</Care>
          </CareContainer>
        </Center>

        <Right>
          <NewsContainer>
            <NewspaperIcon></NewspaperIcon>
            <News>Tin tức</News>
          </NewsContainer>
        </Right>
      </Wrapper>
    </Container>
  );
};

export default Header;
