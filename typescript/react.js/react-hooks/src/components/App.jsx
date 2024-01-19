import React, { useState } from 'react'

const App = () => {
  // const [count, setCount] = useState(10)

  const updateTime = () => {
    let now = new Date().toLocaleTimeString()
    setTime(now)
  }

  const [now, setTime] = useState('TIME')
  const setInterval = (func, interval) => {
    function wrapper() {
      func()
      setTimeout(wrapper, interval)
    }

    setTimeout(wrapper, interval)
  }

  setInterval(updateTime, 1000)

  // const increase = () => {
  //   setCount(count + 1)
  // }

  // const decrease = () => {
  //   setCount(count - 1)
  // }

  return (
    <div className="container">
      <h1>{count}</h1>
      {/* <h1>{now}</h1> */}
      <button onClick={increase}>+</button>
      <button onClick={decrease}>-</button>
      <button onClick={updateTime}>Get Time</button>
    </div>
  )
}

export default App
