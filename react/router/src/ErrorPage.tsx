import { Link } from "react-router-dom";

const ErrorPage = () => {
  return (
    <div>
      <h1>Oh no!, this route doesnt exist!</h1>
      <Link to="/">Go back to the main page</Link>
    </div>
  );
};

export default ErrorPage;
