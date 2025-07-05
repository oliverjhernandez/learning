import { useState, type ReactNode } from "react";
import "./styles.css";

const faqs: Faq[] = [
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

type Faq = {
  title: string;
  text: string;
  number: number;
};

type AccordionProps = {
  faqs: Faq[];
};

const Accordion = (props: AccordionProps) => {
  const [currentOpen, setCurrentOpen] = useState<number | null>(null);

  return (
    <div className="accordion">
      {props.faqs.map((el, i) => {
        return (
          <AccordionItem
            number={el.number}
            title={el.title}
            key={i}
            currentOpen={currentOpen}
            onOpen={setCurrentOpen}
          >
            {el.text}
          </AccordionItem>
        );
      })}
    </div>
  );
};

type AccordionItemProps = {
  number: number;
  title: string;
  currentOpen: number | null;
  onOpen: React.Dispatch<React.SetStateAction<number | null>>;
  children: ReactNode;
};

const AccordionItem = (props: AccordionItemProps) => {
  const isOpen = props.number == props.currentOpen;

  const handleToggle = () => {
    props.onOpen(isOpen ? null : props.number);
  };
  return (
    <div className={`item ${isOpen ? "open" : ""}`} onClick={handleToggle}>
      <p className="number">
        {props.number < 9 ? `0${props.number}` : `props.number + 1`}
      </p>
      <p className="title">{props.title}</p>
      <p className="icon">{isOpen ? "-" : "+"}</p>
      {isOpen && <div className="content-box">{props.children}</div>}
    </div>
  );
};

export default App;
