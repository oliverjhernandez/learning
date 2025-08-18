import { useEffect } from "react";

export function useKey(key: string, action: () => void) {
  const l_key = key.toLowerCase();

  useEffect(
    function () {
      const callback = (e: KeyboardEvent) => {
        if (e.code.toLowerCase() === l_key) {
          action();
        }
      };

      document.addEventListener("keydown", callback);

      return function () {
        document.removeEventListener("keydown", callback);
      };
    },
    [action, l_key],
  );
}
