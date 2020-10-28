import { useEffect } from "react";
import * as S from "./styled";

const LoginForm = () => {
  useEffect(() => {
    // FIXME: just quick fix
    if (location.search.includes("redirect")) {
      location.replace("/admin");
    }
  }, []);

  return (
    <S.LoginForm>
      <a href="/auth/google/login">
        <S.GoogleSignIn />
      </a>
    </S.LoginForm>
  );
};

export default LoginForm;
