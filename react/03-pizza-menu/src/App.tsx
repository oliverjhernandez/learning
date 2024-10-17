import Header from "./Header.tsx";
import Footer from "./Footer.tsx";
import Menu from "./Menu.tsx";
import "./index.css";

const App = () => {
  return (
    <div className="container">
      <Header />
      <Menu />
      <Footer />
    </div>
  );
};

export default App;
