type Product = {
  title: string;
  id: number;
};

const products: Product[] = [
  { title: "Cabbage", id: 1 },
  { title: "tomatoes", id: 2 },
  { title: "potatos", id: 3 },
];

const Products = () => {
  const listItems = products.map((prod) => {
    return <li key={prod.id}>{prod.title}</li>;
  });

  return <ul>{listItems}</ul>;
};

export default Products;
