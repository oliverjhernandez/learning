import { Document, Schema, model } from 'mongoose'

export interface IReview extends Document {
  body: string
  rating: number
}

const reviewSchema = new Schema({
  body: String,
  rating: Number,
})

export const Review = model<IReview>('Review', reviewSchema)
