import { useEffect, useState } from "react";

export function useLocalStorageState<T>(initialState: T[], key: string) {
  const [value, setValue] = useState<T[]>(function () {
    try {
      const storedValue = localStorage.getItem(key);
      return storedValue ? (JSON.parse(storedValue) as T[]) : initialState;
    } catch (err) {
      if (err instanceof Error) {
        console.log(err.message);
      }
      return initialState;
    }
  });

  useEffect(
    function () {
      localStorage.setItem(key, JSON.stringify(value));
    },
    [value, key],
  );

  return [value, setValue] as const;
}
