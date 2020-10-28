import styled from "styled-components";
import { Button as MuiButton } from "@material-ui/core";

export const Container = styled.div`
  background: #fcfdfe;
  border: 1px solid #f0f1f7;
  box-sizing: border-box;
  border-radius: 8px;
  height: 40px;
  padding: 8px;
`;

export const Button = styled(MuiButton)`
  background: none;
  border: 0;
  min-width: 25px !important;
`;

export const Input = styled.input`
  background: none;
  border: 0;
  font-style: normal;
  font-weight: normal;
  font-size: 14px;
  line-height: 20px;
  letter-spacing: 0.3px;
  color: rgba(#4b506d, 0.4);
  outline: none;
`;
