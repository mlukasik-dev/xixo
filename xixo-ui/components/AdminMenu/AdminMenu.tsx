import { ReactElement } from "react";
import Link from "next/link";
import Icon from "@/components/Icon";
import * as S from "./styled";

const AdminMenu = () => {
  return (
    <S.Container>
      <S.Logo>
        <Link href="/">
          <a>
            <Icon name="logo" width={"157px"} height={"38px"} />
          </a>
        </Link>
      </S.Logo>
      <S.Title>Main</S.Title>
      <S.List aria-label="menu">
        {MAIN_NAV.map((item) => (
          <Link href={item.link} key={item.id}>
            <a>
              <S.MenuItem>
                {item.icon}
                <span>{item.text}</span>
              </S.MenuItem>
            </a>
          </Link>
        ))}
      </S.List>
      <S.Title>Tools</S.Title>
      <S.List aria-label="menu">
        {TOOLS_NAV.map((item) => (
          <Link href={item.link} key={item.id}>
            <a>
              <S.MenuItem>
                {item.icon}
                <span>{item.text}</span>
              </S.MenuItem>
            </a>
          </Link>
        ))}
      </S.List>
      <S.Title>Settings</S.Title>
      <S.List aria-label="menu">
        {SETTINGS_NAV.map((item) => (
          <Link href={item.link} key={item.id}>
            <a>
              <S.MenuItem>
                {item.icon}
                <span>{item.text}</span>
              </S.MenuItem>
            </a>
          </Link>
        ))}
      </S.List>
    </S.Container>
  );
};

type NavItem = {
  id: string;
  icon: ReactElement;
  text: string;
  link: string;
};

const ovalIcon = <Icon name="ovalIcon" width="16px" height="16px" />;

const MAIN_NAV: NavItem[] = [
  {
    id: "0",
    icon: ovalIcon,
    text: "Dashboard",
    link: "/",
  },
  {
    id: "1",
    icon: ovalIcon,
    text: "Leads",
    link: "/leads",
  },
  {
    id: "2",
    icon: ovalIcon,
    text: "Appointments",
    link: "/appointments",
  },
  {
    id: "3",
    icon: ovalIcon,
    text: "Clients",
    link: "/clients",
  },
  {
    id: "4",
    icon: ovalIcon,
    text: "Reputation",
    link: "/reputation",
  },
  {
    id: "5",
    icon: ovalIcon,
    text: "Insights",
    link: "/insights",
  },
];
const TOOLS_NAV: NavItem[] = [
  {
    id: "0",
    icon: ovalIcon,
    text: "Website",
    link: "/website",
  },
  {
    id: "1",
    icon: ovalIcon,
    text: "Task",
    link: "/task",
  },
  {
    id: "2",
    icon: ovalIcon,
    text: "Marketing",
    link: "/marketing",
  },
  {
    id: "3",
    icon: ovalIcon,
    text: "Workflows",
    link: "/workflows",
  },
];
const SETTINGS_NAV: NavItem[] = [
  {
    id: "0",
    icon: ovalIcon,
    text: "Configure",
    link: "/configure",
  },
  {
    id: "1",
    icon: ovalIcon,
    text: "Data",
    link: "/data",
  },
  {
    id: "2",
    icon: ovalIcon,
    text: "Integrations",
    link: "",
  },
];

export default AdminMenu;
