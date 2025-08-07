import { useEffect, useState } from "react";
import "./App.css";
import Input from "./components/Input";
import Select from "./components/Select";

// {
//   amount: 100.0,
//   base: "EUR",
//   date: "2025-08-06",
//   rates: { USD: 116.04 },
// }

type ConversionResponse = {
  amount: number;
  base: string;
  date: string;
  rates: Record<string, number>;
  message?: string;
};

export const currencies = ["USD", "EUR", "CAD", "INR"];

function App() {
  const [amount, setAmount] = useState(0);
  const [first, setFirst] = useState("USD");
  const [second, setSecond] = useState("EUR");
  const [conversion, setConversion] = useState(0);
  const [isLoading, setIsLoading] = useState(false);

  const handleAmountChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    if (value === "" || /^\d+$/.test(value)) {
      setAmount(Number(value));
    }
  };

  useEffect(
    function () {
      const controller = new AbortController();

      const fetchConversion = async () => {
        try {
          setIsLoading(true);
          const url = `https://api.frankfurter.app/latest?amount=${amount}&from=${first}&to=${second}`;

          const init: RequestInit = {
            signal: controller.signal,
          };

          const res = await fetch(url, init);

          if (!res.ok)
            throw new Error("something went wrong while fetching conversion");

          const data = (await res.json()) as ConversionResponse;

          const rateValue = data.rates[second];
          if (rateValue === undefined) {
            throw new Error("unexpected response format");
          }
          setConversion(rateValue);

          if (data.message === "not found") {
            throw new Error("not found");
          }

          setIsLoading(false);
        } catch (err) {
          if (err instanceof Error) {
            console.log(err.message);
          }
        }
      };

      if (amount === 0) {
        return;
      }

      if (first === second) {
        setConversion(amount);
        return;
      }

      if (isNaN(amount) || amount <= 0) return;

      fetchConversion();

      return () => {
        controller.abort();
      };
    },
    [amount, first, second],
  );
  return (
    <div>
      <Input amount={amount} onAmountChange={handleAmountChange} />
      <Select
        currencies={currencies}
        name="first"
        select={first}
        setSelect={setFirst}
      />
      <Select
        currencies={currencies}
        name="second"
        select={second}
        setSelect={setSecond}
      />
      <p> {isLoading ? "Loading..." : `${conversion} ${second}`}</p>
    </div>
  );
}

export default App;
