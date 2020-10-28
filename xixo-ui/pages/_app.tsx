import { useEffect } from "react";
import { AppProps } from "next/app";
import { theme, ThemeProvider } from "@/config/theme";
import { QueryCache, ReactQueryCacheProvider } from "react-query";
import { Hydrate } from "react-query/hydration";
import { ReactQueryDevtools } from "react-query-devtools";

const queryCache = new QueryCache();

const CustomApp = ({ Component, pageProps }: AppProps) => {
  useEffect(() => {
    // MUI has 2 sheets: 1 for server-side and 1 for client-side
    // we don't want this duplication, so during initial client-side load,
    // attempt to locate duplicated server-side MUI stylesheets...
    const jssStyles = document.querySelector("#jss-server-side");
    if (jssStyles && jssStyles.parentNode)
      // ...and if they exist remove them from the head
      jssStyles.parentNode.removeChild(jssStyles);
  }, []);

  return (
    <ReactQueryCacheProvider queryCache={queryCache}>
      <Hydrate state={pageProps.dehydratedState}>
        <ThemeProvider theme={theme}>
          <Component {...pageProps} />
        </ThemeProvider>
      </Hydrate>
      <ReactQueryDevtools />
    </ReactQueryCacheProvider>
  );
};

export default CustomApp;
