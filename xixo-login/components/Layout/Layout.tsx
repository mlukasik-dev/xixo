import Head from "next/head";
import * as S from "./styled";

const Layout: React.FC = ({ children }) => {
  return (
    <>
      <Head>
        <link
          href="https://fonts.googleapis.com/css2?family=Mulish:wght@400;600;700&display=swap"
          rel="stylesheet"
        ></link>
      </Head>
      <S.Content>{children}</S.Content>
    </>
  );
};

export default Layout;
