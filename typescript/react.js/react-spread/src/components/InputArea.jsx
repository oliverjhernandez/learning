import React, { useState } from 'react'
import PropTypes from 'prop-types'

const InputArea = props => {
  const [inputText, setInputText] = useState('')

  const handleChange = event => {
    const newValue = event.target.value
    setInputText(newValue)
  }

  return (
    <div className="form">
      <input
        onChange={handleChange}
        type="text"
        name="listInput"
        value={inputText}
      />
      <button
        onClick={() => {
          props.onNew(inputText)
          setInputText('')
        }}
      >
        <span>Add</span>
      </button>
    </div>
  )
}

InputArea.propTypes = {
  onNew: PropTypes.func.isRequired,
  event: PropTypes.object.isRequired
}

export default InputArea
