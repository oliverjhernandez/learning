import { useState } from "react";

type BoxProps = {
  children?: React.ReactNode;
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
