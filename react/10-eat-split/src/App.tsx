import { useState, type ReactNode } from "react";
import "./App.css";

let initialFriends: Friend[] = [
  {
    id: "118836",
    name: "Clark",
    image: "https://i.pravatar.cc/48?u=118836",
    balance: -7,
  },
  {
    id: "933372",
    name: "Sarah",
    image: "https://i.pravatar.cc/48?u=933372",
    balance: 20,
  },
  {
    id: "499476",
    name: "Anthony",
    image: "https://i.pravatar.cc/48?u=499476",
    balance: 0,
  },
];

type Friend = {
  id: string;
  name: string;
  image: string;
  balance: number;
};

function App() {
  const [friends, setFriends] = useState(initialFriends);
  const [showAddFriendForm, setShowAddFriendForm] = useState(false);
  const [selectedFriend, setSelectedFriend] = useState<Friend | null>(null);

  const handleShowAddFriend = () => {
    setShowAddFriendForm(!showAddFriendForm);
  };

  const handleSelectedFriend = (f: Friend) => {
    setSelectedFriend((s) => (s?.id === f.id ? null : f));
    setShowAddFriendForm(false);
  };

  const handleAddFriend = (f: Friend) => {
    setFriends([...friends, f]);
    setShowAddFriendForm(false);
    setSelectedFriend(null);
  };

  const handleSplitBill = (value: number) => {
    setFriends((friends) =>
      friends.map((f) =>
        f.id === selectedFriend?.id
          ? { ...f, balance: (f.balance += value) }
          : f,
      ),
    );
    setSelectedFriend(null);
  };

  return (
    <div className="app">
      <div className="sidebar">
        <FriendsList
          friendList={friends}
          selectedFriend={selectedFriend}
          onSelectedFriend={handleSelectedFriend}
        />

        {showAddFriendForm && <FormAddFriend onAddFriend={handleAddFriend} />}

        <Button onClick={handleShowAddFriend}>
          {showAddFriendForm ? "close" : "Add Friend"}
        </Button>
      </div>

      {selectedFriend && (
        <FormSplitBill
          selectedFriend={selectedFriend}
          onSplitBill={handleSplitBill}
          key={selectedFriend.id}
        />
      )}
    </div>
  );
}

type FriendListProps = {
  friendList: Friend[];
  selectedFriend: Friend | null;
  onSelectedFriend: (f: Friend) => void;
};

const FriendsList = (p: FriendListProps) => {
  return (
    <div>
      <ul>
        {p.friendList.map((f) => {
          return (
            <Friend
              key={f.id}
              friend={f}
              onSelectedFriend={p.onSelectedFriend}
              selectedFriend={p.selectedFriend}
            />
          );
        })}
      </ul>
    </div>
  );
};

type FriendProps = {
  friend: Friend;
  onSelectedFriend: (f: Friend) => void;
  selectedFriend: Friend | null;
};

const Friend = (p: FriendProps) => {
  const isSelected = p.friend.id === p.selectedFriend?.id;

  return (
    <li className={isSelected ? "selected" : ""} key={p.friend.id}>
      <img src={p.friend.image} alt="friend image" />
      <h3>{p.friend.name}</h3>
      {p.friend.balance < 0 && (
        <p className="red">
          You owe {p.friend.name} ${Math.abs(p.friend.balance)}
        </p>
      )}
      {p.friend.balance > 0 && (
        <p className="green">
          {p.friend.name} owes you ${p.friend.balance}
        </p>
      )}
      {p.friend.balance == 0 && <p>You and {p.friend.name} are even</p>}
      <Button
        onClick={() => {
          return p.onSelectedFriend(p.friend);
        }}
      >
        {isSelected ? "Close" : "Select"}
      </Button>
    </li>
  );
};

type FormAddFriendProps = {
  onAddFriend: (f: Friend) => void;
};

const FormAddFriend = (p: FormAddFriendProps) => {
  const [name, setName] = useState("");
  const [image, setImage] = useState("https://i.pravatar.cc/48");

  const id = crypto.randomUUID();

  const friend: Friend = {
    id: id,
    name: name,
    image: `${image}?=${id}`,
    balance: 0,
  };

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    if (!name || !image) return;

    p.onAddFriend(friend);

    setName("");
    setImage("https://i.pravatar.cc/48");
  };

  return (
    <form action="#" className="form-add-friend" onSubmit={handleSubmit}>
      <label htmlFor="friend-name">🧑🏻‍🦰Friend Name</label>
      <input
        id="friend-name"
        type="text"
        value={name}
        onChange={(e) => setName(e.target.value)}
      />
      <label htmlFor="friend-img">🏝️Image Url</label>
      <input
        id="friend-img"
        type="text"
        value={image}
        onChange={(e) => setImage(e.target.value)}
      />
      <Button onClick={() => ""}>Add</Button>
    </form>
  );
};

type ButtonProps = {
  children: ReactNode;
  onClick: () => void;
};

const Button = (p: ButtonProps) => {
  return (
    <button onClick={p.onClick} className="button">
      {p.children}
    </button>
  );
};

type FormSplitBillProps = {
  selectedFriend: Friend;
  onSplitBill: (n: number) => void;
};

const FormSplitBill = (p: FormSplitBillProps) => {
  const [bill, setBill] = useState(0);
  const [paidByUser, setPaidByUser] = useState(0);
  const paidByFriend = bill ? bill - paidByUser : 0;
  const [whoIsPaying, setWhoIsPaying] = useState("user");

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    if (!bill || !paidByUser) return;

    p.onSplitBill(whoIsPaying == "user" ? paidByFriend : -paidByUser);
  };

  return (
    <form action="#" className="form-split-bill" onSubmit={handleSubmit}>
      <h2>Split a bill with {p.selectedFriend.name}</h2>

      <label htmlFor="value">💰Bill Value</label>
      <input
        id={String(bill)}
        type="text"
        onChange={(e) => setBill(Number(e.target.value))}
      />

      <label htmlFor="expense">🏌🏻‍♀️Your Expense</label>
      <input
        id="expense"
        value={String(paidByUser)}
        type="text"
        onChange={(e) =>
          setPaidByUser(
            Number(e.target.value) > bill ? paidByUser : Number(e.target.value),
          )
        }
      />

      <label htmlFor="friend-expense">
        👩🏼‍💻{p.selectedFriend.name}'s expense
      </label>
      <input type="text" id="friend-expense" value={paidByFriend} disabled />

      <label htmlFor="who-paying">🤑Who is paying the bill</label>
      <select
        name="who-paying"
        id="who-paying"
        value={whoIsPaying}
        onChange={(e) => setWhoIsPaying(e.target.value)}
      >
        <option value="you">You</option>
        <option value="friend">Mark</option>
      </select>

      <button className="button">Split Bill</button>
    </form>
  );
};

export default App;
