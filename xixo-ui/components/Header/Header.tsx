import { FC } from "react";
import Search from "@/components/Search";
import Account from "./Account";
import { Box } from "@material-ui/core";
import Select from "@/components/Select";
import NotificationBell from "./NotificationBell";
import * as S from "./styled";

type Props = {
  title: string;
  subtitle: string;
};

const Header: FC<Props> = ({ title, subtitle }) => {
  return (
    <S.Container>
      <Box flex={1}>
        <S.Title>{title}</S.Title>
        <S.Subtitle>{subtitle}</S.Subtitle>
      </Box>
      <Search placeholder="Search contacts" />
      <Select />
      <NotificationBell hasNew />
      <Account />
    </S.Container>
  );
};

export default Header;
