import {} from "./main.tsx";

interface Button1Props {
  count: number;
  onClick: () => void;
}

const Button1: React.FC<Button1Props> = ({ count, onClick }) => {
  return <button onClick={onClick}>Clicked {count} times!</button>;
};

export default Button1;
