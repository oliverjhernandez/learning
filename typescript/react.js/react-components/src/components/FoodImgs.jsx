import React from 'react'
import '../App.css'

const FoodImgs = () => {
  const baconImage =
    'https://assets.epicurious.com/photos/630386df4eaa77f49a537aa5/1:1/w_2240,c_limit/Brown_Sugar_Bacon_RECIPE_081722_38635.jpg'
  const ramenImage = 'https://www.justonecookbook.com/wp-content/uploads/2023/04/Spicy-Shoyu-Ramen-8025-II.jpg'
  const hotDogImage =
    'https://www.thespruceeats.com/thmb/3RnYaB9iJdkzJ-2AROuLzA1vYTY=/750x0/filters:no_upscale():max_bytes(150000):strip_icc():format(webp)/homemade-hot-dogs-1808491-final-01-ddf78878cb784ccdb617a2ae31674b75.jpg'

  return (
    <div>
      <img
        className="FoodImgs"
        src={baconImage}
        alt="Bacon"
      />
      <img
        className="FoodImgs"
        src={ramenImage}
        alt="Ramen"
      />
      <img
        className="FoodImgs"
        src={hotDogImage}
        alt="Hot Dogs"
      />
    </div>
  )
}

export default FoodImgs
