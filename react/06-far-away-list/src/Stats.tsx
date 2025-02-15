import "./index.css";
import { itemProp } from "./PackingList";

type StatsProps = {
  items: itemProp[];
};

const Stats = (props: StatsProps) => {
  if (props.items.length === 0) {
    return (
      <p className="stats">
        <em>Start adding some items to your packing list! 🚀</em>
      </p>
    );
  }
  const numItems = props.items.length;
  const numCheckedItems = props.items.filter((i) => i.packed).length;
  const perctCheckedItems = Math.round((numCheckedItems / numItems) * 100);

  return (
    <footer className="stats">
      <em>
        {perctCheckedItems === 100
          ? "You got everything! Ready to go ✈️ "
          : `🧳 You have ${numItems} items on your list and you already packed
        ${numCheckedItems} (${perctCheckedItems}%)`}
      </em>
    </footer>
  );
};

export default Stats;
