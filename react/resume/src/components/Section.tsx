import "../styles/section.css";

function Section() {
  return (
    <div className="section section-container">
      <div className="section-title">Title1 --- --- ---</div>
      <div className="section-body section-body-container">
        <div className="section-body-item">
          <label>Name</label> <input type="text" />
        </div>
        <div className="section-body-item">item2</div>
        <div className="section-body-item">item3</div>
        <div className="section-body-item">item4</div>
        <div className="section-body-item">item5</div>
        <div className="section-body-item">item6</div>
      </div>
    </div>
  );
}

export default Section;
