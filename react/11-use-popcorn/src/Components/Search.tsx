import { useEffect, useRef } from "react";
import "../index.css";

type SearchProps = {
  query: string;
  setQuery: React.Dispatch<React.SetStateAction<string>>;
};

const Search = ({ query, setQuery }: SearchProps) => {
  const inputElem = useRef<HTMLInputElement>(null);

  useEffect(function () {
    const callback = (e: KeyboardEvent) => {
      if (document.activeElement === inputElem.current) {
        return;
      }

      if (e.code === "Enter" && inputElem.current !== null) {
        inputElem.current.focus();
        setQuery("");
      }
    };

    document.addEventListener("keydown", callback);

    return () => {
      document.removeEventListener("keydown", callback);
    };
  }, []);

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
