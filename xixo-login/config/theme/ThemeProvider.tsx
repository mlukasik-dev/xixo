import { FC } from "react";
import { ThemeProvider as StyledThemeProvider } from "styled-components";
import {
  ThemeProvider as MuiThemeProvider,
  StylesProvider,
} from "@material-ui/core/styles";
import { CssBaseline } from "@material-ui/core";
import { GlobalStyle } from "@/config/theme/GlobalStyle";

type Props = {
  theme: import("@material-ui/core").Theme;
};

const ThemeProvider: FC<Props> = ({ children, theme }) => {
  return (
    <StylesProvider injectFirst>
      <MuiThemeProvider theme={theme}>
        <StyledThemeProvider theme={theme}>
          <CssBaseline />
          <GlobalStyle />
          {children}
        </StyledThemeProvider>
      </MuiThemeProvider>
    </StylesProvider>
  );
};

export default ThemeProvider;
