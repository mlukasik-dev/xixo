import styled from "styled-components";
import { List as MuiList, MenuItem as MuiMenuItem } from "@material-ui/core";

export const Container = styled.nav`
  flex: 255px 0;
  background: #363740;
`;

export const Logo = styled.div`
  padding: 25px 44px 3px;
`;

export const Title = styled.div`
  padding: 22px 22px 5px;
  border-bottom: 1px solid rgba(#dfe0eb, 0.2);
  font-style: normal;
  font-weight: bold;
  font-size: 14px;
  line-height: 18px;
  letter-spacing: 0.2px;
  color: #c5c7cd;
  font-family: "Mulish", sans-serif;
`;

export const List = styled(MuiList).attrs(() => ({
  disablePadding: true,
  component: "ul",
}))`
  padding: 0;
  a {
    color: inherit;
    text-decoration: none;
  }
`;

export const MenuItem = styled(MuiMenuItem)`
  border-left: 3px solid #363740 !important;
  span {
    padding: 4px 10px;
    font-style: normal;
    font-weight: normal;
    font-size: 16px;
    line-height: 20px;
    letter-spacing: 0.2px;
    color: #c5c7cd;
  }
  svg {
    color: #9fa2b4;
    opacity: 0.4;
  }
  &:hover {
    background: rgba(#9fa2b4, 0.08);
    border-left: 3px solid #dde2ff !important;
    color: #dde2ff;
    svg {
      color: #dde2ff;
      opacity: 1;
    }
  }
`;
