import { NextPage } from "next";
import Layout from "@/components/Layout";
import Login from "@/components/Login";
import Slider from "@/components/Slider";

const Index: NextPage = () => {
  return (
    <Layout>
      <Login />
      <Slider />
    </Layout>
  );
};

export default Index;
