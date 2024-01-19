import React, { useState } from 'react'
import ToDoItem from './ToDoItem'
import InputArea from './InputArea'

function App() {
  const [items, setItems] = useState([])

  const addItem = (inputText) => {
    setItems(prevValues => {
      return [...prevValues, inputText]
    })
  }

  const deleteItem = id => {
    setItems(prevItems => {
      return prevItems.filter((elem, i) => {
        return i !== id
      })
    })
  }

  return (
    <div className="container">
      <div className="heading">
        <h1>To-Do List</h1>
      </div>
      <div />
      <InputArea
        onNew={addItem}
      />
      <div>
        <ul>
          {items.map((todoItem, i) => (
            <ToDoItem
              text={todoItem}
              onChecked={deleteItem}
              key={i}
              id={i}
            />
          ))}
        </ul>
      </div>
    </div>
  )
}

export default App
