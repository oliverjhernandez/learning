import "../styles/section.css";

type ItemProps = {
  label: string;
};

function Item({ label }: ItemProps) {
  return (
    <div className="section-body-item">
      <label className="label">{label}</label> <input type="text" />
    </div>
  );
}

export default Item;
