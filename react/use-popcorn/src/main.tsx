import { StrictMode, useState } from "react";
import { createRoot } from "react-dom/client";
import StarRating from "./Components/StarRating";
// import App from "./App";
// import "./index.css";

const Test = () => {
  const [movieRating, setMovieRating] = useState(4);

  return (
    <div>
      <StarRating
        size={24}
        maxRating={10}
        color="red"
        onExternalSetRating={setMovieRating}
      />

      <p>Hello, {movieRating}!!</p>
    </div>
  );
};

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <Test />
  </StrictMode>,
);
