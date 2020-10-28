import styled from "styled-components";
import {
  TableCell as MuiTableCell,
  TableContainer as MuiTableContainer,
} from "@material-ui/core";

type ContainerProps = {
  component?: string | React.FC;
};
export const Container = styled(MuiTableContainer)<ContainerProps>`
  background: #ffffff;
  border: 1px solid #dfe0eb;
  box-sizing: border-box;
  border-radius: 8px;
  width: 100%;

  &.table-documents {
    border: 0;
    margin: 20px 0 0;
    label {
      color: #252733;
    }
    svg {
      color: #dfe0eb;
      margin: 0 12px 0 0;
    }

    a {
      text-decoration: none;
      font-size: 16px;
      line-height: 22px;
      letter-spacing: 0.3px;
      color: #3751ff;
    }
  }

  .btn-status {
    background: #29cc97;
    border-radius: 8px;
    font-style: normal;
    font-weight: bold;
    font-size: 11px;
    line-height: 14px;
    color: white;
    padding: 3px 12px;
    display: inline-block;
  }
`;

export const Th = styled(MuiTableCell)`
  font-style: normal;
  font-weight: bold;
  font-size: 14px;
  line-height: 18px;
  letter-spacing: 0.2px;
  color: #9fa2b4 !important;
`;

export const Td = styled(MuiTableCell).attrs(() => ({
  component: "td",
}))`
  font-size: 14px !important;
  line-height: 20px;
`;

export const IconBlock = styled.span`
  color: #dfe0eb;
  margin: 0 12px;
  cursor: pointer;
`;

export const SourceText = styled.span`
  font-size: 14px;
  line-height: 20px;
  text-align: center;
  letter-spacing: 0.2px;
  color: #3751ff;
  cursor: pointer;
`;

export const PhoneText = styled.span`
  font-size: 14px;
  color: #252733;
  line-height: 22px;
  display: inline-block;
`;

export const UserLink = styled.a`
  color: #252733;
  font-size: 14px;
  text-decoration: none;
`;
