import Square from "./Square";

type square = string;
type squares = square[];

type BoardProps = {
  xIsNext: boolean;
  squares: string[];
  onPlay: (squares: squares) => void;
};

const Board: React.FC<BoardProps> = ({ xIsNext, squares, onPlay }) => {
  const handleClick = (i: number) => {
    if (squares[i] || calculateWinner(squares)) {
      return;
    }

    const nextSquares = squares.slice();
    if (xIsNext) {
      nextSquares[i] = "X";
    } else {
      nextSquares[i] = "O";
    }
    onPlay(nextSquares);
  };

  const winner = calculateWinner(squares);
  let who;
  if (winner) {
    who = "Winner: " + winner;
  } else {
    who = "Next Player: " + (xIsNext ? "X" : "0");
  }

  const rows: JSX.Element[] = [];
  for (let i = 0; i < 3; i++) {
    const sqInRows: JSX.Element[] = [];
    for (let k = 0; k < 3; k++) {
      const index = i * 3 + k;
      sqInRows.push(
        <Square
          key={index}
          value={squares[index]}
          handleClick={() => handleClick(index)}
        />,
      );
    }
    rows.push(
      <div key={i} className="board-row">
        {" "}
        {sqInRows}{" "}
      </div>,
    );
  }

  return (
    <>
      <div className="status">{who}</div>
      {rows}
    </>
  );
};

const calculateWinner = (squares: string[]) => {
  const lines = [
    [0, 1, 2],
    [3, 4, 5],
    [6, 7, 8],
    [0, 3, 6],
    [1, 4, 7],
    [2, 5, 8],
    [0, 4, 8],
    [2, 4, 6],
  ];
  for (let i = 0; i < lines.length; i++) {
    const [a, b, c] = lines[i];
    if (squares[a] && squares[a] === squares[b] && squares[a] === squares[c]) {
      return squares[a];
    }
  }
  return null;
};

export default Board;
