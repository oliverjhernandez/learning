import "../styles/section.css";
import ItemTitle from "./ItemTitle";
import Item from "./Item";

function Section() {
  return (
    <div className="section section-container">
      <ItemTitle label="Personal Info" />
      <div className="section-body section-body-container">
        <Item label="Name" />
        <Item label="Surname" />
        <Item label="Phone" />
        <Item label="Something" />
        <Item label="Else" />
        <Item label="Entirely" />
      </div>
    </div>
  );
}

export default Section;
