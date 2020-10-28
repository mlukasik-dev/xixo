import { NextPage } from "next";
import Layout from "@/components/Layout";
import LeadsListTable from "@/components/Leads";

const Leads: NextPage = () => {
  return (
    <Layout title="Here are your leads" subtitle="We love marketing">
      <LeadsListTable />;
    </Layout>
  );
};

export default Leads;
