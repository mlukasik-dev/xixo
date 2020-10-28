import { NextPage } from "next";
import ClientsListTable from "@/components/Clients/ClientsListTable";
import Layout from "@/components/Layout";

const Clients: NextPage = () => {
  return (
    <Layout
      title="Your clients are amazing"
      subtitle="We love Jones and Castles Law Firm"
    >
      <ClientsListTable />;
    </Layout>
  );
};

export default Clients;
