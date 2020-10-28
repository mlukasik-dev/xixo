import { FC } from "react";
import { Box } from "@material-ui/core";
import Search from "@/components/Search";
import Icon from "@/components/Icon";
import * as S from "./styled";

type Props = {
  title: string;
};

const TableHeader: FC<Props> = ({ title }) => {
  return (
    <S.Header>
      <h2>{title}</h2>
      <Box alignItems="center" display="flex" flexDirection="row">
        <Search type="search" placeholder="search" />
        <Box justifyContent="space-between" alignItems="center" display="flex">
          <S.HeaderOption>
            <Icon name="sort" width="14px" height="12px" />
            Sort
          </S.HeaderOption>
          <S.HeaderOption>
            <Icon name="filter" width="12px" height="12px" />
            Filter
          </S.HeaderOption>
        </Box>
      </Box>
    </S.Header>
  );
};

export default TableHeader;
