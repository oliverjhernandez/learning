import "../styles/section.css";

type ItemTitleProps = {
  label: string;
};

function ItemTitle({ label }: ItemTitleProps) {
  return (
    <div className="section-title">
      <div>{label}</div>
      <span>
        <hr />
      </span>
    </div>
  );
}

export default ItemTitle;
