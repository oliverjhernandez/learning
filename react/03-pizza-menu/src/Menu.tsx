import Pizza, { PizzaProps } from "./Pizza";

const pizzaData: PizzaProps[] = [
  {
    name: "Focaccia",
    ingredients: "Bread with italian olive oil and rosemary",
    price: 6,
    photoName: "pizzas/focaccia.jpg",
    soldOut: false,
  },
  {
    name: "Pizza Margherita",
    ingredients: "Tomato and mozarella",
    price: 10,
    photoName: "pizzas/margherita.jpg",
    soldOut: false,
  },
  {
    name: "Pizza Spinaci",
    ingredients: "Tomato, mozarella, spinach, and ricotta cheese",
    price: 12,
    photoName: "pizzas/spinaci.jpg",
    soldOut: false,
  },
  {
    name: "Pizza Funghi",
    ingredients: "Tomato, mozarella, mushrooms, and onion",
    price: 12,
    photoName: "pizzas/funghi.jpg",
    soldOut: false,
  },
  {
    name: "Pizza Salamino",
    ingredients: "Tomato, mozarella, and pepperoni",
    price: 15,
    photoName: "pizzas/salamino.jpg",
    soldOut: true,
  },
  {
    name: "Pizza Prosciutto",
    ingredients: "Tomato, mozarella, ham, aragula, and burrata cheese",
    price: 18,
    photoName: "pizzas/prosciutto.jpg",
    soldOut: false,
  },
];

const Menu = () => {
  return (
    <main className="menu">
      <h2>Our Menu</h2>

      <p>
        Authentic Italian cuisine. Creative dishes to choose from. All from our
        stone oven, all organic, all delicious.
      </p>

      <ul className="pizzas">
        {pizzaData.map((pizza: PizzaProps) => {
          return (
            <Pizza
              name={pizza.name}
              key={pizza.name}
              ingredients={pizza.ingredients}
              photoName={pizza.photoName}
              soldOut={pizza.soldOut}
              price={pizza.price}
            />
          );
        })}
      </ul>
    </main>
  );
};

export default Menu;
