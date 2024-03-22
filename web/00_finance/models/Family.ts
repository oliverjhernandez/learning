import { Document, Schema, model } from 'mongoose'

interface IFamily extends Document {
  name: string
  currency: string
}

const FamilySchema = new Schema<IFamily>(
  {
    name: { type: String, required: true },
    currency: { type: String, required: true },
  },
  { timestamps: true }
)

export const Family = model<IFamily>('Family', FamilySchema)
