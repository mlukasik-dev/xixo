import { FC } from "react";
import Link from "next/link";
import {
  Table,
  TableHead,
  TableBody,
  TableRow,
  TableCell,
  Box,
  TableSortLabel,
  TablePagination,
} from "@material-ui/core";
import Icon from "@/components/Icon";
import { Ordering, Pagination } from "./LeadsListTable";
import { Lead } from "@/mock-data/leads";
import * as StyledTable from "@/components/styled/table";
import * as S from "./styled";

type Props = {
  rows: Lead[];
  pagination: Pagination;
  ordering: Ordering;
};

const LeadsTable: FC<Props> = ({ rows, pagination, ordering }) => {
  return (
    <>
      <Table>
        <TableHead>
          <TableRow>
            <StyledTable.Th align="left">
              <TableSortLabel
                active={ordering.orderBy === "date"}
                direction={ordering.direction}
                onClick={() => {}}
              >
                Date
              </TableSortLabel>
            </StyledTable.Th>
            <StyledTable.Th color="red" align="left">
              <TableSortLabel
                active={ordering.orderBy === "name"}
                direction={ordering.direction}
                onClick={() => {}}
              >
                Name
              </TableSortLabel>
            </StyledTable.Th>
            <StyledTable.Th align="left">Type</StyledTable.Th>
            <StyledTable.Th align="center">Source</StyledTable.Th>
            <TableCell align="center">&nbsp;</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {rows.map((row) => (
            <TableRow key={row.id}>
              <StyledTable.Td scope="row" align="left">
                {new Date(Date.now()).toDateString()}
              </StyledTable.Td>
              <StyledTable.Td scope="row" align="left">
                <Link href={`leads/${row.id}`}>
                  <StyledTable.UserLink>{row.name}</StyledTable.UserLink>
                </Link>
              </StyledTable.Td>
              <StyledTable.Td scope="row" align="left">
                <S.BtnTable type={row.type}>{row.type}</S.BtnTable>
                <StyledTable.PhoneText>{row.phone}</StyledTable.PhoneText>
              </StyledTable.Td>
              <StyledTable.Td scope="row" align="center">
                <StyledTable.SourceText>{row.source}</StyledTable.SourceText>
              </StyledTable.Td>
              <StyledTable.Td scope="row" align="center">
                <StyledTable.IconBlock as="a" href={`mailto:${row.email}`}>
                  <Icon name={"letter"} width={"22px"} height={"18px"} />
                </StyledTable.IconBlock>
                <StyledTable.IconBlock>
                  <Icon name={"comment"} width={"24px"} height={"21px"} />
                </StyledTable.IconBlock>
              </StyledTable.Td>
            </TableRow>
          ))}
        </TableBody>
      </Table>
      <Box display="flex" alignItems="center" justifyContent="center">
        <TablePagination
          rowsPerPageOptions={[5, 10, 25]}
          component="div"
          count={pagination.total}
          rowsPerPage={pagination.perPage}
          page={pagination.page - 1}
          onChangePage={(event, page) => {}}
          onChangeRowsPerPage={(event) => {}}
        />
      </Box>
    </>
  );
};

export default LeadsTable;
