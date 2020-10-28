import Document, { DocumentContext } from "next/document";
import { ServerStyleSheet } from "styled-components";
import { ServerStyleSheets as MaterialUiServerStyleSheets } from "@material-ui/core/styles";

class CustomDocument extends Document {
  static async getInitialProps(ctx: DocumentContext) {
    // create SC sheet
    const sheet = new ServerStyleSheet();
    // create MUI sheets
    const materialUISheets = new MaterialUiServerStyleSheets();
    // keep a reference to original Next renderPage
    const originalRenderPage = ctx.renderPage;

    try {
      // set the renderPage as a function that...
      ctx.renderPage = () =>
        // utilizes the original renderPage function that...
        originalRenderPage({
          // overrides the enhanceApp property with 2 returned wrapped fn's
          enhanceApp: (App) => (props) =>
            // that collects and returns SC + MUI styles from App
            sheet.collectStyles(materialUISheets.collect(<App {...props} />)),
        });

      // invoke internal Next getInitialProps (which executes the above)
      const initialProps = await Document.getInitialProps(ctx);
      return {
        // from getInitialProps, spread out any initial props
        ...initialProps,
        // and apply any initial style tags...
        // and apply the MUI style tags...
        // and apply the SC style tags...
        // to the document's head:
        // <html>
        //   <head>
        //     ...styles are placed here
        //   </head>
        //   <body>...</body>
        // </html>
        styles: (
          <>
            {initialProps.styles}
            {materialUISheets.getStyleElement()}
            {sheet.getStyleElement()}
          </>
        ),
      };
    } finally {
      // seal the SC sheet -- MUI sheets don't need to be sealed or do so internally
      sheet.seal();
    }
  }
}

export default CustomDocument;
