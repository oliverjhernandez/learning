import React from 'react'

const HourGreeting = () => {
  const date = new Date()
  const currentHour = date.getHours()

  const timeStyle = {
    color: 'red'
  }

  let greeting = ''
  if (currentHour > 12) {
    greeting = 'Good Afternoon'
    timeStyle.color = 'blue'
  } else {
    greeting = 'Good Morning'
    timeStyle.color = 'Red'
  }

  return (
    <div>
      <h1
        className="Hour"
        style={timeStyle}
      >
        Currently its {`${currentHour}. ${greeting}.`}
      </h1>
    </div>
  )
}

export default HourGreeting
