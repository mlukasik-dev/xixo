import { FC } from "react";
import Bell from "@/assets/bell.svg";
import * as S from "./styled";

type Props = {
  hasNew?: boolean;
};

const NotificationBell: FC<Props> = ({ hasNew }) => {
  return (
    <S.NotificationBellContainer>
      {hasNew && <S.NewBadge />}
      <Bell />
    </S.NotificationBellContainer>
  );
};

export default NotificationBell;
