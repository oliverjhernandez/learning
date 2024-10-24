import { useState } from "react";

function App() {
  let [step, setStep] = useState(1);
  const [count, setCount] = useState(0);

  const handleStepPlus = () => {
    setStep((s) => {
      return s + 1;
    });
  };

  const handleStepMinus = () => {
    setStep((s) => {
      return s - 1;
    });
  };

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
        <button onClick={handleStepMinus}>-</button>Step: {step}
        <button onClick={handleStepPlus}>+</button>
      </div>
      <div>
        <button onClick={handleCountMinus}>-</button>Count: {count}
        <button onClick={handleCountPlus}>+</button>
      </div>

      <div>
        <p>{calculateDate()}</p>
      </div>
    </>
  );
}

export default App;
