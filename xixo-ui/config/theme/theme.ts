import { createMuiTheme } from "@material-ui/core";

const theme = createMuiTheme({
  typography: {
    fontFamily: ["Mulish", "sans-serif"].join(","),
  },
  palette: {
    primary: {
      main: "#000",
    },
  },
});

export default theme;
