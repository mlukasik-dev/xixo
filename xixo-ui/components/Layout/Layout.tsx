import Head from "next/head";
import AdminMenu from "../AdminMenu/AdminMenu";
import Header from "@/components/Header";
import * as S from "./styled";

type Props = {
  title: string;
  subtitle: string;
};

const Layout: React.FC<Props> = ({ children, title, subtitle }) => {
  return (
    <>
      <Head>
        <link
          href="https://fonts.googleapis.com/css2?family=Mulish:wght@400;600;700&display=swap"
          rel="stylesheet"
        ></link>
      </Head>
      <S.ParentBlock>
        <AdminMenu />
        <S.Content>
          <Header title={title} subtitle={subtitle} />
          {children}
        </S.Content>
      </S.ParentBlock>
    </>
  );
};

export default Layout;
