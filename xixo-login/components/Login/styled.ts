import styled from "styled-components";
import GoogleButton from "react-google-button";

export const Container = styled.section`
  flex-basis: 557px;
  display: flex;
  flex-direction: column;
  align-items: center;
  background-color: #fff;
  padding-top: 115px;
  padding-bottom: 24px;
`;

export const Heading = styled.h2`
  font-weight: 700;
  font-size: 32px;
  margin-top: 78px;
`;

export const GreyText = styled.span`
  color: ${(props) => props.theme.palette.text.secondary};
`;

export const TextButton = styled.a`
  color: ${(props) => props.theme.palette.primary.main};
  margin-left: 9px;
  :hover {
    text-decoration: underline;
  }
`;

export const Legal = styled.div`
  font-size: 12px;
  line-height: 16px;
  margin-top: 31px;
`;

export const LoginForm = styled.form`
  flex: 1;
`;

export const GoogleSignIn = styled(GoogleButton)`
  outline: none;
  min-width: 300px;
  margin-top: 50px;
`;
