import React from 'react'
import PropTypes from 'prop-types'

const Input = () => {
  return (
    <input
      type="text"
      placeholder="Username"
    />
  )
}

Input.propTypes = {
  type: PropTypes.string,
  placeholder: PropTypes.string
}

export default Input
