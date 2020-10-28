import { FC, useState } from "react";
import { Paper } from "@material-ui/core";
import { LEADS } from "@/mock-data/leads";
import TableHeader from "@/components/TableHeader/";
import LeadsTable from "./LeadsTable";
import * as StyledTable from "@/components/styled/table";
import * as S from "./styled";

type Props = {};

export type Ordering = {
  direction: "desc" | "asc";
  orderBy: string;
};

export type Pagination = {
  perPage: number;
  page: number;
  total: number;
};

const LeadsListTable: FC<Props> = () => {
  const [pagination, setPagination] = useState({ perPage: 5, page: 1 });
  const [ordering, setOrdering] = useState<Ordering>({
    direction: "asc",
    orderBy: "date",
  });

  const perPage = 5;
  const page = 1;
  const total = 1;

  return (
    <S.Container>
      <StyledTable.Container component={Paper}>
        <TableHeader title="All contacts" />
        <LeadsTable
          rows={LEADS}
          pagination={{
            perPage, // perPage : meta.perPage
            page, // page : meta.page
            total, // total : meta.total
          }}
          ordering={ordering}
        />
      </StyledTable.Container>
    </S.Container>
  );
};

export default LeadsListTable;
