import { FC } from "react";
import { Table, TableHead, TableBody, TableRow } from "@material-ui/core";
import * as StyledTable from "@/components/styled/table";
import Icon from "@/components/Icon";
import { Client } from "@/mock-data/clients";

type Props = {
  rows: Client[];
};

const ClientsTable: FC<Props> = ({ rows }) => {
  return (
    <>
      <Table>
        <TableHead>
          <TableRow>
            <StyledTable.Th align="left">Name</StyledTable.Th>
            <StyledTable.Th align="left">Owner</StyledTable.Th>
            <StyledTable.Th align="center">Status</StyledTable.Th>
            <StyledTable.Th align="center">Notes</StyledTable.Th>
            <StyledTable.Th align="center">&nbsp;</StyledTable.Th>
          </TableRow>
        </TableHead>
        <TableBody>
          {rows.map((row) => (
            <TableRow key={row.id}>
              <StyledTable.Td scope="row" align="left">
                {row.name}
              </StyledTable.Td>
              <StyledTable.Td scope="row" align="left">
                {row.owner}
              </StyledTable.Td>
              <StyledTable.Td scope="row" align="left">
                <span className={"btn-status"}>{row.status}</span>
              </StyledTable.Td>
              <StyledTable.Td onClick={() => {}} scope="row" align="center">
                <StyledTable.SourceText>Add</StyledTable.SourceText>
                <StyledTable.SourceText>/ Review</StyledTable.SourceText>
              </StyledTable.Td>
              <StyledTable.Td scope="row" align="center">
                <StyledTable.IconBlock onClick={() => {}}>
                  <Icon name={"folder"} width={"22px"} height={"19px"} />
                </StyledTable.IconBlock>
                <StyledTable.IconBlock>
                  <Icon name={"calendar"} width={"18px"} height={"21px"} />
                </StyledTable.IconBlock>
                <StyledTable.IconBlock onClick={() => {}}>
                  <Icon name={"letter"} width={"22px"} height={"18px"} />
                </StyledTable.IconBlock>
                <StyledTable.IconBlock onClick={() => {}}>
                  <Icon name={"comment"} width={"24px"} height={"21px"} />
                </StyledTable.IconBlock>
              </StyledTable.Td>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </>
  );
};

export default ClientsTable;
