import styled from "styled-components";

export const Header = styled.div`
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 32px;
`;

export const HeaderTitle = styled.h2`
  font-style: normal;
  font-weight: bold;
  font-size: 19px;
  line-height: 24px;
  letter-spacing: 0.4px;
  color: #252733;
  margin: 0 10px 0 0;
`;

export const HeaderOption = styled.span`
  font-weight: 600;
  font-size: 14px;
  line-height: 20px;
  letter-spacing: 0.2px;
  color: #4b506d;
  margin: 0 0 0 30px;
  svg {
    color: #c5c7cd;
    padding: 0 10px 0 0;
  }
`;
