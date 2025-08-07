type InputProps = {
  amount: number;
  onAmountChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
};

const Input = ({ amount, onAmountChange }: InputProps) => {
  return (
    <input name="amount" type="text" value={amount} onChange={onAmountChange} />
  );
};

export default Input;
