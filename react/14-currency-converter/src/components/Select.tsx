type SelectProps = {
  name: string;
  currencies: string[];
  select: string;
  setSelect: React.Dispatch<React.SetStateAction<string>>;
};

const Select = ({ name, currencies, select, setSelect }: SelectProps) => {
  return (
    <select
      name={name}
      value={select}
      onChange={(e) => setSelect(e.target.value)}
    >
      {currencies.map((curr) => (
        <option key={curr} value={curr}>
          {curr}
        </option>
      ))}
    </select>
  );
};

export default Select;
