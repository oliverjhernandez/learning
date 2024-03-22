import { Document, Schema, model } from 'mongoose'

interface ICurrency extends Document {
  name: string
  isoCode: string
}

const CurrencySchema = new Schema<ICurrency>(
  {
    name: { type: String, required: true },
    isoCode: { type: String, required: true },
  },
  { timestamps: true }
)

export const Currency = model<ICurrency>('Currency', CurrencySchema)
