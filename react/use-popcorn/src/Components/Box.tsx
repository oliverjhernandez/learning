import { useState, type ReactElement } from "react";

type BoxProps = {
  children: ReactElement | ReactElement[] | undefined;
};

const Box = (props: BoxProps) => {
  const [isOpen, setIsOpen] = useState(true);

  return (
    <div className="box">
      <button className="btn-toggle" onClick={() => setIsOpen((open) => !open)}>
        {isOpen ? "–" : "+"}
      </button>
      {isOpen && props.children}
    </div>
  );
};

export default Box;
