import type { ReactElement } from "react";

interface NavProps {
  children: ReactElement | ReactElement[];
}

const Nav = (props: NavProps) => {
  return <nav className="nav-bar">{props.children}</nav>;
};

export default Nav;
