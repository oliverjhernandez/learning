import { Document, Schema, model } from 'mongoose'

interface IAccount extends Document {
  subtype: string
  familyId: Schema.Types.ObjectId
  name: string
  accountableType: string
  accountableId: Schema.Types.ObjectId
  balance: number
  currency: Schema.Types.ObjectId
  classification: string
}

const AccountSchema = new Schema<IAccount>(
  {
    subtype: { type: String, required: true },
    familyId: { type: Schema.Types.ObjectId, required: true },
    name: { type: String, required: true },
    accountableType: { type: String, required: true },
    accountableId: { type: Schema.Types.ObjectId, required: true },
    balance: { type: Number, required: true },
    currency: { type: Schema.Types.ObjectId, required: true },
    classification: { type: String, required: true },
  },
  { timestamps: true }
)

export const Account = model<IAccount>('Account', AccountSchema)
