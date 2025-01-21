import { useState } from "react";

function App() {
  let [step, setStep] = useState(1);
  const [count, setCount] = useState(0);

  // const handleStepPlus = () => {
  //   setStep((s) => {
  //     return s + 1;
  //   });
  // };

  // const handleStepMinus = () => {
  //   setStep((s) => {
  //     return s - 1;
  //   });
  // };

  const handleCountPlus = () => {
    setCount((c) => {
      return c + step;
    });
  };

  const handleCountMinus = () => {
    setCount((c) => {
      return c - step;
    });
  };

  const calculateDate = () => {
    const currentDate = new Date();
    currentDate.setDate(currentDate.getDate() + count);
    if (count < 0) {
      return `${Math.abs(count)} days ago was ${currentDate.toDateString()}`;
    } else if (count > 0) {
      return `${Math.abs(count)} days from today is ${currentDate.toDateString()}`;
    } else {
      return `today is ${currentDate.toDateString()}`;
    }
  };

  return (
    <>
      <div>
        <input
          type="range"
          defaultValue="1"
          min="1"
          max="10"
          onChange={(e) => {
            setStep(Number(e.target.value));
          }}
        />
        Step: {step}
      </div>
      <div>
        <button onClick={handleCountMinus}>-</button>
        <input
          type="text"
          onChange={(e) => {
            setCount(Number(e.target.value));
          }}
        />
        <button onClick={handleCountPlus}>+</button>
      </div>
      <div>
        <p>{calculateDate()}</p>
      </div>
      {count !== 1 || step !== 1 ? (
        <div>
          <button
            onClick={() => {
              setCount(1);
              setStep(1);
            }}
          >
            Reset
          </button>
        </div>
      ) : null}
    </>
  );
}

export default App;
