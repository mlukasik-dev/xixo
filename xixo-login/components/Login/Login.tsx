import Logo from "@/assets/logo.svg";
import LoginForm from "./LoginForm";
import * as S from "./styled";

const Login = () => {
  return (
    <S.Container>
      <Logo />
      <S.Heading>Welcome Back</S.Heading>
      <LoginForm />
      <div>
        <S.GreyText>Want to sign your firm up with xixo?</S.GreyText>
        <S.TextButton href="/">Click here</S.TextButton>
      </div>
      <S.Legal>
        <S.GreyText>&copy; {new Date().getFullYear()} xixo inc.</S.GreyText>
      </S.Legal>
    </S.Container>
  );
};

export default Login;
