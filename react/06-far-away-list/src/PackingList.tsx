type PackingListProps = {
  items: itemProp[];
  onDeleteItem: (id: number) => void;
  onToggleItem: (id: number) => void;
};

const PackingList = (props: PackingListProps) => {
  return (
    <div className="list">
      <ul>
        {props.items.map((i) => {
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
    </div>
  );
};

export default PackingList;

export type itemProp = {
  id: number;
  description: string;
  quantity: number;
  packed: boolean;
};

type ItemProps = {
  onDeleteItem: (id: number) => void;
  item: itemProp;
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
