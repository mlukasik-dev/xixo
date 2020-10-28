import { Paper } from "@material-ui/core";
import { CLIENTS } from "@/mock-data/clients";
import ClientsTable from "./ClientsTable";
import TableHeader from "@/components/TableHeader";
import * as StyledTable from "@/components/styled/table";
import * as S from "./styled";

const ClientsListTable = () => {
  return (
    <S.Container>
      <StyledTable.Container className={"table-block"} component={Paper}>
        <TableHeader title="" />
        <ClientsTable rows={CLIENTS} />
      </StyledTable.Container>
    </S.Container>
  );
};

export default ClientsListTable;
