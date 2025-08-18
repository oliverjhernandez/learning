import { useRef } from "react";
import "../index.css";
import { useKey } from "../Hooks/useKey";

type SearchProps = {
  query: string;
  setQuery: React.Dispatch<React.SetStateAction<string>>;
};

const Search = ({ query, setQuery }: SearchProps) => {
  const inputElem = useRef<HTMLInputElement>(null);

  useKey("Enter", () => {
    if (document.activeElement === inputElem.current) return;
    if (inputElem.current) {
      inputElem.current.focus();
    }
    setQuery("");
  });

  return (
    <input
      className="search"
      type="text"
      placeholder="Search movies..."
      value={query}
      onChange={(e) => setQuery(e.target.value)}
      ref={inputElem}
    />
  );
};

export default Search;
