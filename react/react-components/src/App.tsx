import { StrictMode, useState } from "react";
import Button1 from "./Button1";
import Title from "./Title";
import Products from "./Products";
import Profile from "./Profile";

const App = () => {
  const [count, setCount] = useState(0);
  const handleClick = () => {
    setCount(count + 1);
  };

  return (
    <StrictMode>
      <Title />
      <Button1 count={count} onClick={handleClick} />
      <Button1 count={count} onClick={handleClick} />
      <Profile />
      <Products />
    </StrictMode>
  );
};

export default App;
