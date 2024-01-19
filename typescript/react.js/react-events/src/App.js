import React, { useState } from 'react'

function App() {
  const [isMouseOver, setIsMouseOver] = useState(false)
  const [name, setName] = useState('')
  const [headingText, setHeadingText] = useState('')

  const handleClick = (event) => {
    setHeadingText(name)
    event.preventDefault()
  }

  const handleMouseOver = () => {
    setIsMouseOver(true)
  }

  const handleMouseOut = () => {
    setIsMouseOver(false)
  }

  const handleChange = event => {
    console.log(event.target.value)
    setName(event.target.value)
  }

  return (
    <div className="container">
      <h1> Hello {headingText} </h1>
      <form onSubmit={handleClick}>
        <input
          onChange={handleChange}
          type="text"
          placeholder="What's your name?"
          value={name}
        />
        <button
          type='submit'
          style={{ backgroundColor: isMouseOver ? 'black' : 'white' }}
          onClick={handleClick}
          onMouseOver={handleMouseOver}
          onMouseOut={handleMouseOut}
        >
          Submit
        </button>
      </form>
    </div>
  )
}

export default App
