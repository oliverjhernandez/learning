import { useState } from "react";

type PackingListProps = {
  items: itemProp[];
  onDeleteItem: (id: number) => void;
  onToggleItem: (id: number) => void;
  onHandleClearList: () => void;
};

const PackingList = (props: PackingListProps) => {
  const [sortBy, setSortBy] = useState("input");

  let sortedItems: itemProp[] = [];

  if (sortBy == "input") sortedItems = props.items;
  if (sortBy == "description")
    sortedItems = props.items.slice().sort((a, b) => {
      return a.description.localeCompare(b.description);
    });
  if (sortBy == "packed")
    sortedItems = props.items
      .slice()
      .sort((a, b) => Number(a.packed) - Number(b.packed));

  const emptyList = (e: React.SyntheticEvent<HTMLFormElement>) => {};

  return (
    <div className="list">
      <ul>
        {sortedItems.map((i) => {
          return (
            <Item
              key={i.id}
              item={i}
              onDeleteItem={props.onDeleteItem}
              onToggleItem={props.onToggleItem}
            />
          );
        })}
      </ul>

      <div className="actions">
        <select
          onChange={(e) => {
            setSortBy(e.target.value);
          }}
          value={sortBy}
        >
          <option value="input">Sort by input order</option>
          <option value="description">Sort by description</option>
          <option value="packed">Sort by packed status</option>
        </select>
        <button onClick={props.onHandleClearList}>Clear List</button>
      </div>
    </div>
  );
};

export default PackingList;

export type itemProp = {
  id: number;
  description: string;
  packed: boolean;
  quantity: number;
};

type ItemProps = {
  item: itemProp;
  onDeleteItem: (id: number) => void;
  onToggleItem: (id: number) => void;
};

const Item = (props: ItemProps) => {
  return (
    <li>
      <input
        type="checkbox"
        value={String(props.item.packed)}
        onChange={() => props.onToggleItem(props.item.id)}
      />
      <span style={props.item.packed ? { textDecoration: "line-through" } : {}}>
        {props.item.quantity} {props.item.description}
      </span>
      <button onClick={() => props.onDeleteItem(props.item.id)}>❌</button>
    </li>
  );
};
