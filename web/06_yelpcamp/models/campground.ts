import { Document, Schema, model } from 'mongoose'

interface ICampGround extends Document {
  title: string
  price: string
  description: string
  location: String
}

const CampGroundSchema = new Schema({
  title: String,
  price: String,
  description: String,
  location: String,
})

export const CampGround = model<ICampGround>('Campground', CampGroundSchema)
