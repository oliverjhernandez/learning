import Logo from "./Logo";
import Form from "./Form";
import PackingList, { itemProp } from "./PackingList";
import Stats from "./Stats";
import "./index.css";
import { useState } from "react";

function App() {
  const [items, setItems] = useState<itemProp[]>([]);

  const handleAddItem = (i: itemProp) => {
    setItems((items) => [...items, i]);
  };

  const handleDeleteItem = (id: number) => {
    setItems((items) => items.filter((i) => i.id !== id));
  };

  const handleToggleItem = (id: number) => {
    setItems((items) =>
      items.map((i) => {
        return i.id === id ? { ...i, packed: !i.packed } : i;
      }),
    );
  };

  return (
    <div className="app">
      <Logo />
      <Form onAddItems={handleAddItem} />
      <PackingList
        items={items}
        onDeleteItem={handleDeleteItem}
        onToggleItem={handleToggleItem}
      />
      <Stats />
    </div>
  );
}

export default App;
