import { createMuiTheme } from "@material-ui/core";

const theme = createMuiTheme({
  typography: {
    fontFamily: ["Mulish", "sans-serif"].join(","),
  },
  palette: {
    primary: {
      main: "#3751FF",
    },
    secondary: {
      main: "#252733",
    },
    text: {
      secondary: "#9FA2B4",
      disabled: "#4B506D",
    },
  },
});

export default theme;
