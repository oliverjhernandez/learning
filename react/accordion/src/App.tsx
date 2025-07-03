import { useState } from "react";
import "./styles.css";

const faqs: AccordionItemProps[] = [
  {
    number: 1,
    title: "Where are these chairs assembled?",
    text: "Lorem ipsum dolor sit amet consectetur, adipisicing elit. Accusantium, quaerat temporibus quas dolore provident nisi ut aliquid ratione beatae sequi aspernatur veniam repellendus.",
  },
  {
    number: 2,
    title: "How long do I have to return my chair?",
    text: "Pariatur recusandae dignissimos fuga voluptas unde optio nesciunt commodi beatae, explicabo natus.",
  },
  {
    number: 3,
    title: "Do you ship to countries outside the EU?",
    text: "Excepturi velit laborum, perspiciatis nemo perferendis reiciendis aliquam possimus dolor sed! Dolore laborum ducimus veritatis facere molestias!",
  },
];

const App = () => {
  return (
    <div>
      <Accordion faqs={faqs} />
    </div>
  );
};

type Faqs = {
  title: string;
  text: string;
  number: number;
};

type AccordionProps = {
  faqs: Faqs[];
};

const Accordion = (props: AccordionProps) => {
  return (
    <div className="accordion">
      {props.faqs.map((el, i) => {
        return (
          <AccordionItem
            title={el.title}
            number={el.number}
            text={el.text}
            key={i}
          />
        );
      })}
    </div>
  );
};

type AccordionItemProps = {
  number: number;
  title: string;
  text: string;
};

const AccordionItem = (props: AccordionItemProps) => {
  const [isOpen, setIsOpen] = useState(false);

  const handleToggle = () => {
    setIsOpen(!isOpen);
  };

  return (
    <div className={`item ${isOpen ? "open" : ""}`} onClick={handleToggle}>
      <p className="number">
        {props.number < 9 ? `0${props.number}` : `props.number + 1`}
      </p>
      <p className="title">{props.title}</p>
      <p className="icon">{isOpen ? "-" : "+"}</p>
      {isOpen && <div className="content-box">{props.text}</div>}
    </div>
  );
};

export default App;
