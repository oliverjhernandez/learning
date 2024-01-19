import React from 'react'
import Avatar from './Avatar'
import Detail from './Detail'
import PropTypes from 'prop-types'

const Card = props => {
  return (
    <div className="card">
      <div className="top">
        <p>{props.id}</p>
        <h2 className="name">{props.name}</h2>
        <Avatar img={props.img} />
      </div>
      <div className="bottom">
        <Detail detailInfo={props.tel} />
        <Detail detailInfo={props.email} />
      </div>
    </div>
  )
}

Card.propTypes = {
  id: PropTypes.number,
  name: PropTypes.string,
  img: PropTypes.string,
  tel: PropTypes.string,
  email: PropTypes.string
}

export default Card
