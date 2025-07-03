import { ReactNode, useState } from "react";
import "./index.css";

const messages = [
  "Learn React ⚛️",
  "Apply for jobs 💼",
  "Invest your new income 🤑",
];

function App() {
  let [step, setStep] = useState(1);
  let [isOpen, setIsOpen] = useState(true);

  const handlePrevious = () => {
    if (step > 1) {
      setStep((s) => {
        return s - 1;
      });
    }
  };
  const handleNext = () => {
    if (step < 3) {
      setStep((s) => {
        return s + 1;
      });
    }
  };

  return (
    <>
      <button className="close" onClick={() => setIsOpen(!isOpen)}>
        &times;
      </button>
      {isOpen && (
        <div className="steps">
          <div className="numbers">
            <div className={`${step >= 1 ? "active" : ""}`}>1</div>
            <div className={`${step >= 2 ? "active" : ""}`}>2</div>
            <div className={`${step == 3 ? "active" : ""}`}>3</div>
          </div>

          <StepMessage step={step}>{messages[step - 1]}</StepMessage>

          <div className="buttons">
            <Button onCLick={handlePrevious}>
              <span>👈🏻</span>Prevous
            </Button>
            <Button onCLick={handleNext}>
              <span>Next 👉🏻</span>
            </Button>
          </div>
        </div>
      )}
    </>
  );
}

type ButtonProps = {
  onCLick: () => void;
  children: ReactNode;
};

const Button = (props: ButtonProps) => {
  return <button onClick={props.onCLick}>{props.children}</button>;
};

type StepMessageProps = {
  step: number;
  children: ReactNode;
};

const StepMessage = (props: StepMessageProps) => {
  return (
    <p className="message">
      <h3>Step {props.step}</h3>
      {props.children}
    </p>
  );
};

export default App;
