import styled from "styled-components";

export const Container = styled.div`
  flex: 1 1;
  padding: 35px 30px 60px;
`;

type BtnTableProps = {
  type: string;
};
export const BtnTable = styled.span<BtnTableProps>`
  border-radius: 8px;
  width: 88px;
  height: 24px;
  text-transform: uppercase;
  text-align: center;
  margin: 0 10px 0 0;
  font-style: normal;
  font-weight: bold;
  font-size: 11px;
  display: inline-block;
  color: #fff;
  line-height: 24px;
  background-color: ${(props) => {
    switch (props.type) {
      case "call":
        return "#fec400";
      case "chat":
        return "#3751ff";
      case "sms":
        return "#f12b2c";
      case "form":
        return "#29cc97";
    }
  }};
`;
