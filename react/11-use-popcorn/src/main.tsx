import { StrictMode, useState } from "react";
import { createRoot } from "react-dom/client";
import StarRating from "./Components/StarRating";
import App from "./App";
import "./index.css";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <App />
  </StrictMode>,
);
