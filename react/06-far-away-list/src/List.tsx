const initialItems: itemProp[] = [
  { id: 1, description: "Passports", quantity: 3, packed: false },
  { id: 2, description: "Socks", quantity: 2, packed: true },
];

const List = () => {
  return (
    <div className="list">
      <ul>
        {initialItems.map((item) => {
          return <Item key={item.id} {...item} />;
        })}
      </ul>
    </div>
  );
};

export default List;

type itemProp = {
  id: number;
  description: string;
  quantity: number;
  packed: boolean;
};

const Item = (item: itemProp) => {
  return (
    <li>
      <span style={item.packed ? { textDecoration: "line-through" } : {}}>
        {item.quantity} {item.description}
      </span>
      <button>❌</button>
    </li>
  );
};
