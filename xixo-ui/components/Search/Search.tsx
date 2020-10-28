import { FC, InputHTMLAttributes } from "react";
import Icon from "@/components/Icon";
import * as S from "./styled";

type Props = {} & InputHTMLAttributes<HTMLInputElement>;

const Search: FC<Props> = ({ ...inputProps }) => {
  return (
    <S.Container>
      <S.Button>
        <Icon name="search" width="13px" height="13px" />
      </S.Button>
      <S.Input {...inputProps} />
    </S.Container>
  );
};

export default Search;
