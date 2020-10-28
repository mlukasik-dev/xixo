// Imports

// Node modules
import {
  branch,
  compose,
  defaultProps,
  renderNothing,
  setDisplayName,
  renderComponent,
} from "recompose";

// Local
// Views
import Logo from "./Logo";
import iconInfo from "./iconInfo";
import Oval from "./Oval";
import ArrowNext from "./ArrowNext";
import Search from "./Search";
import Filter from "./Filter";
import Sort from "./Sort";
import Comment from "./Comment";
import Letter from "./Letter";
import Calendar from "./Calendar";
import Folder from "./Folder";
import DotsTop from "./DotsTop";
import PrivateTag from "./PrivateTag";
import Basket from "./Basket";
import Xls from "./Xls";
import DocX from "./Docx";
import Pdf from "./Pdf";

// File
const build = (type, Component) =>
  branch(({ name }) => name === type, renderComponent(Component));

const withDefaultProps = defaultProps({
  width: 22,
  height: 22,
});

const Icon = compose(
  setDisplayName("Icon"),
  withDefaultProps,
  build("logo", Logo),
  build("infoIcon", iconInfo),
  build("ovalIcon", Oval),
  build("arrowNext", ArrowNext),
  build("search", Search),
  build("filter", Filter),
  build("sort", Sort),
  build("comment", Comment),
  build("letter", Letter),
  build("calendar", Calendar),
  build("folder", Folder),
  build("dotsTop", DotsTop),
  build("tag", PrivateTag),
  build("basket", Basket),
  build("xls", Xls),
  build("docx", DocX),
  build("pdf", Pdf)
)(renderNothing());

// Exports
export default Icon;
