import type { ReactElement } from "react";

type MainProps = {
  children: ReactElement | ReactElement[];
};

const Main = (props: MainProps) => {
  return <main className="main">{props.children}</main>;
};

export default Main;
