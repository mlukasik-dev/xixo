import styled from "styled-components";

export const Container = styled.section`
  position: relative;
  background-color: ${(props) => props.theme.palette.secondary.main};
  flex: 1;
  color: #fff;
`;

export const Slide = styled.div`
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 80%;
  display: flex;
  justify-content: center;
`;

export const Content = styled.div`
  max-width: 335px;
  margin-left: 73px;
`;

export const Title = styled.h1`
  max-width: 319px;
  font-weight: 700;
  font-size: 32px;
  letter-spacing: 0.3px;
  margin-bottom: 9px;
`;

export const SubTitle = styled.h2`
  margin-top: 0;
  max-width: 343px;
  font-weight: 700;
  font-size: 19px;
  letter-spacing: 0.4px;
`;

export const Input = styled.input`
  /* TODO: put into theme */
  background-color: #fcfdfe;
  color: ${(props) => props.theme.palette.text.disabled};
  outline: none;
  border: 1px solid #f0f1f7;
  min-width: 335px;
  padding: 10px 16px;
  text-align: center;
  border-radius: 8px;
`;

export const Button = styled.button`
  outline: none;
  border: 0;
  border-radius: 8px;
  background-color: ${(props) => props.theme.palette.primary.main};
  min-width: 335px;
  padding: 15px 24px;
  color: #fff;
  font-size: 14px;
  letter-spacing: 0.2px;
  margin-top: 11px;
`;

export const ShareLink = styled.div`
  margin-top: 28px;
  display: flex;
  align-items: center;
  span {
    margin-right: 20px;
  }
  svg {
    margin-right: 15px;
  }
`;
