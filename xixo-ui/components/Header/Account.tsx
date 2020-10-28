import { Avatar, Box } from "@material-ui/core";
import * as S from "./styled";

const Account = () => {
  return (
    <Box display="flex" alignItems="center">
      <h4>John Ferdinand</h4>
      <S.Avatar alt="John Ferdinand" src="/avatar.png" />
    </Box>
  );
};

export default Account;
