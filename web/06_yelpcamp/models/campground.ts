import { Document, Schema, model } from 'mongoose'
import { Review } from './Reviews'

interface ICampground extends Document {
  title: string
  price: number
  image: string
  description: string
  location: string
  reviews: [Schema.Types.ObjectId]
}

const campgroundSchema = new Schema({
  title: { type: String, required: true },
  price: { type: Number, required: true },
  image: { type: String, required: true },
  description: { type: String, required: true },
  location: { type: String, required: true },
  reviews: [{ type: Schema.Types.ObjectId, ref: 'Review' }],
})

campgroundSchema.post('findOneAndDelete', async function (doc) {
  console.log(doc)

  if (doc) {
    await Review.deleteMany({
      _id: {
        $in: doc.reviews,
      },
    })
  }
})

export const Campground = model<ICampground>('Campground', campgroundSchema)
