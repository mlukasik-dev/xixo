import XixoGuy from "@/assets/xixo-guy.svg";
import Facebook from "@/assets/facebook.svg";
import Twitter from "@/assets/twitter.svg";
import Linkedin from "@/assets/linkedin.svg";
import * as S from "./styled";

const Slider = () => {
  return (
    <S.Container>
      <S.Slide>
        <XixoGuy />
        <S.Content>
          <S.Title>
            <div>Love ❤ xixo?</div>
            <div>Refer and earn $200</div>
          </S.Title>
          <S.SubTitle>
            You could get a $200 Amazon gift card when your friends sign up.
          </S.SubTitle>
          <S.Input value="https://xixo.com/i/{unique-link}" readOnly />
          <S.Button>Copy your link</S.Button>
          <S.ShareLink>
            <span>Share your link on:</span>
            <Facebook />
            <Twitter />
            <Linkedin />
          </S.ShareLink>
        </S.Content>
      </S.Slide>
    </S.Container>
  );
};

export default Slider;
