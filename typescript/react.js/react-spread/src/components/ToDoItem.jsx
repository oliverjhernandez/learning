import React from 'react'
import PropTypes from 'prop-types'

const ToDoItem = props => {
  return (
    <div
      onClick={() => {
        props.onChecked(props.id)
      }}
    >
      <li>{props.text}</li>
    </div>
  )
}

ToDoItem.propTypes = {
  text: PropTypes.string.isRequired,
  onChecked: PropTypes.func.isRequired,
  id: PropTypes.number.isRequired
}

export default ToDoItem
