import { Document, Schema, model } from 'mongoose'

interface ICampGround extends Document {
  title: string
  price: Number
  image: string
  description: string
  location: String
}

const CampGroundSchema = new Schema({
  title: String,
  price: Number,
  image: String,
  description: String,
  location: String,
})

export const CampGround = model<ICampGround>('Campground', CampGroundSchema)
