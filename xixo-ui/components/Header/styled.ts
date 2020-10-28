import styled from "styled-components";
import { Avatar as MuiAvatar } from "@material-ui/core";

export const Container = styled.div`
  display: flex;
  width: 100%;
  height: 129px;
  align-items: center;
`;

export const Title = styled.h3`
  font-weight: 700;
  font-size: 32px;
  line-height: 40px;
  letter-spacing: 0.3px;
  margin: 0;
`;

export const Subtitle = styled.h5`
  font-weight: 700;
  font-size: 14px;
  line-height: 18px;
  letter-spacing: 0.2px;
  color: #9fa2b4;
`;

export const Avatar = styled(MuiAvatar)`
  margin-left: 14px;
  border: 1.5px solid #dfe0eb;
`;

export const NotificationBellContainer = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  width: 49px;
  height: 32px;
  border-left: 1px solid #dfe0eb;
  border-right: 1px solid #dfe0eb;
  margin-left: 32px;
  margin-right: 16px;
  position: relative;
`;

export const NewBadge = styled.div`
  background-color: #3751ff;
  border: 1.5px solid #f7f8fc;
  width: 7px;
  height: 7px;
  position: absolute;
  top: 5px;
  right: 17px;
  border-radius: 50%;
`;
