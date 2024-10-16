import { Link } from "react-router-dom";

function App() {
  return (
    <div>
      <h1>Hello from the Main Page!</h1>
      <p>Here are some examples of links to other pages</p>
      <nav>
        <ul>
          <li>
            <Link to="profile">Profile Page</Link>
          </li>
        </ul>
      </nav>
    </div>
  );
}

export default App;
