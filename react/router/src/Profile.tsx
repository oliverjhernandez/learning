import { useParams } from "react-router-dom";
import Popeye from "./Popeye";
import Spinach from "./Spinach";
import DefaultProfile from "./DefaultProfile";

const Profile = () => {
  const { name } = useParams();

  return (
    <div>
      <h1>Hello from the Profile Page</h1>
      <p>So, how are you?!</p>
      <hr />
      <h2>The profile visited is here:</h2>
      {name == "popeye" ? (
        <Popeye />
      ) : name == "spinach" ? (
        <Spinach />
      ) : (
        <DefaultProfile />
      )}
    </div>
  );
};

export default Profile;
